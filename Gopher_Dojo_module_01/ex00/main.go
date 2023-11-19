package main

import (
	"io"
	"os"
)

func main() {
    var ret int

    if len(os.Args[1:]) > 0 {
        for _, arg := range os.Args[1:] {
            switch arg {
                case "-":
                    ret = cat(os.Stdin, os.Stdout)
                default:
                    if file := open(arg); file != nil {
                        ret = cat(file, os.Stdout)
						file.Close()
                    } else {
                        ret = 1;
                    }
                }
            }
    } else {
        ret = cat(os.Stdin, os.Stdout)
    }
    os.Exit(ret)
}

func cat(reader io.Reader, writer io.Writer) int {
    var buff [1]byte
    var _, err = reader.Read(buff[:])
    for err == nil {
        _, err = writer.Write(buff[:]); if err != nil { return 1 }
        _, err = reader.Read(buff[:])
    }
    if err != io.EOF { return 1 }
    return 0
}

func open(path string) *os.File {
	var file, err = os.Open(path)
	if err != nil {
		os.Stderr.Write([]byte("ft_cat: " + err.Error()[5:] + "\n"))
		return nil
	}
    info, err := os.Stat(path)
    if err != nil {
        os.Stderr.Write([]byte("ft_cat: " + err.Error()[5:] + "\n"))
		file.Close()
        return nil
    }
    if info.IsDir() {
        os.Stderr.Write([]byte("ft_cat: " + path + ": Is a directory\n"))
		file.Close()
        return nil
    }

    return file
}
