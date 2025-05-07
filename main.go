package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
)

var IDs []string

func main() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error loading .env file.")
	}

	envIds := os.Getenv("IDS")

	IDs = strings.Split(envIds, ",")

	documentsFolderPath, err := os.UserHomeDir()

	if err != nil {
		fmt.Printf("Error getting user's documents directory: %v\n", err)
	}

	paintsFolder := filepath.Join(documentsFolderPath, "/Documents/iRacing/paint")

	filesDeleted, diskSpaceFreed := iterateOverFiles(paintsFolder)
	fmt.Printf("Files deleted: %v \nSpace freed: %v", filesDeleted, formatSpaceFreed(diskSpaceFreed))
}

func formatSpaceFreed(spaceFreed int64) string {
	spaceInKB := float32(spaceFreed / 1024)

	if spaceInKB > 1024 {
		spaceInMB := spaceInKB / 1024
		if spaceInMB > 1024 {
			spaceInGB := spaceInMB / 1024

			return fmt.Sprintf("%.2f GB", spaceInGB)
		}
		return fmt.Sprintf("%.2f MB", spaceInMB)
	}
	return fmt.Sprintf("%.2f KB", spaceInKB)
}

func iterateOverFiles(path string) (int, int64) {
	var filesDeleted int
	var diskSpaceFreed int64

	files, err := os.ReadDir(path)

	if err != nil {
		fmt.Printf("Error reading directory %v", err)
		return filesDeleted, diskSpaceFreed
	}

	for _, file := range files {
		fileName, fileSize := readFile(file)
		filePath := filepath.Join(path, fileName)

		if file.IsDir() {
			recFilesDel, recDiskSpace := iterateOverFiles(filePath)

			filesDeleted += recFilesDel
			recDiskSpace += recDiskSpace
		} else {
			if !stringArrayContains(IDs, fileName) {
				err := os.Remove(filePath)

				if err != nil {
					fmt.Printf("Error deleting file %v", err)
					continue
				}

				filesDeleted++
				diskSpaceFreed += fileSize
			}
		}
	}

	return filesDeleted, diskSpaceFreed
}

func stringArrayContains(arr []string, str string) bool {
	for _, element := range arr {
		if strings.Contains(str, element) {
			return true
		}
	}
	return false
}

func readFile(file os.DirEntry) (string, int64) {
	fileInfo, err := file.Info()

	if err != nil {
		fmt.Printf("Error reading file %v", err)
	}

	fileName := fileInfo.Name()
	fileSize := fileInfo.Size()

	return fileName, fileSize
}
