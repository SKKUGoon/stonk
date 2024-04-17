from transformers import AutoModelForQuestionAnswering, AutoTokenizer
from transformers import pipeline

from text_data import clean_text

# Text
text = clean_text("conference_call_nvda.txt")

# model_name = "deepset/roberta-base-squad2"
#
# # a) Get predictions
# nlp = pipeline('question-answering', model=model_name, tokenizer=model_name, trust_remote_code=True)
# QA_input = {
#     'question': 'Any future guidance on Q1 2024?',
#     'context': text
# }
# res = nlp(QA_input)
# print(res)
#
# QA_input2 = {
#     'question': 'The company name in question?',
#     'context': text
# }
# res = nlp(QA_input2)
# print(res)
#
# QA_input3 = {
#     'question': "What's the future prospect on Neutron rocket?",
#     'context': text
# }
# res = nlp(QA_input3)
# print(res)

# b) Load model & tokenizer - for fine - tuning
# model = AutoModelForQuestionAnswering.from_pretrained(model_name)
# tokenizer = AutoTokenizer.from_pretrained(model_name)

summary_model_name = "facebook/bart-large-cnn"
summary_nlp = pipeline("summarization", model=summary_model_name)

d = summary_nlp(text, truncation=True, max_length=130, min_length=30)
print(d)

# def chunk_text(text, max_length):
#     # Initialize the tokenizer
#     tokenizer = AutoTokenizer.from_pretrained("facebook/bart-large-cnn")
#
#     # Tokenize the text and split into chunks of max_length
#     tokens = tokenizer.encode(text, add_special_tokens=False)
#     chunk_size = max_length - 2  # account for special tokens [CLS] and [SEP]
#     chunks = [tokens[i:i + chunk_size] for i in range(0, len(tokens), chunk_size)]
#
#     # Decode tokens back to text
#     chunked_texts = [tokenizer.decode(chunk, clean_up_tokenization_spaces=True) for chunk in chunks]
#     return chunked_texts
#
#
# def summarize_chunks(chunks):
#     # Initialize the pipeline
#     summary_nlp = pipeline("summarization", model="facebook/bart-large-cnn")
#
#     # Summarize each chunk
#     summaries = [summary_nlp(chunk)[0]['summary_text'] for chunk in chunks]
#     return " ".join(summaries)
#
#
# chunks = chunk_text(text, 1024)
# summary = summarize_chunks(chunks)
# print(summary)

