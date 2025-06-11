# Build all components (C and Go)
`make`

# Build only C programs
`make c_programs`

# Build only Go program
`make go_program`

# Clean all built binaries
`make clean`

# Run C writer (sends messages to queue)
`make run_writer`

# Run C reader (reads messages from queue)
`make run_reader`

# Run Go reader (reads messages from queue)
`make run_go_reader`

# Create shmfile if it doesn't exist
`make shmfile`
