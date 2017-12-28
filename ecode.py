from PIL import Image,ImageFont,ImageDraw,ImageFilter
import random
 
 def charRandom():
    return chr((random.randint(65,90)))
 
def numRandom():
    return random.randint(0,9)
 
def colorRandom1():
    return (random.randint(64,255),random.randint(64,255),random.randint(64,255))
 
def colorRandom2():
    return (random.randint(32,127),random.randint(32,127),random.randint(32,127))
 
width = 60 * 4
height = 60
image = Image.new('RGB', (width,height), (255,255,255));
font = ImageFont.truetype('LemonMilk.otf',36);
 
draw = ImageDraw.Draw(image)
for x in range(width):
    for y in range(height):
        draw.point((x,y), fill=colorRandom1())
         
for t in range(4):
    draw.text((60*t+10,10), charRandom(),font=font, fill=colorRandom2())
 
image = image.filter(ImageFilter.BLUR)
image.save('code.jpg','jpeg')