import json
import codecs
from scrapy import Spider, Request
from scrapy_selenium import SeleniumRequest
from selenium.webdriver.common.by import By
from selenium.webdriver.support import expected_conditions as EC

from cron.items import ChapterItem
from cron.db_connect import BookDB


class WikidichChapSpider(Spider):
    name = 'wikidich_book_chap'
    allowed_domains = ['wikidich.com']
    custom_settings = {
        'ITEM_PIPELINES': {'cron.pipelines.WKBookChapPipeline': 300},
    }

    def start_requests(self):
        db = BookDB()
        book_list = db.get_book_reading()

        for book in book_list:
            try:
                yield SeleniumRequest(
                    url=book[1], 
                    callback=self.book_parser,
                    wait_time=3000, 
                    meta={'book_id': book[0]},
                    wait_until=EC.presence_of_element_located((By.CSS_SELECTOR, '.chapter-name')),
                )
            except Exception as e:
                print('\n\n\n', e.message, '\n\n')

    def book_parser(self, response):
        json_data = json.loads(response.body)

        for data in json_data['content']['results']:
            print (data, ' ')

        for section in response.xpath('//div[@class="volume-list"]'):
            for li in section.xpath('.//li'):
                url=f"https://wikidich.com{li.xpath('.//a/@href').get()}"
                yield Request(url=url, callback=self.chap_parser, meta=response.meta)
            

    def chap_parser(self, response):

        chapter = ChapterItem()
        chapter['book_id'] = response.meta['book_id']
        chapter['url'] = response.url
        chapter['title'] = chapter.css('.book-title:nth-child(2)').get()
        chapter['content'] = chapter.css('#bookContentBody::text').get()

        yield chapter
