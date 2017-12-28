from flask import Flask
import tensorflow as tf
import numpy as np
import sklearn as sk
import pymysql

app = Flask(__name__)

@app.route('/')
def hello_world():
    return 'Hello World!'

@app.route('/user')
def hello():
	return 'wow'

@app.route('/sk')
def main_sk:
	pass

@app.route('/tf')
def main_tf:
	pass

if __name__ == '__main__':
    app.run(host='0.0.0.0')
