package main

import (
	"difffile/internal/diff"
	"difffile/internal/export"
	"difffile/internal/initialize"
	"flag"
	"fmt"
	"strings"
)

var srcPath *string = flag.String("p", ".", "src path")
var fileSuffix *string = flag.String("f", "", "suffix file")
var exportPath *string = flag.String("o", "./difffile", "dst export path")
var mode *string = flag.String("t", "init", "run type [init|diff]")
var version *bool = flag.Bool("v", false, "the version of difffile")

const banner string = `
difffile v1.0.1
`

func main() {
	fmt.Print(banner)
	flag.Parse()
	if *version {
		return
	}

	if *srcPath == "" {
		fmt.Println("must use a src path")
		return
	}

	var suffixs []string
	if *fileSuffix != "" {
		suffixs = strings.Split(*fileSuffix, ",")
	}

	switch strings.ToLower(*mode) {
	case "init":
		i := initialize.NewInit(*srcPath, *srcPath, suffixs...)
		if err := i.Exec(); err != nil {
			fmt.Println("difffile init exec fail", err)
			return
		}
		fmt.Println("difffile init success!!!")
	case "diff":
		d := diff.NewDiff(*srcPath, suffixs...)
		diffFiles, err := d.Diff()
		if err != nil {
			fmt.Println("difffile diff exec fail", err)
			return
		}
		fmt.Println("difffile diff success!!!")
		fmt.Println("change file size:", len(diffFiles))
		if len(diffFiles) > 0 {
			fmt.Println("change files:")
			for i := range diffFiles {
				fmt.Printf("\t%s\n", diffFiles[i])
			}

			if err = export.ExportFile(diffFiles, *srcPath, *exportPath); err != nil {
				fmt.Println("difffile export exec fail", err)
				return
			}
		}

		/*
			i := initialize.NewInit(*srcPath, *srcPath, suffixs...)
			if err := i.Exec(); err != nil {
				fmt.Println("difffile init exec fail", err)
				return
			}
			fmt.Println("difffile init success!!!")
		*/
	default:
		fmt.Println("invalid run type", *mode)
	}
}
