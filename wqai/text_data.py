def clean_text(path: str):
    with open(path, 'r') as file:
        text = file.read()
        return text


if __name__ == "__main__":
    clean_text("conference_call_rklb.txt")
