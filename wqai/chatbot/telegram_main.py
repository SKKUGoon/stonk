from telethon import TelegramClient, events

from telegram_auth import TelegramAuth
from telegram_handler import start_message, iden
from telegram_nosql_db import JsonDatabase

from models import qa_nlp, summary_nlp

auth = TelegramAuth("test_session")
client = TelegramClient(auth.session_name, auth.app_id, auth.app_hash)
client.start(bot_token=auth.token)
jdb = JsonDatabase("test_db")
jdb.default(override=True)


# Define the wrapper function for each handler functions
# Define handler functions in `telegram_handler.py` for
# better maintenance.
@client.on(events.NewMessage(pattern='/start'))
async def start(event):
    sender = await event.get_sender()
    id_ = iden(sender)

    print(id_, "connected, added to database")
    jdb.add_user(id_)
    msg = start_message()

    await event.respond(msg)


@client.on(events.NewMessage())
async def watchman(event):
    """
    Watch all the message.
    :param event:
    :return:
    """
    print(event.message)


@client.on(events.NewMessage())
async def get_context_text(event):
    if event.message.file:
        sender = await event.get_sender()
        id_ = iden(sender)

        if event.message.file.name.endswith('.txt'):
            ctx_id = event.message.file.name
            await event.respond("Processing your text file... please wait")

            path = await event.message.download_media(file='./tmp')
            with open(path, 'r') as file:
                content = file.read()

                # Create a context data
                jdb.add_user_context(id_, ctx_id)
                jdb.update_context_context(id_, content)

                # Create a summary
                article = jdb.get_current_context(id_)
                summary = summary_nlp(article, truncation=True, max_length=130, min_length=30, do_sample=False)
                jdb.update_context_summary(id_, summary[0]['summary_text'])

            await event.respond(f"Successfully uploaded. Here's a brief summary\n\n{jdb.get_current_summary(id_)}")

        else:
            await event.respond("Please send a text file(`.txt`)")


@client.on(events.NewMessage(pattern='/summary'))
async def answer_question(event):
    sender = await event.get_sender()
    id_ = iden(sender)

    summary = jdb.get_current_summary(id_)
    if summary is not None:
        await event.respond(summary)
    else:
        await event.respond("Summary has not been created yet")


@client.on(events.NewMessage(pattern='/question'))
async def answer_question(event):
    sender = await event.get_sender()
    id_ = iden(sender)

    # Prepare the question and its context
    text = jdb.get_current_context(id_)
    question = str(event.message.message).replace("/question ", "")
    qa_input = {"question": question, "context": text}
    ans = qa_nlp(qa_input)

    await event.respond(f"Answer\n{ans['answer']}")

    jdb.update_question_answer_log(id_, {"q": question, "a": ans['answer']})


def main():
    """
    Start the bot client
    """
    print("------ telegram bot starting ------")
    client.run_until_disconnected()


if __name__ == "__main__":
    main()
