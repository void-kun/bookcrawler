from flask import Flask, jsonify
from flask_sqlalchemy import SQLAlchemy
from sqlalchemy import create_engine

app = Flask(__name__)
app.config.from_object('src.config.Config')
engine = create_engine(app.config['SQLALCHEMY_DATABASE_URI'])
engine.execute('CREATE SCHEMA IF NOT EXISTS crawl_local;')
db = SQLAlchemy(app)

@app.route("/")
def hello_world():
    return jsonify(hello="dd")
