from threading import Thread

i = 0

def function1():
	global i
	x=0
	for x in xrange(1,1000000):
		i+=1

def function2():
	global i
	x=0
	for x in xrange(1,1000000):
		i-=1

def main():
	thread1 = Thread(target = function1, args = (),)
	thread1.start()

	thread2 = Thread(target = function2, args = (),)
	thread2.start()

	thread1.join()
	thread2.join()

	print(i)

main()

