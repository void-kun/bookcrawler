#!/bin/bash

# Start the run once job.
echo "Docker container has been started"

declare -p | grep -Ev 'BASHOPTS|BASH_VERSINFO|EUID|PPID|SHELLOPTS|UID' > /container.env

# Setup a cron schedule
echo "
SHELL=/bin/bash
BASH_ENV=/container.env

# schedule crawl books wikidich basic
0 */3 * * * /var/www/cron/wk_books_cron.sh
# 0 */3 * * * /var/www/cron/wk_chaps_cron.sh

" > scheduler.txt

crontab scheduler.txt
cron -f