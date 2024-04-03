package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

const defaultDestination = "bin/execfile"

func main() {
	// 1. 현재 프로그램 runtime의 directory
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Failed while resolving current directory.")
		return 
	}

	// 2. go-build-n-run을 실행할 수 있는 standalone go 환경인지 check (main.go file이 있는가)
	result := checkFileExists(filepath.Join(currentDir, "main.go"))
	if !result {
		fmt.Println("main.go file is missing in your directory...")
		return 
	}

	// 3. go build의 output dir
	destination := getDestination()

	// 4. script command 실행 (go build -o {outputDir} && ./{outputDir})
	execBuildAndRun(destination)
}

func getDestination() string {
	var destination string
	if len(os.Args) > 1 {
		destination = os.Args[1]
	} else {
		destination = defaultDestination
	}
	return destination
}

func checkFileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}

func execBuildAndRun(dest string) {
	cmd := exec.Command("/bin/sh", "-c", fmt.Sprintf("go build -o %s && ./%s", dest, dest))
	output, err := cmd.CombinedOutput()
	printOutput(output, err)
}

func printOutput(out []byte, e error) {
	if e != nil {
		fmt.Println("Comamnd ERROR: \n", e)
		fmt.Println()
	}
	fmt.Println("Command output: \n", string(out))
}