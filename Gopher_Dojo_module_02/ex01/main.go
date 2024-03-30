package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
    if(len(os.Args[1:]) != 1) {
        fmt.Fprintln(os.Stderr, "Bad Args")
        os.Exit(1)
    }
    _, err := getSize(os.Args[1])
    if(err != nil) {
        fmt.Fprintln(os.Stderr, os.Args[0] + ": " + err.Error())
        os.Exit(1)
    }
}

func getSize(url string) (int, error) {
    req, err := http.NewRequest("HEAD", url, strings.NewReader("")); if err != nil { return 0, err }
    client := http.Client{}
    res, err := client.Do(req); if err != nil { return 0, err }
    defer res.Body.Close()

    if res.StatusCode != 200 && res.StatusCode != 206 {
        return 0, errors.New(strconv.Itoa(res.StatusCode))
    }
    if res.Header.Get("Accept-Ranges") != "bytes" {
        return 0, errors.New("server do not accepte ranges")
    }

    cont_len, err := strconv.Atoi(res.Header.Get("Content-Length")) ; if err != nil { return 0, err }

    return cont_len, nil
}