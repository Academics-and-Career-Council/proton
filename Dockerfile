
FROM python:3.10.8-slim-bullseye
RUN mkdir /ocr
WORKDIR /ocr
COPY req.txt .
RUN pip3 install -r req.txt
COPY . /ocr
EXPOSE 7000
ENTRYPOINT [ "gunicorn", "server:app", "-w", "1", "-b", "0.0.0.0:7000" ]