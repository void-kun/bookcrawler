from datetime import datetime

from src import db

book_author = db.Table(
    'book_author', db.Model.metadata,
    db.Column('book_id', db.Integer, db.ForeignKey(
        'book.id'), primary_key=True),
    db.Column('author_id', db.Integer, db.ForeignKey(
        'author.id'), primary_key=True)
)

book_category = db.Table(
    'book_category', db.Model.metadata,
    db.Column('book_id', db.Integer, db.ForeignKey(
        'book.id'), primary_key=True),
    db.Column('category_id', db.Integer, db.ForeignKey(
        'category.id'), primary_key=True),
    db.Column('update_at', db.DateTime,
              nullable=False, default=datetime.utcnow),
    db.Column('create_at', db.DateTime,
              nullable=False, default=datetime.utcnow)
)


class Book(db.Model):
    __tablename__ = 'book'

    id = db.Column(db.Integer, primary_key=True)
    name = db.Column(db.String(128), unique=True, nullable=False)
    name_slug = db.Column(db.String(255), unique=True, nullable=False)
    image_url = db.Column(db.String(255))
    source_id = db.Column(db.Integer, db.ForeignKey('source.id'),
                          nullable=False)
    # author relation
    authors = db.relationship('Author', secondary=book_author)
    # chapter relation
    chapters = db.relationship('Chapter', lazy='select',
                               backref=db.backref('book', lazy='joined'))
    # category relation
    categories = db.relationship('Category', secondary=book_category)

    def __init__(self, name) -> None:
        self.name = name
        self.name_slug = name

    def __repr__(self) -> str:
        return '<Book %r>' % self.name
