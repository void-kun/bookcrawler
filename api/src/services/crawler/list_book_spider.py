import scrapy


class ListBookSpider(scrapy.Spider):
    name = 'list_book'
    allowed_domains = 'wikidich.com'

    def start_requests(self):
        urls = []
        for num in range(0, 0):
            urls.append(f'https://wikidich.com/tim-kiem?status=5794f03dd7ced228f4419192&qs=1&gender=5794f03dd7ced228f4419196&m=2&start={num}&so=4&y=2021&vo=1')   
        for url in urls:
            yield scrapy.Request(url=url, callback=self.parser)

    def parser(self, response):
        print(response.content)
