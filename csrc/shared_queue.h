#ifndef SHARED_QUEUE_H
#define SHARED_QUEUE_H

#define QUEUE_SIZE 10
#define ITEM_SIZE 64

// Shared memory queue structure
struct shared_queue {
    int front;
    int rear;
    int count;
    char items[QUEUE_SIZE][ITEM_SIZE];
};

#endif