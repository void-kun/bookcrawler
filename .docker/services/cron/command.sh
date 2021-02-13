#!/bin/bash

cd cron
apt-get update
apt-get upgrade -y
apt-get install -y gcc
pip install -r requirements.txt
