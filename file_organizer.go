package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var (
	sourceDir string
	dryRun    bool
	logFile   = "sortify.log"
	categories map[string]string
)

// loadCategories loads categories from a JSON file
func loadCategories(configFile string) error {
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &categories)
	if err != nil {
		return err
	}
	return nil
}

// organizeFiles moves files into categorized directories
func organizeFiles(dir string) error {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			continue // Skip directories
		}

		ext := strings.ToLower(strings.TrimPrefix(filepath.Ext(file.Name()), "."))
		category, exists := categories[ext]
		if !exists {
			category = "Other"
		}

		targetDir := filepath.Join(dir, category)
		if !dryRun {
			err := os.MkdirAll(targetDir, os.ModePerm)
			if err != nil {
				return err
			}
		}

		srcPath := filepath.Join(dir, file.Name())
		destPath := filepath.Join(targetDir, file.Name())

		if dryRun {
			fmt.Printf("[DRY RUN] Moving: %s -> %s\n", srcPath, destPath)
		} else {
			err := os.Rename(srcPath, destPath)
			if err != nil {
				logAction(fmt.Sprintf("Error moving file: %s -> %s", srcPath, destPath))
			} else {
				logAction(fmt.Sprintf("Moved: %s -> %s", srcPath, destPath))
			}
		}
	}
	return nil
}

// logAction logs actions to a log file
func logAction(message string) {
	logEntry := fmt.Sprintf("%s - %s\n", time.Now().Format(time.RFC3339), message)
	_ = ioutil.WriteFile(logFile, []byte(logEntry), os.ModeAppend|os.ModePerm)
	fmt.Println(message)
}

func main() {
	flag.StringVar(&sourceDir, "source", "C:/Users/sgari/Downloads", "Source directory to organize files")
	flag.BoolVar(&dryRun, "dry-run", false, "Preview changes without moving files")
	flag.Parse()

	fmt.Println("Starting Sortify...")

	// Load categories
	err := loadCategories("categories.json")
	if err != nil {
		fmt.Println("Error loading categories:", err)
		return
	}

	// Organize files
	err = organizeFiles(sourceDir)
	if err != nil {
		fmt.Println("Error organizing files:", err)
	} else {
		fmt.Println("File organization complete!")
	}
}
