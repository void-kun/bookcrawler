from src import db


class Category(db.Model):
    __table_args__ = {"schema": "crawl_local"}
    __tablename__ = 'category'

    id = db.Column(db.Integer, nullable=False, primary_key=True)
    # list books
    name = db.Column(db.String(128), unique=True, nullable=False)
    name_slug = db.Column(db.String(255), unique=True, nullable=False)
    book_total = db.Column(db.Integer, nullable=False, default=0)

    def __init__(self, name) -> None:
        self.name = name
        self.name_slug = name

    def __repr__(self) -> str:
        return '<Category %r>' % self.name
