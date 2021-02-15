from src import db


class Source(db.Model):
    __table_args__ = {"schema": "crawl_local"}
    __tablename__ = 'source'

    id = db.Column(db.Integer, nullable=False, primary_key=True)
    books = db.relationship('Book', backref=db.backref('source', lazy=True))
    name = db.Column(db.String(128), unique=True, nullable=False)
    name_slug = db.Column(db.String(255), unique=True, nullable=False)
    url = db.Column(db.String(255), unique=True, nullable=False)
    book_total = db.Column(db.Integer, nullable=False, default=0)

    def __init__(self, name, url) -> None:
        self.name = name
        self.name_slug = name
        self.url = url

    def __repr__(self) -> str:
        return '<Source %r>' % self.name
