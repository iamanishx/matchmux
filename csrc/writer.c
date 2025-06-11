#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/ipc.h>
#include <sys/shm.h>
#include <sys/stat.h> 
#include <fcntl.h>    
#include <unistd.h>   
#include <errno.h>   

#include "shared_queue.h"

#define FTOK_FILE "shmfile" 
#define FTOK_PROJ_ID 65     

int main() {

    int fd = open(FTOK_FILE, O_CREAT | O_RDWR, 0666);
    if (fd == -1) {
        perror("Failed to create/open ftok file");
        exit(1);
    }
    close(fd); 

    key_t key = ftok(FTOK_FILE, FTOK_PROJ_ID);
    if (key == -1) {
        perror("ftok failed");
        exit(1);
    }

    int shmid = shmget(key, sizeof(struct shared_queue), IPC_CREAT | IPC_EXCL | 0666);
    int is_first_process = (shmid != -1); 

    if (shmid == -1) {
        if (errno == EEXIST) {
            shmid = shmget(key, sizeof(struct shared_queue), 0666);
            if (shmid == -1) {
                perror("shmget failed to get existing segment");
                exit(1);
            }
            printf("Attached to existing shared memory segment.\n");
        } else {
            perror("shmget failed to create segment");
            exit(1);
        }
    } else {
        printf("Created new shared memory segment.\n");
    }
    struct shared_queue *queue = (struct shared_queue*)shmat(shmid, NULL, 0);
    if (queue == (void*) -1) { 
        perror("shmat failed");
        exit(1);
    }

    if (is_first_process) {
        queue->front = 0;
        queue->rear = 0;
        queue->count = 0;

        pthread_mutexattr_t mutex_attr;
        pthread_mutexattr_init(&mutex_attr);
        pthread_mutexattr_setpshared(&mutex_attr, PTHREAD_PROCESS_SHARED);
        pthread_mutex_init(&queue->mutex, &mutex_attr);
        pthread_mutexattr_destroy(&mutex_attr);

        pthread_condattr_t cond_attr;
        pthread_condattr_init(&cond_attr);
        pthread_condattr_setpshared(&cond_attr, PTHREAD_PROCESS_SHARED);
        pthread_cond_init(&queue->cond_not_empty, &cond_attr);
        pthread_cond_init(&queue->cond_not_full, &cond_attr);
        pthread_condattr_destroy(&cond_attr);

        printf("Shared queue and synchronization primitives initialized.\n");
    } else {
        sleep(1);
        printf("Using already initialized shared queue.\n");
    }

    printf("C Writer started. Enter messages to add to the queue (type 'exit' to quit).\n");
    
    char buffer[ITEM_SIZE];
    while (1) {
        printf("Enter message: ");
        if (fgets(buffer, ITEM_SIZE, stdin) == NULL) {
            break; 
        }
        buffer[strcspn(buffer, "\n")] = 0;
        
        if (strcmp(buffer, "exit") == 0) {
            printf("Exiting C writer...\n");
            break;
        }

        pthread_mutex_lock(&queue->mutex);

        while (queue->count == QUEUE_SIZE) {
            printf("C Writer: Queue full, waiting for space...\n");
            pthread_cond_wait(&queue->cond_not_full, &queue->mutex);
        }
        
        strncpy(queue->items[queue->rear], buffer, ITEM_SIZE - 1);
        queue->items[queue->rear][ITEM_SIZE - 1] = '\0'; 
        
        queue->rear = (queue->rear + 1) % QUEUE_SIZE;
        queue->count++;
        
        printf("C Writer: Added '%s'. Queue count: %d\n", buffer, queue->count);

        pthread_cond_signal(&queue->cond_not_empty);

        pthread_mutex_unlock(&queue->mutex);
    }

    shmdt(queue);
    printf("C Writer: Detached from shared memory.\n");

    // Important: Do NOT remove the shared memory segment here (IPC_RMID).
    // This segment is meant to be shared by multiple processes.
    // Cleanup (e.g., shmctl with IPC_RMID) should be done by a dedicated
    // cleanup utility or when all processes have detached and the system
    // is shutting down.

    return 0;
}