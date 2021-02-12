from flask.cli import FlaskGroup
from src import app, db

from src.models.book import Book
from src.models.author import Author
from src.models.category import Category
from src.models.chapter import Chapter
from src.models.source import Source


cli = FlaskGroup(app)


@cli.command('create_db')
def create_db():
    db.drop_all()
    db.create_all()
    source = Source('wikidich', 'https://wikidich.com/')
    category = Category('Tiên hiệp')
    author = Author('Ngô Thiên')
    book = Book('Ngã Thị Phong Thiên')
    chapter = Chapter(1, 'lời giới thiệu', 'this is content')
    # append relation
    book.chapters.append(chapter)
    book.categories.append(category)
    book.authors.append(author)
    source.books.append(book)

    db.session.add_all([book, chapter, category, author, source])
    db.session.commit()


if __name__ == "__main__":
    cli()
