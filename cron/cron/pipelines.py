import json
import codecs
import psycopg2


class CronPipeline:

    def open_spider(self, spider):
        hostname = 'db'
        username = 'zrik'
        password = 'secret'
        database = 'bookcrawler'
        self.connection = psycopg2.connect(
            host=hostname, user=username, password=password, dbname=database)
        self.cur = self.connection.cursor()

    def close_spider(self, spider):
        self.cur.close()
        self.connection.close()

    def process_item(self, item, spider):
        self.cur.execute("insert into quotes_content(content,author) values(%s,%s)",
                         (item['content'], item['author']))
        self.connection.commit()
        return item
