import scrapy
from re import search

class WikidichBookSpider(scrapy.Spider):
    name = 'wikidich_list_book'
    allowed_domains = 'wikidich.com'

    def start_requests(self):
        urls = []
        for num in range(0, 980, 20):
            urls.append(
                f'https://wikidich.com/tim-kiem?status=5794f03dd7ced228f4419192&qs=1&gender=5794f03dd7ced228f4419196&m=2&start={num}&so=4&y=2021&vo=1')
        for url in urls:
            yield scrapy.Request(url=url, callback=self.parser)

    def parser(self, response):
        for book in response.css('.book-item'):
            addNum = book.css('.info-col .book-stats-box .book-stats')[0].css('span::text').get()
            if search(r'[km]', addNum) or search(r'(\d){4,}',addNum):
                yield {
                    'title': book.css('.info-col .book-title::text').extract_first(),
                    'url': book.css('.info-col .tooltipped::attr(href)').extract_first(),
                    'addNum': addNum,
                    'author': book.css('.book-author a::text').extract_first(),
                    'gender': book.css('.book-gender::text').extract_first()
                }
