#!/bin/bash

apt-get update
apt-get upgrade -y
pip install -r requirements.txt
# gunicorn --bind 0.0.0.0:3000 manage:app
flask run --host=0.0.0.0 --port=3000
