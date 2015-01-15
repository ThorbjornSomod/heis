#include <pthread.h>
#include <stdio.h>

	int i;

void* function1(){
	int x;
	for (x = 0; x<1000000; x++){
	i++;}
	
	return NULL;
}

void* function2(){
	int x;
	for (x = 0; x<1000000; x++){
	i--;}
	
	return NULL;
}


int main(){
	i = 0;
	pthread_t thread1;
	pthread_create(&thread1, NULL,function1,NULL );
	pthread_t thread2;
	pthread_create(&thread2, NULL,function2,NULL );
	pthread_join(thread1,NULL);
	pthread_join(thread2,NULL);

	printf("%d\n", i);
	return 0;
}