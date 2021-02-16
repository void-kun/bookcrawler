#!/bin/bash

cd cron
apt-get update
apt-get upgrade -y
apt-get install -y gcc cron
pip install -r requirements.txt

chmod +x /var/www/app/.docker/services/cron/cron.sh
crontab -l | { cat; echo "*/10 * * * * /var/www/app/.docker/services/cron/cron.sh"; } | crontab -

tail -F /dev/null
