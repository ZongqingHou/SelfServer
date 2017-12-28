import tensorflow as tf
import sklearn as sk

from PIL import Image
import numpy as np
import util
import os

OUTPUT_NODE = 4

NUM_CHANNELS = 3

CONV1_DEEP = 32
CONV1_SIZE = 5
CONV2_DEEP = 64
CONV2_SIZE = 5

FC_SIZE = 512

LEARNING_RATE_BASE = 0.01
LEARNING_RATE_DECAY = 0.99
REGULARIZATRION_RATE = 0.0001
MOVING_AVERAGE_DECAY = 0.99
NUM_CHANNELS = 3
TRAINING_STEPS = 100

def process_image(input_tensor, train, regularizer):
	with tf.variable_scope('layer1-conv1', reuse=tf.AUTO_REUSE):
		conv1_weights = tf.get_variable('weight', [CONV1_SIZE, CONV1_SIZE, NUM_CHANNELS, CONV1_DEEP], initializer=tf.truncated_normal_initializer(stddev=0.1))
		conv1_biases = tf.get_variable('bias', [CONV1_DEEP], initializer=tf.constant_initializer(0.0))
		conv1 = tf.nn.conv2d(input_tensor, conv1_weights, strides=[1, 1, 1, 1], padding='SAME')
		relu1 = tf.nn.relu(tf.nn.bias_add(conv1, conv1_biases))
	
	with tf.name_scope('layer2-pool1'):
		pool1 = tf.nn.max_pool(relu1, ksize=[1, 2, 2, 1], strides=[1, 2, 2, 1], padding='SAME')

	with tf.variable_scope('layer3-conv2', reuse=tf.AUTO_REUSE):
		conv2_weights = tf.get_variable('weight', [CONV2_SIZE, CONV2_SIZE, CONV1_DEEP, CONV2_DEEP], initializer=tf.truncated_normal_initializer(stddev=0.1))
		conv2_biases = tf.get_variable('bias', [CONV2_DEEP], initializer=tf.constant_initializer(0.0))
		conv2 = tf.nn.conv2d(pool1, conv2_weights, strides=[1, 1, 1, 1], padding='SAME')
		relu2 = tf.nn.relu(tf.nn.bias_add(conv2, conv2_biases))

	with tf.name_scope('layer4-pool2'):
		pool2 = tf.nn.max_pool(relu2, ksize=[1, 2, 2, 1], strides=[1, 2, 2, 1], padding='SAME')
		pool_shape = pool2.get_shape().as_list()
		nodes = pool_shape[1] * pool_shape[2] * pool_shape[3]
		reshaped = tf.reshape(pool2, [pool_shape[0], nodes])

	with tf.variable_scope('layer5-fc1', reuse=tf.AUTO_REUSE):
		fc1_weights = tf.get_variable("weight", [nodes, FC_SIZE], initializer=tf.truncated_normal_initializer(stddev=0.1))
		if regularizer != None: tf.add_to_collection('losses', regularizer(fc1_weights))
		fc1_biases = tf.get_variable("bias", [FC_SIZE], initializer=tf.constant_initializer(0.1))

		fc1 = tf.nn.relu(tf.matmul(reshaped, fc1_weights) + fc1_biases)
		if train: fc1 = tf.nn.dropout(fc1, 0.5)

	with tf.variable_scope('layer6-fc2', reuse=tf.AUTO_REUSE):
		fc2_weights = tf.get_variable("weight", [FC_SIZE, OUTPUT_NODE], initializer=tf.truncated_normal_initializer(stddev=0.1))
		if regularizer != None: tf.add_to_collection('losses', regularizer(fc2_weights))
		fc2_biases = tf.get_variable("bias", [OUTPUT_NODE], initializer=tf.constant_initializer(0.1))
		logit = tf.matmul(fc1, fc2_weights) + fc2_biases
	
	return logit

def train(length, width, image, check):
	x = tf.placeholder(tf.float32, [1, width, length, NUM_CHANNELS], name='x-input')
	y_ = tf.placeholder(tf.float32, [None, 4], name='y-input')

	regularizer = tf.contrib.layers.l2_regularizer(REGULARIZATRION_RATE)

	y = process_image(x, False, regularizer)

	global_step = tf.Variable(0, trainable=False)

	variable_averages = tf.train.ExponentialMovingAverage(MOVING_AVERAGE_DECAY, global_step)
	variable_averages_op = variable_averages.apply(tf.trainable_variables())
	cross_entropy = tf.nn.sparse_softmax_cross_entropy_with_logits(logits=y, labels=tf.argmax(y_, 1))
	cross_entropy_mean = tf.reduce_mean(cross_entropy)
	loss = cross_entropy_mean + tf.add_n(tf.get_collection('losses'))
	learning_rate = tf.train.exponential_decay(LEARNING_RATE_BASE, global_step, TRAINING_STEPS, LEARNING_RATE_DECAY, staircase=True)

	train_step = tf.train.GradientDescentOptimizer(learning_rate).minimize(loss, global_step=global_step)

	with tf.control_dependencies([train_step, variable_averages_op]):
		train_op = tf.no_op(name='train')

	saver = tf.train.Saver()

	with tf.Session() as sess:
		if os.path.exists('path/'):
			saver.restore(sess, 'path/train.ckpt')
		else:
			tf.global_variables_initializer().run()

		for step in range(TRAINING_STEPS):
			test, loss_value, step = sess.run([train_op, loss, global_step], feed_dict={x: image, y_: check})
			util.view_bar("processing step of " , step, 100)

		saver.save(sess, 'path/train.ckpt')
		return sess.run(y, feed_dict={x: image})

if __name__ == '__main__':
	image = Image.open('hello.jpg')
	length, width = image.size

	input_data = np.array(image)
	np.reshape(input_data, (length, width, 3))

	# train(length, width, [input_data])