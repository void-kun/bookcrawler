import psycopg2
import os

from cron.items import BookItem
from cron.helper import convert_date


class CronPipeline:

    def open_spider(self, spider):
        hostname = 'db'
        username = 'zrik'
        password = 'secret'
        database = 'bookcrawler'
        self.connection = psycopg2.connect(
            host=hostname, user=username, password=password, dbname=database)
        self.cur = self.connection.cursor()

        self.create_db()

    def close_spider(self, spider):
        self.cur.close()
        self.connection.close()

    def process_item(self, item: BookItem, spider):
        try:
            query = f"""
                INSERT INTO raw_local.craw_book_raw(
                    title, url, avatar_url, visibility, author, state, 
                    last_chapter_title, last_chapter_at, 
                    categories, summary) 
                values('{item['title']}', '{item['url']}', '{item['avatar_url']}',
                {item['visibility']}, '{item['author']}', '{item['state']}',
                '{item['last_chapter_title']}', '{convert_date(item['last_chapter_at'])}',
                ARRAY [{','.join(item['categories'])}], '{item['summary']}')
                ON CONFLICT (url) DO UPDATE 
                SET visibility = excluded.visibility, 
                    state = excluded.state,
                    last_chapter_title = last_chapter_title,
                    last_chapter_at = last_chapter_at;
            """
            self.cur.execute(query)
            self.connection.commit()
        except Exception as e:
            with open('error.log', 'a') as f:
                f.write(str(e))
            self.connection.rollback()
        return item

    def create_db(self):
        if self.cur:
            # create schema
            schema_create = """
                CREATE SCHEMA IF NOT EXISTS raw_local;
            """
            table_create = """
            CREATE TABLE IF NOT EXISTS raw_local.craw_book_raw (
                id int PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
                title varchar(255) NOT NULL,
                url varchar(255) NOT NULL,
                avatar_url varchar(255) UNIQUE NOT NULL,
                visibility int NOT NULL,
                author varchar(128) NOT NULL,
                state varchar(30) NOT NULL,
                last_chapter_title varchar(255),
                last_chapter_at date,
                categories varchar(30) ARRAY,
                summary TEXT
            );    
            """
            self.cur.execute(schema_create)
            self.cur.execute(table_create)
            self.connection.commit()
