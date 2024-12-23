

---

# Sortify

**Sortify** is a Go-based CLI tool designed to organize files in a specified directory (e.g., your Downloads folder) into categorized subdirectories based on their extensions. It simplifies file management and keeps your directories neat and tidy.

---

## Features

- **Organize by File Type**: Moves files into folders like `Documents`, `Images`, `Videos`, etc., based on their extensions.
- **Dry Run Mode**: Preview the changes without actually moving files.
- **Customizable Categories**: Define file categories using a JSON configuration file.
- **Logging**: Logs actions performed for easy tracking and debugging.

---

## Getting Started

### Prerequisites

- Install [Go](https://golang.org/doc/install).
- A basic understanding of how to run commands in a terminal.

---

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/<your-username>/sortify.git
   cd sortify
   ```

2. Build the project:
   ```bash
   go build -o sortify.exe file_organizer.go
   ```

3. Create the `categories.json` file in the project root with the following content:
   ```json
   {
     "pdf": "Documents",
     "docx": "Documents",
     "zip": "Archives",
     "png": "Images",
     "jpg": "Images",
     "mp4": "Videos",
     "exe": "Applications",
     "msi": "Applications",
     "other": "Other"
   }
   ```

---

### Usage

1. Run the tool to organize files:
   ```cmd
   sortify.exe --source="C:/Users/sgari/Downloads"
   ```

2. Preview changes with `--dry-run`:
   ```cmd
   sortify.exe --source="C:/Users/sgari/Downloads" --dry-run
   ```

---

## File Categories

By default, the following categories are supported (configurable via `categories.json`):

| File Extension | Folder Name   |
|----------------|---------------|
| `.pdf`, `.docx` | Documents     |
| `.zip`         | Archives      |
| `.png`, `.jpg` | Images        |
| `.mp4`         | Videos        |
| `.exe`, `.msi` | Applications  |
| Others         | Other         |

---

## Logging

All actions are logged to `sortify.log` in the project directory. Each log entry includes a timestamp and the file action performed.

---

## Example

### Before Running Sortify:
```
Downloads/
├── example.pdf
├── photo.jpg
├── video.mp4
├── setup.exe
```

### After Running Sortify:
```
Downloads/
├── Documents/
│   └── example.pdf
├── Images/
│   └── photo.jpg
├── Videos/
│   └── video.mp4
├── Applications/
    └── setup.exe
```

---

## Future Enhancements

- Recursive directory scan.
- Interactive mode for manual file organization.
- Restore mode to revert moved files.
- Progress bar and enhanced logging.

---

## Contributing

Contributions are welcome! Please open an issue or submit a pull request if you'd like to help improve **Sortify**.

---

## License

No license. This is a personal project and is not licensed for external use.

---


