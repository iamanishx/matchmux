package main

/*
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/ipc.h>
#include <sys/shm.h>
#include <time.h>

#define QUEUE_SIZE 10
#define ITEM_SIZE 64

struct shared_queue {
    int front;
    int rear;
    int count;
    char items[QUEUE_SIZE][ITEM_SIZE];
};

key_t get_key() {
    return ftok("shmfile", 65);
}
*/
import "C"
import (
	"fmt"
	"log"
	"time"
	"unsafe"
)

func main() {
	key := C.get_key()

	shmid := C.shmget(key, C.sizeof_struct_shared_queue, 0666)
	if shmid == -1 {
		log.Fatal("Failed to get shared memory segment")
	}

	shmaddr := C.shmat(shmid, nil, 0)
	if uintptr(unsafe.Pointer(shmaddr)) == ^uintptr(0) {
		log.Fatal("Failed to attach to shared memory")
	}

	queue := (*C.struct_shared_queue)(shmaddr)

	fmt.Println("Go Reader started. Reading messages from queue.")
	fmt.Println("Press Ctrl+C to exit.")

	for {
		if queue.count == 0 {
			time.Sleep(500 * time.Millisecond)
			continue
		}

		message := C.GoString(&queue.items[queue.front][0])
		fmt.Printf("Received: '%s'\n", message)

		if message == "EXIT_READER" {
			fmt.Println("Exit command received. Stopping reader.")
			break
		}

		queue.front = (queue.front + 1) % C.QUEUE_SIZE
		queue.count--
	}

	C.shmdt(shmaddr)

	C.shmctl(shmid, C.IPC_RMID, nil)

	fmt.Println("Shared memory segment removed")
}
