package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/larshelmer/md2present/internal/report"
)

//nolint
var (
	version  string
	revision string
	created  string
)

func main() {
	fmt.Println("md2report", version, revision, created)
	filename := flag.String("file", "", "file to generate report from")
	flag.Parse()
	if *filename == "" {
		out := flag.CommandLine.Output()
		fmt.Fprintf(out, "missing flag: -file\n")
		fmt.Fprintf(out, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(2)
	}

	err := report.Generate(*filename)
	if err != nil {
		log.Fatal(err)
	}
}
