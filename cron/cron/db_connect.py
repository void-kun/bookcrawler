import psycopg2
import os


class DB(object):

    def __init__(self):
        self.hostname = 'db'
        self.username = os.getenv('POSTGRES_USER')
        self.password = os.getenv('POSTGRES_PASSWORD')
        self.database = os.getenv('POSTGRES_DB')
        self.connection = None

    def open_db(self):
        self.connection = psycopg2.connect(
            host=self.hostname, user=self.username, password=self.password, dbname=self.database)

    def close_db(self):
        self.connection.close()


class BookDB(DB):

    def get_book_reading(self):
        try:
            self.open_db()
            with self.connection.cursor() as cursor:
                query = "SELECT id, url FROM raw_local.craw_book_raw WHERE id = 2;"
                cursor.execute(query)
                return cursor.fetchall()
        except Exception:
            self.connection.rollback()
        finally:
            if self.connection:
                self.close_db()
        return None
