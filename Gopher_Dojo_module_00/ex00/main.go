package main

import (
	"Road-to-Mercari/ex00/convert/utils"
	"fmt"
	"os"
)

func main() {
    var in, out, err = utils.ParseArgs(os.Args[1:])
    if err != nil {
        fmt.Fprintln(os.Stderr, "error: " + err.Error())
        os.Exit(1)
    }

    var conv = utils.NewConverter(in, out)
	
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
