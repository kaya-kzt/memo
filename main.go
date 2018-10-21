package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/jessevdk/go-flags"
)

type options struct {
	Dir    string `short:"d" long:"dir" env:"HOME" description:"directory you want to put memo"`
	Editor string `short:"e" long:"editor" env:"EDITOR" description:"editor name you want to use"`
	Format string `short:"f" long:"format" default:"20060102" description:"specify time format used in file name"`
	Type   string `short:"t" long:"filetype" default:"txt" description:"specify file content format. you can select 'txt' or 'md'"`
	Args   struct {
		Title string
	} `positional-args:"yes"`
}

func main() {
	// prepare options
	var opts options
	_, err := flags.Parse(&opts)
	if err != nil {
		return
	}

	// set default title or not
	if len(opts.Args.Title) == 0 {
		opts.Args.Title = "memo"
	}

	// if Dir is not exist, try to make directory
	dir := filepath.Clean(opts.Dir)
	baseDir := dir + "/memo"
	if _, err := os.Stat(baseDir); err != nil {
		if err := os.MkdirAll(baseDir, 0755); err != nil {
			fmt.Println(err)
			return
		}
	}

	// verify file type
	if opts.Type != "txt" && opts.Type != "md" {
		fmt.Println("invalid file type")
		return
	}

	// get current time
	t := time.Now()

	// try to open file
	fileName := baseDir + "/" + t.Format(opts.Format) + "_" + opts.Args.Title + "." + opts.Type
	if _, err := os.Stat(fileName); err != nil {
		file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
		}
		defer file.Close()
		fmt.Fprintln(file, getContent(opts.Args.Title, opts.Type, t))
		// fmt.Fprintln(file, "書き込み〜！") //書き込み
		fmt.Printf("finished to make file: %s\n", fileName)
	} else {
		fmt.Println("the same filename has already exist!")
	}
}
