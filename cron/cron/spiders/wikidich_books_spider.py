from scrapy import Spider, Request
from re import search

from cron.items import BookItem


class WikidichBookSpider(Spider):
    name = 'wikidich_list_book'
    allowed_domains = ['wikidich.com']

    def start_requests(self):
        for num in range(0, 980, 20):
            yield Request(f'https://wikidich.com/tim-kiem?status=5794f03dd7ced228f4419192&qs=1&gender=5794f03dd7ced228f4419196&m=2&start={num}&so=4&y=2021&vo=1', callback=self.parser)

    def parser(self, response):
        for book in response.css('.book-item'):
            addNum = book.css('.book-stats:first-child span::text').get()
            if search(r'[km]', addNum) or search(r'(\d){4,}', addNum):
                url = f"https://wikidich.com{book.css('.tooltipped::attr(href)').get()}"
                yield Request(url=url, callback=self.book_parser)

    def book_parser(self, response):
        book = BookItem()
        cover_info = response.css('.cover-info')
        book_desc = response.css('.book-desc')
        book['url'] = response.url
        book['avatar_url'] = f"https://wikidich.com{response.css('.book-info img::attr(src)').get()}"
        book['title'] = cover_info.xpath('.//h2//text()').get()
        book['visibility'] = cover_info.xpath('.//p[1]//span[1]//span//text()').get()
        book['author'] = cover_info.xpath('.//p[3]//a//text()').get()
        book['state'] = cover_info.xpath('.//p[4]//a//text()').get()
        book['last_chapter_title'] = cover_info.xpath('.//p[5]//a//text()').get()
        book['last_chapter_at'] = cover_info.xpath('.//p[6]//span//text()').get()
        book['categories'] = [f"\'{i}\'" for i in book_desc.xpath('.//p//span//a//text()').getall()]
        book['summary'] = '\n'.join([i.get() for i in book_desc.css('.book-desc-detail p::text')])
        return book
