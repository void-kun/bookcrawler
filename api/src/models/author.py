from src import db


class Author(db.Model):
    __tablename__ = 'author'

    id = db.Column(db.Integer, nullable=False, primary_key=True)
    # list books
    name = db.Column(db.String(128), unique=True, nullable=False)
    name_slug = db.Column(db.String(255), unique=True, nullable=False)
    pseudonym = db.Column(db.String(255), unique=True, nullable=True)
    book_total = db.Column(db.Integer, nullable=False, default=0)

    def __init__(self, name):
        self.name = name
        self.name_slug = name

    def __repr__(self) -> str:
        return "<Author %r>" % (self.name)
