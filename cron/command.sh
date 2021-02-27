#!/bin/bash

apt-get update
apt-get upgrade -y
apt-get install -y gcc cron
pip install -r requirements.txt

chmod +x /var/www/cron/cron.sh
crontab -l | { cat; echo "*/10 * * * * /var/www/cron/cron.sh"; } | crontab -

tail -F /dev/null
