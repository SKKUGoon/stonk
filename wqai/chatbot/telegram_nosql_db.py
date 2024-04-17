from typing import Dict
from datetime import datetime
import os
import json
import copy


class JsonDatabase:
    # Temporarily store the data on json database
    # {
    #     "<user_id>": {
    #         "<context_id>": {
    #             "context": "<context text>",
    #             "filename": "<filename of the context>",
    #             "summary": "<summary of the context>",
    #             "used": bool,
    #             "log": {
    #                 "qa": [
    #                     {
    #                         "q": "<question>",
    #                         "a": "<answer>"
    #                     },
    #                     ...
    #                 ]
    #             },
    #         }
    #     }
    # }
    def __init__(self, name: str, verbose: bool = False):
        self.dbpath = f"{name}.json"
        self.db = dict()
        self.verbose = verbose

    def default(self, override: bool = False):
        """
        Create json database
        :param override: if override is true, delete the past json and create a new one
        """
        if os.path.exists(self.dbpath):
            print(f"{self.dbpath} already exists")
            self._dn_sync()
            self._db_stat(self.db)

        else:
            print(f"No existing {self.dbpath} found. Creating new empty file")
            self._up_sync()

        if override:
            print("Overriding past database. Deleting all contents")
            self.db = dict()
            self._up_sync()
            self._db_stat(self.db)
            return

    @staticmethod
    def _db_stat(db: Dict):
        usr_num = len(list(db.keys()))
        ctx_num = 0
        for usr in db.keys():
            ctx_num += len(list(db[usr].keys()))

        msg = f"Total User    : {usr_num}\nTotal Context : {ctx_num}"
        print(msg)

    def _dn_sync(self):
        with open(self.dbpath, 'r') as file:
            self.db = json.load(file)

    def _up_sync(self):
        with open(self.dbpath, "w") as file:
            json.dump(self.db, file)

    @staticmethod
    def _new_context_for_user(ctx_id: str):
        return {
            "context": None,
            "filename": ctx_id,
            "summary": None,
            "used": True,
            "log": {
                "qa": list()
            }
        }

    def _id_ctx_id_guard(self, user_id: str, context_id: str):
        try:
            assert user_id in self.db.keys()
            assert context_id in self.db[user_id].keys()
        except AssertionError:
            print(f"no {user_id} or no {context_id}. create it first")

    def _id_guard(self, user_id: str):
        try:
            assert user_id in self.db.keys()
        except AssertionError:
            print(f"no {user_id}. create it first")

    def add_user(self, user_id: str) -> str:
        """
        Add user_id to the json database
        :param user_id:
        :return: user_id
        """
        if user_id in self.db.keys():
            return user_id

        print(f"created user with id {user_id}")
        self.db[user_id] = dict()
        self._up_sync()

        return user_id

    def add_user_context(self, user_id: str, context_id: str):
        """
        Set filename to context_id
        :param user_id:
        :param context_id:
        :return:
        """
        try:
            assert user_id in self.db.keys()
        except AssertionError:
            print(f"no {user_id}, create user first")
            return

        if context_id in self.db[user_id].keys():
            # Force update everytime the new file is given
            print(f"context id {context_id} already exists. Archiving")
            ts = datetime.now().timestamp()

            # Move the new context to old
            self.db[user_id][context_id]["used"] = False
            self.db[user_id][f"{context_id}_{ts}"] = copy.deepcopy(self.db[user_id][context_id])

        for ctx_id in self.db[user_id].keys():
            self.db[user_id][ctx_id]["used"] = False

        # Create a new context
        self.db[user_id][context_id] = self._new_context_for_user(context_id)
        self._up_sync()

    def update_context_summary(self, user_id: str, summarized: str):
        self._id_guard(user_id)

        context_id = None
        for ctx_id, ctx_config in self.db[user_id].items():
            if ctx_config["used"]:
                context_id = ctx_id
                break

        if context_id is None:
            print(f"no current context_id {context_id} found")
            return

        self.db[user_id][context_id]["summary"] = summarized
        self._up_sync()

    def update_context_context(self, user_id: str, cleaned: str):
        self._id_guard(user_id)

        context_id = None
        for ctx_id, ctx_config in self.db[user_id].items():
            if ctx_config["used"]:
                context_id = ctx_id
                break

        if context_id is None:
            print(f"no current context_id {context_id} found")
            return

        self.db[user_id][context_id]["context"] = cleaned
        self._up_sync()

    def update_question_answer_log(self, user_id: str, qa: Dict):
        self._id_guard(user_id)

        try:
            assert "q" in qa.keys() and "a" in qa.keys()

            context_id = None
            for ctx_id, ctx_config in self.db[user_id].items():
                if ctx_config["used"]:
                    context_id = ctx_id
                    break

            if context_id is None:
                print(f"no current context_id {context_id} found")
                return

            self.db[user_id][context_id]["log"]["qa"].append(qa)
        except AssertionError:
            print(f"qa dict {qa} does not match our format. Must include `q` and `a` key")

        self._up_sync()

    # Each user can have only one context running at a time
    def get_current_context(self, user_id: str) -> str:
        self._id_guard(user_id)

        # There is only one working context
        text_ctx = None
        for ctx_id, ctx_config in self.db[user_id].items():
            if ctx_config["used"]:
                text_ctx = ctx_config["context"]
                break

        if text_ctx is not None:
            return text_ctx
        else:
            print(f"no used context for {user_id}")

    def get_current_summary(self, user_id: str) -> str:
        self._id_guard(user_id)

        text_ctx = None
        for ctx_id, ctx_config in self.db[user_id].items():
            if ctx_config["used"]:
                text_ctx = ctx_config["summary"]

        if text_ctx is not None:
            return text_ctx
        else:
            print(f"no context for {user_id}")


if __name__ == "__main__":
    j = JsonDatabase("testdb")
    j.default(override=True)

    j.add_user("test_id")
    j._db_stat(j.db)

    j.add_user_context("test_id", "context_id")
    j._db_stat(j.db)

