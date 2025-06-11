#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <sys/ipc.h>
#include <sys/shm.h>
#include "shared_queue.h"

int main() {
    key_t key = ftok("shmfile", 65);
    int shmid = shmget(key, sizeof(struct shared_queue), 0666);
    if (shmid == -1) {
        perror("shmget");
        exit(1);
    }
    
    struct shared_queue *queue = (struct shared_queue*)shmat(shmid, NULL, 0);
    if (queue == (void*)-1) {
        perror("shmat");
        exit(1);
    }
    
    printf("Reader started. Reading messages from queue.\n");
    printf("Press Ctrl+C to exit.\n");
    
    while (1) {
        if (queue->count == 0) {
            usleep(500000);  
            continue;
        }
        
        char *message = queue->items[queue->front];
        printf("Received: '%s'\n", message);
        
        if (strcmp(message, "EXIT_READER") == 0) {
            printf("Exit command received. Stopping reader.\n");
            break;
        }
        
        queue->front = (queue->front + 1) % QUEUE_SIZE;
        queue->count--;
    }
    
    shmdt(queue);
    
    shmctl(shmid, IPC_RMID, NULL);
    
    return 0;
}