package main

import (
	"archive/zip"
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// reconstructFile function to reconstruct the original file from parts
func reconstructFile(targetDir string, reconstructedFilePath string) {
	reconstructedFile, err := os.Create(reconstructedFilePath)
	checkError(err)
	defer reconstructedFile.Close()

	writer := bufio.NewWriter(reconstructedFile)
	defer writer.Flush()

	folders, err := filepath.Glob(filepath.Join(targetDir, "folder_*"))
	checkError(err)

	totalParts := 0
	for _, folder := range folders {
		parts, err := filepath.Glob(filepath.Join(folder, "part_*.zip"))
		checkError(err)
		totalParts += len(parts)
	}

	currentPart := 0
	for _, folder := range folders {
		parts, err := filepath.Glob(filepath.Join(folder, "part_*.zip"))
		checkError(err)

		for _, partZip := range parts {
			partFile, err := unzipPart(partZip)
			checkError(err)

			err = appendPart(writer, partFile)
			checkError(err)

			err = os.Remove(partFile)
			checkError(err)

			currentPart++
			displayProgressBar(currentPart, totalParts)
		}
	}

	fmt.Printf("\nReconstruction complete. The original file is saved as %s\n", reconstructedFilePath)
}

// unzipPart function to unzip a part file and return the unzipped file path
func unzipPart(zipFilePath string) (string, error) {
	r, err := zip.OpenReader(zipFilePath)
	if err != nil {
		return "", err
	}
	defer r.Close()

	var unzippedFilePath string
	for _, f := range r.File {
		unzippedFilePath = filepath.Join(filepath.Dir(zipFilePath), f.Name)
		outFile, err := os.Create(unzippedFilePath)
		if err != nil {
			return "", err
		}
		defer outFile.Close()

		rc, err := f.Open()
		if err != nil {
			return "", err
		}
		defer rc.Close()

		_, err = io.Copy(outFile, rc)
		if err != nil {
			return "", err
		}
	}

	return unzippedFilePath, nil
}

// appendPart function to append the contents of a part file to the writer
func appendPart(writer *bufio.Writer, partFilePath string) error {
	partFile, err := os.Open(partFilePath)
	if err != nil {
		return err
	}
	defer partFile.Close()

	_, err = io.Copy(writer, partFile)
	return err
}
