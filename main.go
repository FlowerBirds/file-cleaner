package main

import (
	"flag"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"text/template"
)

const DELETE_CONF = "delete.conf"

var (
	cloudHome = flag.String("c", "", "")
	deleteConf = flag.String("d", DELETE_CONF, "")
	help = flag.Bool("h", false, "")
)

var Version = "0.0.0"

type templateData struct {
	Version string
}

var usage = `Usage: file-cleaner [options...]
Version: {{ .Version }}
删除delete.conf文件中列出来的文件，需要设置CLOUD_HOME环境变量或者通过-c进行指定根目录

Options:
	-c set CLOUD_HOME directory for find delete.conf and delete files.
	-d set the delete conf file name at CLOUD_HOME dir, default is delete.conf.
	-h help info.
`
func main() {
	flag.Usage = func() {
		tmpl := template.Must(template.New("example").Parse(usage))
		data := templateData{Version: Version}
		tmpl.Execute(os.Stderr, data)
		// fmt.Fprint(os.Stderr, fmt.Sprintf(usage, runtime.NumCPU()))
	}
	flag.Parse()
	if *cloudHome == "" {
		*cloudHome = os.Getenv("CLOUD_HOME")
	}

	if flag.NArg() < 1 && *cloudHome == "" {
		usageAndExit("")
	}
	if *help {
		usageAndExit("")
	}
	if *cloudHome == "" {
		usageAndExit("CLOUD_HOME is not set")
	}

	fi, err := os.Stat(*cloudHome)
	if os.IsNotExist(err) {
		log.Fatalln("directory is not exist: " + *cloudHome)
		os.Exit(2)
	}
	if err != nil {
		log.Fatalln(err)
		os.Exit(2)
	}
	if !fi.IsDir() {
		log.Fatalln("directory is not a directory: " + *cloudHome)
		os.Exit(2)
	}
	if runtime.GOARCH == "linux" && !strings.HasSuffix(*cloudHome, "/") {
		*cloudHome = *cloudHome + "/"
	}
	if runtime.GOOS == "windows" && !strings.HasSuffix(*cloudHome, "\\") {
		*cloudHome = *cloudHome + "\\"
	}
	log.Println("CLOUD_HOME is " + *cloudHome)
	deleteListFile := *cloudHome + *deleteConf
	_, err = os.Stat(deleteListFile)
	if os.IsNotExist(err) {
		log.Fatalln("deleted list file is not exist: " + deleteListFile)
		os.Exit(2)
	}
	existList, erorList, err := readFile(deleteListFile)
	if err != nil {
		os.Exit(2)
	}
	for _, f := range erorList {
		log.Fatalf("delete file is not exist: %s\n", f)
	}
	for _, line := range existList {
		deleteDirAndFile(line)
	}

}

func usageAndExit(msg string) {
	if msg != "" {
		fmt.Fprintf(os.Stderr, msg)
		fmt.Fprintf(os.Stderr, "\n\n")
	}
	flag.Usage()
	fmt.Fprintf(os.Stderr, "\n")
	os.Exit(1)
}

func PathExist(path string) (bool, bool) {
	fi, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false, false
	}
	return true, fi.IsDir()
}

func readFile(file string) ([]string, []string, error)  {
	var existList []string
	var errorList []string
	b, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("read file: %v error: %v", file, err)
		return existList, errorList, err
	}
	s := string(b)
	for _, lineStr := range strings.Split(s, "\n") {
		lineStr = strings.TrimSpace(lineStr)
		if lineStr == "" {
			continue
		}
		tmp := *cloudHome + lineStr
		if e, _  := PathExist(tmp); e {
			existList = append(existList, tmp)
		} else {
			errorList = append(errorList, tmp)
		}
	}
	return existList, errorList, nil
}

func deleteDirAndFile(path string) {
	fi, err := os.Stat(path)
	if os.IsNotExist(err) {
		log.Fatalf("read file: %v error: %v", path, err)
	}
	if fi.IsDir() {
		filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
			log.Println("delete file(walk): " + path)
			return nil
		})
	}
	os.RemoveAll(path)
	log.Println("delete file(or directory): " + path)
}
