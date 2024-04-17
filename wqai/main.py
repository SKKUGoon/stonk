from transformers import AutoModelForQuestionAnswering, AutoTokenizer
from transformers import pipeline

from text_data import clean_text

# Text
text = clean_text("conference_call_rklb.txt")

model_name = "deepset/roberta-base-squad2"

# a) Get predictions
nlp = pipeline('question-answering', model=model_name, tokenizer=model_name, trust_remote_code=True)
QA_input = {
    'question': 'Any future guidance on Q1 2024?',
    'context': text
}
res = nlp(QA_input)
print(res)

QA_input2 = {
    'question': 'The company name in question?',
    'context': text
}
res = nlp(QA_input2)
print(res)

QA_input3 = {
    'question': "What's the future prospect on Neutron rocket?",
    'context': text
}
res = nlp(QA_input3)
print(res)

# b) Load model & tokenizer
model = AutoModelForQuestionAnswering.from_pretrained(model_name)
tokenizer = AutoTokenizer.from_pretrained(model_name)
