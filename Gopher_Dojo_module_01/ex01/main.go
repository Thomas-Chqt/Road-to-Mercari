package main

import (
	"Road-to-Mercari/Gopher-Dojo-module-01/ex01/convert/imgconv/imgconv"
	"fmt"
	"os"
)

func main() {
	
	var in, out, err = imgconv.ParseArgs(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, "error: " + err.Error())
		os.Exit(1)
	}

    var conv = imgconv.NewConverter(in, out)
	
    err = conv.AnalizeFiles(os.Args[len(os.Args)-1])
    if err != nil {
        fmt.Fprintln(os.Stderr, "error: " + err.Error())
        os.Exit(1)
    }

    err = conv.Convert(os.Args[len(os.Args)-1])
    if err != nil {
        fmt.Fprintln(os.Stderr, "error: " + err.Error())
        os.Exit(1)
    }
}
