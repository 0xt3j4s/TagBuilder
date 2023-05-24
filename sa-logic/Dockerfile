FROM python:3.6-slim
COPY sa /app
WORKDIR /app
RUN pip3 install -r requirements.txt && \
    python3 -m nltk.downloader stopwords && \
    python3 -m nltk.downloader punkt && \
    python3 -m nltk.downloader wordnet
EXPOSE 5000
ENTRYPOINT ["python3"]
CMD ["tag_builder.py"]
