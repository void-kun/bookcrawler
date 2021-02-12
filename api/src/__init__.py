from flask import Flask, jsonify
from flask_sqlalchemy import SQLAlchemy

app = Flask(__name__)
app.config.from_object('src.config.Config')
db = SQLAlchemy(app)


@app.route("/")
def hello_world():
    return jsonify(hello="dd")
