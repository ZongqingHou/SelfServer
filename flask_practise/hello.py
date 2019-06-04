from flask import Flask, Response, redirect

import tensorflow as tf
import numpy as np
import sklearn as sk
import pymysql
import re
import glob

import sys
import os
import random
import datetime

app = Flask(__name__)

@app.route('/')
def to_git():
	return redirect('https://github.com/ZongqingHou')

@app.errorhandler(404)
def to_stack():
	return redirect('https://stackoverflow.com/users/5478751/hdd')

@app.route('/ai/veri/<number>', methods=['GET'])
def ai(number):
	with connection.cursor() as cursor:
		sql = 'select * from CHECK_VERIFICATION order by req_time desc limit 1'
		cursor.execute(sql)
		connection.commit()
		row = cursor.fetchone()
	return str(row['act_value'] == number)

@app.route('/image/<type>', methods=['GET'])
def create_cha(type):
	if type == None:
		return 'Come on! Let me give you some challenge'
	elif type == 'verification':
		import ecode
		path, number = ecode.create_img()
		trans = open(path, 'rb')
		now = datetime.datetime.now()
		now.strftime('%Y-%m-%d %H:%M:%S')

		with connection.cursor() as cursor:
			sql = 'insert into CHECK_VERIFICATION (req_time, img_path, act_value) values(%s, %s, %s)'
			cursor.execute(sql, (now, path, number))
			connection.commit()
		return Response(trans, mimetype="image/jpeg")

if __name__ == '__main__':
	connection = pymysql.connect(host='localhost',
								 port=3306,
								 user='root',
								 password='kevinwill21hand',
								 db='AI',
								 cursorclass=pymysql.cursors.DictCursor)
	app.run(host='0.0.0.0')
	connection.close()