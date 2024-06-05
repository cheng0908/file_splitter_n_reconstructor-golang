package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:")
		fmt.Println("  To split: <command> split <path_to_file> <target_directory> <split_size_MB> [<max_files_threshold>]")
		fmt.Println("  To reconstruct: <command> reconstruct <target_directory> <reconstructed_file_path>")
		os.Exit(1)
	}

	command := os.Args[1]

	switch strings.ToLower(command) {
	case "split":
		if len(os.Args) < 5 || len(os.Args) > 6 {
			fmt.Println("Usage: <command> split <path_to_file> <target_directory> <split_size_MB> [<max_files_threshold>]")
			os.Exit(1)
		}
		originalFile := os.Args[2]
		targetDirectory := os.Args[3]
		splitSizeMB := os.Args[4]
		maxFiles := "1000"
		if len(os.Args) == 6 {
			maxFiles = os.Args[5]
		}
		splitFile(originalFile, targetDirectory, splitSizeMB, maxFiles)

	case "reconstruct":
		if len(os.Args) != 4 {
			fmt.Println("Usage: <command> reconstruct <target_directory> <reconstructed_file_path>")
			os.Exit(1)
		}
		targetDirectory := os.Args[2]
		reconstructedFilePath := os.Args[3]
		reconstructFile(targetDirectory, reconstructedFilePath)

	default:
		fmt.Println("Unknown command:", command)
		fmt.Println("Usage:")
		fmt.Println("  To split: <command> split <path_to_file> <target_directory> <split_size_MB> [<max_files_threshold>]")
		fmt.Println("  To reconstruct: <command> reconstruct <target_directory> <reconstructed_file_path>")
		os.Exit(1)
	}
}
