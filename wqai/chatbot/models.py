from transformers import AutoModelForQuestionAnswering, AutoTokenizer
from transformers import pipeline

# Using models
qa_model_name = "deepset/roberta-base-squad2"
summary_model_name = "facebook/bart-large-cnn"

# Question - Answering based on context
print("------   Q-A model loading   ------")
qa_nlp = pipeline(
    'question-answering',
    model=qa_model_name,
    tokenizer=qa_model_name,
    trust_remote_code=True
)

# Summary
print("------ Summry model loading  ------")
summary_nlp = pipeline("summarization", model=summary_model_name)

# summarizer(ARTICLE, max_length=130, min_length=30, do_sample=False)
# print()
# >>> [{'summary_text': 'Liana Barrientos, 39, is charged with two counts of "offering a false instrument
# for filing in the first degree" In total, she has been married 10 times, with nine of her marriages occurring
# between 1999 and 2002. She is believed to still be married to four men.'}]
