# Compiler settings
CC = gcc
CFLAGS = -Wall -Wextra
GO = go

# Directories
C_SRC_DIR = csrc
GO_SRC_DIR = gosrc

# Targets
C_TARGETS = $(C_SRC_DIR)/shared-memory/writer $(C_SRC_DIR)/shared-memory/reader
GO_TARGET = $(GO_SRC_DIR)/shared-memory/queue_reader

.PHONY: all clean c_programs go_program

all: c_programs go_program

c_programs: $(C_TARGETS)

go_program: $(GO_TARGET)

# C programs
$(C_SRC_DIR)/shared-memory/writer: $(C_SRC_DIR)/shared-memory/writer.c $(C_SRC_DIR)/shared-memory/shared_queue.h
	$(CC) $(CFLAGS) -o $@ $<

$(C_SRC_DIR)/shared-memory/reader: $(C_SRC_DIR)/shared-memory/reader.c $(C_SRC_DIR)/shared-memory/shared_queue.h
	$(CC) $(CFLAGS) -o $@ $<

$(GO_SRC_DIR)/shared-memory/queue_reader: $(GO_SRC_DIR)/shared-memory/queue_reader.go
	cd $(GO_SRC_DIR)/shared-memory && $(GO) build -o queue_reader queue_reader.go

clean:
	rm -f $(C_TARGETS) $(GO_TARGET)

# Run targets
run_writer:
	$(C_SRC_DIR)/shared-memory/writer

run_reader:
	$(C_SRC_DIR)/shared-memory/reader

run_go_reader:
	$(GO_SRC_DIR)/shared-memory/queue_reader
shmfile:
	@if [ ! -f shmfile ]; then \
		echo "Creating shmfile..."; \
		echo "IPC shared file - Do not delete" > shmfile; \
		chmod 644 shmfile; \
		echo "shmfile created successfully."; \
	else \
		echo "shmfile already exists."; \
	fi