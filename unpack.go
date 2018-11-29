package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/bmatcuk/doublestar"
)

var (
	buildCommit   = ""
	supportedExts = []string{".zip", ".tar.gz", ".tgz", ".gz"}
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: unpack FILE")
		os.Exit(0)
	}
	srcFile := os.Args[1]
	extract(srcFile)
}

func extract(srcFile string) {
	destDir, files, err := uncompress(srcFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(files) == 0 {
		return
	}

	wd, _ := os.Getwd()
	os.Chdir(destDir)
	for _, f := range files {
		extract(f)
	}
	os.Chdir(wd)
}

func getFileWithoutExt(name string, ext string) string {
	return name[:len(name)-len(ext)]
}

func getSupportedFileExt(name string) string {
	for _, ext := range supportedExts {
		if strings.HasSuffix(name, ext) {
			return ext
		}
	}
	return ""
}

func uncompress(srcFile string) (string, []string, error) {
	ext := getSupportedFileExt(srcFile)
	destDir := getFileWithoutExt(srcFile, ext)
	files := []string{}

	cmd := ""
	cmdArgs := []string{}
	isDir := false

	switch ext {
	case ".zip":
		cmd = "unzip"
		cmdArgs = []string{srcFile, "-d", destDir}
		isDir = true
	case ".tgz", ".tar.gz":
		cmd = "tar"
		cmdArgs = []string{"-zxvf", srcFile, "-C", destDir}
		isDir = true
	case ".gz":
		cmd = "gunzip"
		cmdArgs = []string{srcFile}
		isDir = false
	default:
		return destDir, files, nil
	}

	fmt.Printf("Extracting %s (%s)\n", srcFile, cmd)

	if isDir {
		if _, err := os.Stat(destDir); !os.IsNotExist(err) {
			fmt.Printf("Directory \"%s\" already exists!\n", destDir)
			os.Exit(1)
		}
		os.Mkdir(destDir, 0755)
	}

	output, err := exec.Command(cmd, cmdArgs...).CombinedOutput()
	if err != nil {
		os.Stderr.WriteString(string(output))
		os.Stderr.WriteString(err.Error())
		return destDir, files, err
	}

	if isDir {
		wd, _ := os.Getwd()
		os.Chdir(destDir)
		files, _ = doublestar.Glob("**")
		os.Chdir(wd)
	}

	return destDir, files, nil
}
