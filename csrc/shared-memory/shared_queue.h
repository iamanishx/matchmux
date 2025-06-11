#ifndef SHARED_QUEUE_H
#define SHARED_QUEUE_H

#include <pthread.h> 

#define QUEUE_SIZE 10  
#define ITEM_SIZE 64   

struct shared_queue {
    int front;                           
    int rear;                            
    int count;                          
    char items[QUEUE_SIZE][ITEM_SIZE];   

    pthread_mutex_t mutex;               // Mutex to protect shared_queue access
    pthread_cond_t cond_not_empty;       // Condition variable for consumers (queue not empty)
    pthread_cond_t cond_not_full;        // Condition variable for producers (queue not full)
};

#endif 