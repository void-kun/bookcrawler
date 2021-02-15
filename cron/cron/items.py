# Define here the models for your scraped items
#
# See documentation in:
# https://docs.scrapy.org/en/latest/topics/items.html

import scrapy


class BookItem(scrapy.Item):
    url = scrapy.Field()
    title = scrapy.Field()
    avatar_url = scrapy.Field()
    visibility = scrapy.Field()
    author = scrapy.Field()
    state = scrapy.Field()
    last_chapter_title = scrapy.Field()
    last_chapter_url = scrapy.Field()
    last_chapter_at = scrapy.Field()
    categories = scrapy.Field()
    summary = scrapy.Field()
