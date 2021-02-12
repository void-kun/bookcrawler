from datetime import datetime

from src import db


class Chapter(db.Model):
    __tablename__ = 'chapter'

    id = db.Column(db.Integer, nullable=False, primary_key=True)
    book_id = db.Column(db.Integer, db.ForeignKey('book.id'), nullable=False)
    chap_number = db.Column(db.Integer, nullable=False, default=0)
    title = db.Column(db.Text, nullable=False)
    title_slug = db.Column(db.String(255), unique=True, nullable=False)
    content = db.Column(db.Text, nullable=False)
    update_at = db.Column(db.DateTime, nullable=False,
                          default=datetime.utcnow)
    create_at = db.Column(db.DateTime, nullable=False,
                          default=datetime.utcnow)

    def __init__(self, chap_number, title, content) -> None:
        self.chap_number = chap_number
        self.title = title
        self.title_slug = title
        self.content = content

    def __repr__(self) -> str:
        return '<Chapter %r>' % self.name
