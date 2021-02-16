#!/bin/bash

cd /var/www/app/cron
scrapy crawl --set FEED_EXPORT_ENCODING=utf-8 wikidich_list_book
