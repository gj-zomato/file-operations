## 📁 File Operations CLI (Go)

A terminal-based file management tool written in Go 🧑‍💻. This CLI app allows you to **create**, **edit**, **view**, and **delete** `.txt` files interactively, with support for importing data from APIs and displaying file metadata (like size).

---

### ✨ Features

* 📁 Create text files or import `.json` files from any API
* ✏️ Edit `.txt` files interactively (append content)
* 📄 View contents of any file
* 🗑️ Delete files with confirmation
* 📂 View directory contents with file sizes
* 🌐 Prevent editing of API-based `.json` files
* 🚫 Prevent editing of non-`.txt` files
* ❗ Graceful error handling and input validation
* 🧠 Interactive, menu-based user experience with emojis for friendliness

---

### 📦 Folder Structure

```bash
file-operations/
├── main.go                    # CLI entry point with interactive menu
├── function/
│   └── file.go                # File utilities (create, write, update, view, etc.)
├── validation/
│   └── check.go               # Validation functions (existence, emptiness, etc.)
```

---

### 🚀 Getting Started

#### 1. Clone the Repository

```bash
git clone https://github.com/gj-zomato/file-operations.git
cd file-operations
```

#### 2. Build the Project

```bash
go build -o file-cli
```

#### 3. Run the CLI

```bash
./file-cli
```

> 📝 On Windows, use `file-cli.exe` instead.

---

### 📌 Usage Guide

Upon launching the program, you'll be prompted to enter a directory path. From there, you'll access a main menu:

```
🤖 What would you like to do?
  1️⃣  Create 🆕
  2️⃣  Edit ✏️
  3️⃣  View 📂
  4️⃣  Delete ❌
  5️⃣  Change Directory 📂
  6️⃣  Quit 🚪
```

Each option is self-guided and walks you through the necessary inputs.

---

### 🧠 File Rules

* ✅ Only `.txt` files are **editable**
* 🔒 `.json` files (imported from APIs) are **view-only**
* 📝 You can add content line by line, ending input by typing `END`

---

### ⚙️ Functions Breakdown

#### `function/file.go`

* `CreateFile`, `WriteFile`, `DisplayFile`, `UpdateFile`
* `ImportfromAPI`: saves API response to `.json`
* `ReadDir`: prints all files with human-readable size
* `PromptChoice`, `MultiLineInput`: for interactive CLI input

#### `validation/check.go`

* `DirPathExist`, `FileExist`, `IsEmpty`: for path and file validations

---

### 🧪 Example API Import Flow

1. Choose **Create > Import from API**
2. Enter API URL (e.g., `https://jsonplaceholder.typicode.com/todos/1`)
3. Enter directory and filename (no `.json` extension needed)
4. File will be created with the API response

---

### 📦 Dependencies

Only uses the Go standard library:

* `os`, `io`, `fmt`, `bufio`, `net/http`, `path/filepath`, `strings`, `errors`

No external packages required ✅

---

### 📋 To-Do / Improvements

* [ ] Add file renaming support
* [ ] Add copy/move functionality
* [ ] Add search inside files
* [ ] Optional logging to a log file
* [ ] Add unit tests

---

### 🙌 Contribution

PRs and suggestions welcome! Please open an [issue](https://github.com/gj-zomato/file-operations/issues) or submit a PR if you'd like to contribute.

---

### 📄 License

This project is open-source and available under the [MIT License](LICENSE).

---

Let me know if you want:

* Screenshots/gif previews
* A version with usage examples and sample inputs
* A short walkthrough video command set
