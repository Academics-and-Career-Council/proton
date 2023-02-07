
# FROM python:3.10.8-slim-bullseye
# RUN mkdir /ocr
# WORKDIR /ocr
# COPY req.txt .
# RUN pip3 install -r req.txt
# COPY . /ocr
# EXPOSE 7000
# ENTRYPOINT [ "gunicorn", "server:app", "-w", "1", "-b", "0.0.0.0:7000" ]

# Preparing environment
FROM ubuntu:22.04
RUN apt-get update 
RUN apt install software-properties-common apt-utils python3-pip -y
RUN apt install -y default-jre
RUN apt install -y default-jdk

# Working with code
RUN mkdir /ocr
WORKDIR /ocr
COPY req.txt .
RUN pip3 install -r req.txt
COPY . /ocr
EXPOSE 7000
# CMD ["python3","-u","server.py"]
ENTRYPOINT [ "gunicorn", "server:app", "-w", "1", "-b", "0.0.0.0:7000" ]
