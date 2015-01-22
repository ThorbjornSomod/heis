from threading import Thread
from threading import Lock 

i = 0
lock = Lock()

def function1():
	global i
	x=0
	for x in xrange(1,1000000):
		lock.acquire()
		i+=1
		lock.release()

def function2():
	global i
	x=0
	for x in xrange(1,1000001):
		lock.acquire()
		i-=1
		lock.release()

def main():
	thread1 = Thread(target = function1, args = (),)
	thread1.start()

	thread2 = Thread(target = function2, args = (),)
	thread2.start()

	thread1.join()
	thread2.join()

	print(i)

main()

