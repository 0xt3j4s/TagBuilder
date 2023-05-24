# from textblob import TextBlob
import nltk
nltk.download('punkt')
nltk.download('stopwords')
from flask import Flask, request, jsonify
from nltk.tokenize import word_tokenize
from nltk.corpus import stopwords
from flask_cors import CORS, cross_origin

app = Flask(__name__)
CORS(app)


@app.route("/generate-tags/", methods=['POST'])
def analyse_sentence():
    sentence = request.get_json()['sentence']
    tags = extract_tags(sentence)
    return jsonify(
        tags=tags
    )

def extract_tags(sentence):
    # using Nltk for extracting tags from the sentence
    words = word_tokenize(sentence)
    stop_words = set(stopwords.words('english'))
    words = [w for w in words if not w in stop_words]
    
    tags = []

    for w, pos in nltk.pos_tag(words):
        if pos.startswith('NN'):
            tags.append(w)

    return tags



if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)
