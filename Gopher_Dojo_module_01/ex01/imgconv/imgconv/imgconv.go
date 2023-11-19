// Provide all the functions and user defined types used
// by the main package for converting images
package imgconv

import (
	"bytes"
	"errors"
	"image"
	"image/jpeg"
	"image/png"
	"net/http"
	"os"
	"strings"
)

// define an image format
type Format int8

// enum for all the available image format
const (
    invalid Format = iota
    JPG     Format = iota
    PNG     Format = iota
)

// data for the conversion
type Converter struct {
    input_format  Format
    output_format Format
}

// create a new converter with a given input and output type
func NewConverter(in Format, out Format) *Converter {
    var conv = new(Converter)
    conv.input_format = in
    conv.output_format = out
    return conv
}

// start the recursive conversion on a directory ignore the files that are not of the input type and ignore all the errors
func (conv *Converter) Convert(dir string) error {
    var files, err = os.ReadDir(dir); if err != nil { return remove_prefix(err) }
    for _, file := range files {
        if file.IsDir() {
            if err = conv.Convert(dir + "/" + file.Name()); err != nil { return remove_prefix(err) }
        } else {
            var bytes, err = os.ReadFile(dir + "/" + file.Name()); if err != nil { return remove_prefix(err) }
            if formatToMIME(conv.input_format) == http.DetectContentType(bytes) {
                img, err := decodeImg(bytes, conv.input_format); if err != nil { return remove_prefix(err) }
                byte_buff, err := encodeImg(img, conv.output_format); if err != nil { return remove_prefix(err) }
                if err = os.WriteFile(convertFileName(dir + "/" + file.Name(), conv.input_format, conv.output_format), byte_buff, 0644); err != nil { return remove_prefix(err) }
            }
        }
    }
    return nil
}

// check if all the files are of the of the input type, should be run before Convert if the programe need to stop on bad type or error
func (conv *Converter) AnalizeFiles(dir string) error {
    var files, err = os.ReadDir(dir); if err != nil { return remove_prefix(err) }
    for _, file := range files {
        if file.IsDir() {
            if err = conv.AnalizeFiles(dir + "/" + file.Name()); err != nil { return remove_prefix(err) }
        } else {
            var bytes, err = os.ReadFile(dir + "/" + file.Name()); if err != nil { return remove_prefix(err) }
            if formatToMIME(conv.input_format) != http.DetectContentType(bytes) { 
                return errors.New(file.Name() + " is not a valid file")
            }
        }
    }
    return nil
}

// parse the commande line argument to retive the input and output formats
func ParseArgs(argv []string) (Format, Format, error) {
    if len(argv) != 1 && len(argv) != 3 {
        return invalid, invalid, errors.New("invalid argument")
    }
    if len(argv) == 1 {
        return JPG, PNG, nil
    }
    var in, out = invalid, invalid
    for _, v := range argv[:2] {
        if strings.HasPrefix(v, "-i=") {
            if in != invalid {
                return invalid, invalid, errors.New(v + ": input format redefinition")
            }
            in = strToFormat(v[3:])
            if in == invalid { return invalid, invalid, errors.New(v + ": invalid input format") }
        } else if strings.HasPrefix(v, "-o=") {
            if out != invalid {
                return invalid, invalid, errors.New(v + ": output format redefinition")
            }
            out = strToFormat(v[3:])
            if out == invalid { return invalid, invalid, errors.New(v + ": invalid output format") }
        } else { 
            return invalid, invalid, errors.New(v + ": invalid argument") 
        }
    }
    if in == out {
        return invalid, invalid, errors.New("same input and output format")
    }
    return in, out, nil
}

func strToFormat(str string) Format {
    switch strings.ToLower(str) {
        case "jpg":  return JPG
        case "jpeg": return JPG
        case "png":  return PNG
    }
    return invalid
}

func formatToMIME(f Format) string {
    switch f {
        case JPG: return "image/jpeg"
        case PNG: return "image/png"
    }
    return ""
}

func decodeImg(byte_buff []byte, format Format) (image.Image, error) {
    switch format {
        case JPG: return jpeg.Decode(bytes.NewReader(byte_buff))
        case PNG: return png.Decode(bytes.NewReader(byte_buff))
    }
    return *new(image.Image), nil
}

func encodeImg(img image.Image, format Format) ([]byte, error) {
    var buf = new(bytes.Buffer)
    switch format {
    case JPG: 
        if err := jpeg.Encode(buf, img, nil); err != nil { return *new([]byte), err }
    case PNG:  
        if err :=  png.Encode(buf, img);      err != nil { return *new([]byte), err }
    }
    return buf.Bytes(), nil
}

func convertFileName(name string, src Format, dst Format) string {
    var new_name string

    if src == PNG {
        if strings.HasSuffix(name, ".png") {
            new_name = name[0:len(name) - 4]
        } else {
            new_name = name
        }
    }
    if src == JPG {
        if strings.HasSuffix(name, ".jpg") {
            new_name = name[0:len(name) - 4]
        } else if strings.HasSuffix(name, ".jpeg") {
            new_name = name[0:len(name) - 5]
        } else {
            new_name = name
        }
    }
    switch dst {
        case PNG: return new_name + ".png"
        case JPG: return new_name + ".jpg"
    }
    return ""
}

func remove_prefix(err error) error {
    if strings.HasPrefix(err.Error(), "open ") {
        return errors.New(err.Error()[5:])
    } else if strings.HasPrefix(err.Error(), "fdopendir ") {
        return errors.New(err.Error()[10:])
    } 
    return err
}