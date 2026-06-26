package main

import (
	"io/ioutil"
	"path/filepath"
)

const (
	DryRun = true
)

//func mv(a, b string) error {
//
//}

func ls(path string) ([]string, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var filePaths []string
	for _, f := range files {
		filePaths = append(filePaths, filepath.Join(path, f.Name()))
	}

	return filePaths, nil
}

func main() {

	//targetDir, err := os.Getwd()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//filePaths, err := ls(targetDir)
	//
	//// Find the images and PDFs
	//var ignored []string
	//var process []string
	//for _, f := range filePaths {
	//	if isImage(f) || isPDF(f) {
	//		process = append(process, f)
	//	} else {
	//		ignored = append(ignored, f)
	//	}
	//}
	//
	//// Ensure they are formatted
	//for _, f := range process {
	//	baseName := filepath.Base(f)
	//	newBaseName := removeHHMMSS(baseName)
	//
	//	if baseName != newBaseName {
	//		log.Printf(">> %s\n", baseName)
	//		log.Printf("== %s\n", newBaseName)
	//	}
	//
	//	oldPath := filepath.Join(filepath.Dir(f), baseName)
	//	newPath := filepath.Join(filepath.Dir(f), newBaseName)
	//
	//	log.Printf("mv %q %q\n", oldPath, newPath)
	//}
	//
	//if err != nil {
	//	log.Fatalln(err)
	//}
}
