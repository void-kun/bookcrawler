#!/bin/bash

cd api
pip install -r requirements.txt
gunicorn --bind 0.0.0.0:3000 manage:app
