package sharedmemory

/*
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/ipc.h>
#include <sys/shm.h>
#include <sys/stat.h>
#include <fcntl.h>
#include <unistd.h>
#include <errno.h>
#include <pthread.h>

#define QUEUE_SIZE 10
#define ITEM_SIZE 64
#define FTOK_FILE "shmfile"
#define FTOK_PROJ_ID 65

struct shared_queue {
    int front;
    int rear;
    int count;
    char items[QUEUE_SIZE][ITEM_SIZE];

    pthread_mutex_t mutex;
    pthread_cond_t cond_not_empty;
    pthread_cond_t cond_not_full;
};

int get_errno() {
    return errno;
}

const char* get_error_str() {
    return strerror(errno);
}

key_t get_key() {
    int fd = open(FTOK_FILE, O_CREAT | O_RDWR, 0666);
    if (fd == -1) {
        return -1; 
    }
    close(fd);
    key_t k = ftok(FTOK_FILE, FTOK_PROJ_ID);
    return k;
}

int shmget_wrapper(key_t key, size_t size, int shmflg) {
    return shmget(key, size, shmflg);
}

void* shmat_wrapper(int shmid, const void* shmaddr, int shmflg) {
    return shmat(shmid, shmaddr, shmflg);
}

int shmdt_wrapper(const void* shmaddr) {
    return shmdt(shmaddr);
}

int shmctl_wrapper(int shmid, int cmd, struct shmid_ds* buf) {
    return shmctl(shmid, cmd, buf);
}

size_t get_shared_queue_size() {
    return sizeof(struct shared_queue);
}

void lock_mutex(pthread_mutex_t* m) {
    pthread_mutex_lock(m);
}

void unlock_mutex(pthread_mutex_t* m) {
    pthread_mutex_unlock(m);
}

void cond_wait_not_empty(pthread_cond_t* cond, pthread_mutex_t* mutex) {
    pthread_cond_wait(cond, mutex);
}

void cond_signal_not_full(pthread_cond_t* cond) {
    pthread_cond_signal(cond);
}


void read_message(struct shared_queue* q, char* buf) {
    strncpy(buf, q->items[q->front], ITEM_SIZE - 1);
    buf[ITEM_SIZE - 1] = '\0'; // Ensure null termination

    q->front = (q->front + 1) % QUEUE_SIZE;
    q->count--;
}
*/
import "C" 

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
	"unsafe" 
)

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	done := make(chan bool, 1) 

	go func() {
		<-sigs 
		fmt.Println("\nReceived shutdown signal. Initiating graceful exit...")
		done <- true 
	}()

	key := C.get_key()
	if key == -1 {
		log.Fatalf("Failed to get ftok key. Error: %s", C.GoString(C.get_error_str()))
	}

	
	shmid := C.shmget_wrapper(key, C.get_shared_queue_size(), 0666)
	if shmid == -1 {
		if C.get_errno() == C.ENOENT { 
			log.Fatalf("Shared memory segment (key %d) does not exist. Please ensure the C writer is running and has created it. Error: %s", key, C.GoString(C.get_error_str()))
		}
		log.Fatalf("Failed to get shared memory segment (key %d). Error: %s", key, C.GoString(C.get_error_str()))
	}

	shmaddr := C.shmat_wrapper(shmid, nil, 0)
	if uintptr(unsafe.Pointer(shmaddr)) == ^uintptr(0) { 
		log.Fatalf("Failed to attach to shared memory. Error: %s", C.GoString(C.get_error_str()))
	}
	queue := (*C.struct_shared_queue)(shmaddr)
	msgBuf := make([]byte, C.ITEM_SIZE)
	cMsgBuf := (*C.char)(unsafe.Pointer(&msgBuf[0]))

	fmt.Println("Go Reader started. Reading messages from queue. Press Ctrl+C to quit.")

	for {
		select {
		case <-done: 
			fmt.Println("Shutting down Go reader...")
			goto cleanup 
		default:
			C.lock_mutex(&queue.mutex)

			for queue.count == 0 {
				fmt.Println("Go Reader: Queue empty, waiting for messages...")
				C.cond_wait_not_empty(&queue.cond_not_empty, &queue.mutex)
				select {
				case <-done:
					C.unlock_mutex(&queue.mutex) 
					fmt.Println("Shutting down Go reader from wait state...")
					goto cleanup
				default:
				}
			}

			C.read_message(queue, cMsgBuf) 
			message := C.GoString(cMsgBuf) 

			fmt.Printf("Go Reader: Received '%s'. Queue count: %d\n", message, queue.count)

			C.cond_signal_not_full(&queue.cond_not_full)

			C.unlock_mutex(&queue.mutex)

			time.Sleep(10 * time.Millisecond)

			if message == "EXIT_READER" {
				fmt.Println("Exit command received. Stopping reader.")
				goto cleanup 
			}
		}
	}

cleanup:
	
	ret := C.shmdt_wrapper(shmaddr)
	if ret == -1 {
		log.Printf("Failed to detach from shared memory. Error: %s", C.GoString(C.get_error_str()))
	} else {
		fmt.Println("Go Reader: Detached from shared memory.")
	}

	ret = C.shmctl_wrapper(shmid, C.IPC_RMID, nil)
	if ret == -1 {
		log.Printf("Failed to remove shared memory segment. Error: %s", C.GoString(C.get_error_str()))
	} else {
		fmt.Println("Shared memory segment removed.")
	}

	fmt.Println("Go Reader gracefully exited.")
}
