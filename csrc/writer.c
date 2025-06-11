#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/ipc.h>
#include <sys/shm.h>
#include "shared_queue.h"


#define SHM_SIZE 1024

int main() {

    //shmget(): Creates or accesses a shared memory segment.
    key_t key = ftok("shmfile", 65); 
    int shmid = shmget(key, sizeof(struct shared_queue) , IPC_CREAT | 0666);
     if (shmid == -1) {
        perror("shmget");
        exit(1);
    }

    struct shared_queue *queue = (struct shared_queue*)shmat(shmid, NULL, 0);
    if (queue == (char*) -1) {
        perror("shmat");
        exit(1);
    }

    queue -> front = 0;
    queue -> rear = 0;
    queue -> count = 0;

    printf("Reader started. Reading messages from queue.\n");
    printf("Press Ctrl+C to exit.\n");
    
    char buffer[ITEM_SIZE];
    while (1) {
        // Get input from terminal
        printf("> ");
        if (fgets(buffer, ITEM_SIZE, stdin) == NULL) {
            break;
        }
        
        buffer[strcspn(buffer, "\n")] = 0;
        
        if (strcmp(buffer, "exit") == 0) {
            printf("Exiting writer...\n");
            break;
        }
        
        while (queue->count == QUEUE_SIZE) {
            printf("Queue full, waiting...\n");
            sleep(1);
        }
        
        strncpy(queue->items[queue->rear], buffer, ITEM_SIZE-1);
        queue->items[queue->rear][ITEM_SIZE-1] = '\0';
        
        queue->rear = (queue->rear + 1) % QUEUE_SIZE;
        queue->count++;
        
        printf("Added to queue: '%s'\n", buffer);
    }

    shmdt(queue);
    return 0;
}