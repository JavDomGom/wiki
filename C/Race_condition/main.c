#include <pthread.h>
#include <stdio.h>

int counter = 0;
pthread_mutex_t mutex = PTHREAD_MUTEX_INITIALIZER;

void *count ()
{
    int i;

    for (i = 0;i<10;i++)
    {
        pthread_mutex_lock (&mutex);
        counter++;
        printf ("%d ", counter);
        pthread_mutex_unlock (&mutex);
    }
}

int main ()
{
    pthread_t th1, th2;
    pthread_create (&th1, NULL, &count, NULL);
    pthread_create (&th2, NULL, &count, NULL);
    pthread_join (th1, NULL);
    pthread_join (th2, NULL);
    printf ("\nCounter %d\n", counter);
    return 0;
}