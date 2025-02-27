<p align="center">
  <img src="logo.svg" alt="Logo">
</p>

# Halcyon 🔍

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/tejastn10/halcyon?logo=go&logoColor=white)
[![Unit Tests](https://github.com/tejastn10/halcyon/actions/workflows/unit-test.yml/badge.svg)](https://github.com/tejastn10/halcyon/actions/workflows/unit-test.yml)
[![Release Workflow](https://github.com/tejastn10/halcyon/actions/workflows/release.yml/badge.svg)](https://github.com/tejastn10/halcyon/actions/workflows/release.yml)
![License](https://img.shields.io/badge/License-MIT-yellow?logo=open-source-initiative&logoColor=white)

Halcyon is a Go-based CLI tool designed to identify and manage duplicate files in a directory, including nested subdirectories. With Halcyon, you can quickly detect duplicate files, view their paths, and take appropriate actions like moving, deleting, or listing them—all while ensuring a lightweight and fast execution.

---

## Features 🌟

- **Directory Traversal**: Recursively scans directories, including nested subdirectories, for a comprehensive duplicate search.
- **Duplicate Detection**:
  - Identifies duplicate files by name and size.
  - Detects similar file names (e.g., `file_copy`, `file (1)`).
- **Flexible Duplicate Management**:
  - **Move**: Move duplicates to a specified folder.
  - **Print**: Display paths of duplicate files.
  - **Delete**: Permanently delete duplicates.
- **User-Friendly CLI**: Simple and intuitive command-line interface for easy operation.

---

## Getting Started

### Installation ⚙️

You can install Halcyon using the following command:

```bash
curl -sSf https://raw.githubusercontent.com/tejastn10/halcyon/main/scripts/install.sh | bash
```

This will download and install the latest version of Halcyon.

### Uninstallation

To remove Halcyon from your system, run:

```bash
curl -sSf https://raw.githubusercontent.com/tejastn10/halcyon/main/scripts/uninstall.sh | bash
```

This will remove the installed binary from your system.

---

### Example Output

```bash
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

### Contributing 🤝

Contributions are welcome! Check out the [Contribution Guide](CONTRIBUTING.md) for details on how to get started. Feel free to open an issue or submit a pull request if you have ideas to enhance Halcyon.

### To-Do ✅

- Add support for advanced duplicate detection using file hashes.
- Allow configurable file actions via a config file.
- Improve performance for large directories.

---

## License 📜

This project is licensed under the MIT License. See the [LICENSE](LICENSE.md) file for details.

---

## Acknowledgments 🙌

- Named after **Halcyon**, symbolizing tranquility and clarity.
- Built with ❤️ and Go.
