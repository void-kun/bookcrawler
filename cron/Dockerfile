# pull official base image
FROM python:3.8.1-slim-buster

# set work directory
WORKDIR /var/www/cron

COPY . .

RUN touch job-cron.log

# set environment variables
ENV PYTHONDONTWRITEBYTECODE 1
ENV PYTHONUNBUFFERED 1

# install dependencies
RUN pip install --upgrade pip

RUN apt-get update &&\
    apt-get upgrade -y &&\
    apt-get install -y gcc cron &&\
    apt-get install htop

RUN pip install -r requirements.txt

# cron.sh
RUN chmod +x cron.sh schedule.sh

ENTRYPOINT ./schedule.sh
