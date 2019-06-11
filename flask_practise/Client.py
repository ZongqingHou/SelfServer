from io import BytesIO
from PIL import Image
import numpy as np
import requests
import ai

ai_ = None
flag = True

while True:
	url = 'http://0.0.0.0:5000/image/verification'
	res = requests.get(url)
	image = Image.open(BytesIO(res.content))
	image.show()
	length, width = image.size

	if flag:
		user = input('your chose:\n')

	if user == 'ai':
		url = 'http://0.0.0.0:5000/ai/{}'.format(ai_)
		flag = False
		res = requests.get(url)
		print(res.text)
	else:
		check_list = [[int(temp) for temp in user]]
		input_data = np.array(image)
		np.reshape(input_data, (length, width, 3))
		print(check_list)
		ai_ = ai.train(length, width, [input_data], check_list)
		print(ai_)