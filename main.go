package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"os"
	"path"
	"path/filepath"
	"rmToTrash/utils"
	"strings"
	"time"
)

var trashPath string

type flagOption struct {
	Force     bool `short:"f" long:"force" description:"Force move to trash"`
	Recursive bool `short:"r" long:"recursive" description:"Recursive move directory all files to trash"`
}

func init() {
	home, err := utils.Home()
	if err != nil {
		panic(err)
	}

	trashPath = fmt.Sprintf("%s/.Trash/", home)
}

func main() {
	var opt flagOption
	parser := flags.NewParser(&opt, flags.Default)

	args, err := parser.ParseArgs(os.Args[1:])
	if err != nil {
		panic(err)
	}

	if len(args) == 0 {
		parser.WriteHelp(os.Stdout)
		return
	}

	for _, file := range args {
		absPath, _ := filepath.Abs(file)
		if utils.IsDir(absPath) && !opt.Recursive {
			fmt.Printf(" 无法删除\"%s\": 是一个目录", file)
			return
		}
	}

	for _, file := range args {
		absPath, _ := filepath.Abs(file)
		fileFullName := path.Base(absPath)
		ext := path.Ext(absPath)
		name := strings.TrimSuffix(fileFullName, ext)

		// The file not exists
		if !utils.FileExist(absPath) {
			if !opt.Force {
				fmt.Printf("无法删除\"%s\": 没有那个文件或目录\n", file)
			}
			continue
		}

		// The file exists in the recycle bin
		if utils.FileExist(trashPath + fileFullName) {
			now := time.Now()
			nano := now.Nanosecond()
			fileFullName = fmt.Sprintf("%s-%d%s", name, nano, ext)
		}

		// Whether to delete the file
		if !opt.Force {
			term := "是否删除目录"
			if utils.IsFile(absPath) {
				term = "是否删除文件"
			}

			var confirm string
			fmt.Printf("%s \"%s\"?", term, file)
			_, _ = fmt.Scanln(&confirm)

			if strings.ToLower(confirm) != "y" {
				continue
			}
		}

		err := os.Rename(absPath, trashPath+fileFullName)
		if err != nil {
			panic(err)
		}
	}
}
