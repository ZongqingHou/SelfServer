from PIL import Image,ImageFont,ImageDraw,ImageFilter
import random
 
def charRandom():
	return chr((random.randint(65,90)))
 
def numRandom():
	return str(random.randint(0,9))
 
def colorRandom1():
	return (random.randint(64,255),random.randint(64,255),random.randint(64,255))
 
def colorRandom2():
	return (random.randint(32,127),random.randint(32,127),random.randint(32,127))

def create_img():
	width = 60 * 4
	height = 60
	image = Image.new('RGB', (width,height), (255,255,255));
	font = ImageFont.truetype('LemonMilk.otf',36);
 
	draw = ImageDraw.Draw(image)
	return_number = ''
	for x in range(width):
		for y in range(height):
			draw.point((x,y), fill=colorRandom1())
		 
	for t in range(4):
		temp = numRandom()
		draw.text((60*t+10,10), temp,font=font, fill=colorRandom2())
		return_number += temp
	
	image = image.filter(ImageFilter.BLUR)
	path = 'hello.jpg'
	image.save(path, 'jpeg')
	return path, return_number