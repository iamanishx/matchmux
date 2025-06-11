
### 🔧 Build Targets

| Command           | Description                             |
| ----------------- | --------------------------------------- |
| `make`            | Build **all** C and Go components       |
| `make c_programs` | Build only the **C programs**           |
| `make go_program` | Build only the **Go program**           |
| `make clean`      | Remove all built binaries and artifacts |

---

### 🚀 Run Targets

| Command              | Description                                       |
| -------------------- | ------------------------------------------------- |
| `make run_writer`    | Run the **C writer** (sends messages to queue)    |
| `make run_reader`    | Run the **C reader** (reads messages from queue)  |
| `make run_go_reader` | Run the **Go reader** (reads messages from queue) |

---

### 📄 Utility

| Command        | Description                          |
| -------------- | ------------------------------------ |
| `make shmfile` | Create `shmfile` if it doesn’t exist |

