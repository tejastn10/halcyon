# Halcyon 🔍

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/tejastn10/halcyon)

Halcyon is a Go-based CLI tool designed to identify and manage duplicate files in a directory, including nested subdirectories. With Halcyon, you can quickly detect duplicate files, view their paths, and take appropriate actions like moving, deleting, or listing them—all while ensuring a lightweight and fast execution.

## Features

- [x] Traverse directories, including nested subdirectories.
- [x] Identify duplicate files by name and size.
- [x] Detect similar file names (e.g., `file_copy`, `file (1)`).
- [x] Provide multiple options for managing duplicates:
  - Move duplicates to a specific folder.
  - Print paths of duplicates.
  - Delete duplicates.
- [x] Simple and intuitive CLI interface.

---

## Getting Started

### Prerequisites

- [Go 1.23+](https://go.dev/doc/install) installed on your machine.
- [Git](https://git-scm.com/) for version control (optional).

---

### Installation

1. Clone this repository:

    ```bash
    git clone https://github.com/tejastn10/halcyon.git
    cd halcyon
    ```

2. Install dependencies:

    ```bash
    go mod tidy
    ```

3. Run the project:

    ```bash
    go run main.go
    ```

---

### Usage

1. **Build the project:**

    ```bash
    go build -o halcyon
    ```

2. **Run the CLI tool:**

    You can pass the directory path as a flag (defaults to the current directory if no path is provided):

    ```bash
    ./halcyon --dir="/your/directory/path"
    ```

3. **Available actions:**

    After detecting duplicates, Halcyon will provide options:
    - Move duplicates to a folder.
    - Print duplicate file paths.
    - Delete duplicates.

---

### Example Output

```markdown
Total files: 2
Processed: 2
Skipped: 0
+-------+-----------------------+---------+---------------------+
| INDEX |         PATH          |  SIZE   |      MODIFIED       |
+-------+-----------------------+---------+---------------------+
|     0 | duplicates/file.txt   | 0.00 MB | 2024-12-31 22:21:48 |
|     1 | duplicates/file 1.txt | 0.00 MB | 2024-12-31 22:21:43 |
+-------+-----------------------+---------+---------------------+
Use the arrow keys to navigate: ↓ ↑ → ← 
? Choose action: 
  ▸ Keep all files
    Delete specific files
    Delete all duplicates
    Move files to backup
↓   Skip these duplicates
```

### Project Structure

```bash
halcyon/
├── cmd/                  # CLI commands
│   ├── root.go           # Main CLI entry point
├── tasks/                # Core logic for file operations
│   ├── traverse.go       # Directory traversal logic
├── utils/                # Utility functions
│   ├── utils.go          # Reusable helpers
├── go.mod                # Go module definition
├── go.sum                # Go module checksum
├── main.go               # Application entry point
├── README.md             # Project documentation
├── LICENSE.md            # Project license
```

### Contributing

Contributions are welcome! Feel free to open an issue or submit a pull request if you have ideas to enhance Halcyon.

### To-Do

- Add support for advanced duplicate detection using file hashes.
- Allow configurable file actions via a config file.
- Improve performance for large directories.

---

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

## Acknowledgments

- Named after **Halcyon**, symbolizing tranquility and clarity.
- Built with ❤️ and Go.
