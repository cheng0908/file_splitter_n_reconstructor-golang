package main

import (
	"archive/zip"
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
)

func splitFile(filePath string, targetDir string, splitSizeMB string, maxFiles string) {
	file, err := os.Open(filePath)
	checkError(err)
	defer file.Close()

	fileInfo, err := file.Stat()
	checkError(err)
	totalSize := fileInfo.Size()

	splitSizeMBInt, err := strconv.Atoi(splitSizeMB)
	checkError(err)
	splitSizeBytes := splitSizeMBInt * 1024 * 1024

	maxFilesInt, err := strconv.Atoi(maxFiles)
	checkError(err)

	buffer := make([]byte, splitSizeBytes)
	reader := bufio.NewReader(file)
	partCount := 0
	folderCount := 0
	totalParts := (totalSize / int64(splitSizeBytes)) + 1

	for {
		n, err := reader.Read(buffer)
		if n == 0 && err == io.EOF {
			break
		}
		checkError(err)

		// Create new subdirectory if partCount is a multiple of maxFiles
		if partCount%maxFilesInt == 0 {
			folderCount++
		}

		subDir := filepath.Join(targetDir, fmt.Sprintf("folder_%d", folderCount))
		err = os.MkdirAll(subDir, os.ModePerm)
		checkError(err)

		partFilePath := filepath.Join(subDir, fmt.Sprintf("part_%05d", partCount))
		err = writePartFile(partFilePath, buffer[:n])
		checkError(err)

		zipFilePath := partFilePath + ".zip"
		err = zipPartFile(partFilePath, zipFilePath)
		checkError(err)

		err = os.Remove(partFilePath)
		checkError(err)

		partCount++

		// Show progress
		displayProgressBar(partCount, int(totalParts))
	}

	fmt.Printf("\nSplitting complete. Zipped files are saved in subdirectories within %s\n", targetDir)
}

func writePartFile(filePath string, data []byte) error {
	partFile, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer partFile.Close()

	_, err = partFile.Write(data)
	return err
}

func zipPartFile(srcFilePath string, destZipFilePath string) error {
	destZipFile, err := os.Create(destZipFilePath)
	if err != nil {
		return err
	}
	defer destZipFile.Close()

	zipWriter := zip.NewWriter(destZipFile)
	defer zipWriter.Close()

	srcFile, err := os.Open(srcFilePath)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	w, err := zipWriter.Create(filepath.Base(srcFilePath))
	if err != nil {
		return err
	}

	_, err = io.Copy(w, srcFile)
	return err
}
