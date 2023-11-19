package imgconv_test

//go test -coverprofile=coverage.out -coverpkg=./imgconv ; go tool cover -html=coverage.out

import (
	"Road-to-Mercari/Gopher-Dojo-module-01/ex01/convert/imgconv/imgconv"
	"testing"
)

type parseArgsTestCase struct {
    args 		[]string
    in  		imgconv.Format
    out 		imgconv.Format
    should_fail	bool
}

func TestParseArgs(t *testing.T) {
    t.Parallel()

    var tests = []parseArgsTestCase{
        {[]string{"-i=png", "-o=jpg", ""},  imgconv.PNG, imgconv.JPG, false},
        {[]string{"-o=jpg", "-i=png", ""},  imgconv.PNG, imgconv.JPG, false},
        {[]string{"-i=jpg", "-o=png", ""},  imgconv.JPG, imgconv.PNG, false},
        {[]string{"-i=jpeg", "-o=png", ""}, imgconv.JPG, imgconv.PNG, false},
        {[]string{"-o=png", "-i=jpg", ""},  imgconv.JPG, imgconv.PNG, false},
        {[]string{"-i=png", "-o=jpg", ""},  imgconv.PNG, imgconv.JPG, false},
        {[]string{"-o=jpg", "-i=png", ""},  imgconv.PNG, imgconv.JPG, false},
        {[]string{"-i=jpg", "-o=png", ""},  imgconv.JPG, imgconv.PNG, false},
        {[]string{"-o=png", "-i=jpg", ""},  imgconv.JPG, imgconv.PNG, false},
        {[]string{""}, 					    imgconv.JPG, imgconv.PNG, false},
        {[]string{"", ""}, 				    imgconv.JPG, imgconv.PNG, true},
        {[]string{"", "", ""},			    imgconv.JPG, imgconv.PNG, true},
        {[]string{"-o=p_g", "-i=jpg", ""},  imgconv.JPG, imgconv.PNG, true},
        {[]string{"-o=png", "-i=j_g", ""},  imgconv.JPG, imgconv.PNG, true},
        {[]string{"-o=png", "-o=jpg", ""},  imgconv.JPG, imgconv.PNG, true},
        {[]string{"-i=png", "-i=jpg", ""},  imgconv.JPG, imgconv.PNG, true},
        {[]string{"-i=png", "-o=png", ""},  imgconv.JPG, imgconv.PNG, true},
    }

    for _, test := range(tests) {
        var in, out, err = imgconv.ParseArgs(test.args)
        if err != nil && test.should_fail == false { t.Error(test.args, err.Error()); continue }
        if err == nil && test.should_fail == true { t.Error(test.args, "No error"); continue }
		if test.should_fail == true { continue }
        if in != test.in || out != test.out { t.Error(test.args, "Parsing error"); continue }
    }
}

func TestAnalizeFiles(t *testing.T) {
	t.Parallel()

    testAnalizeFilesJPGtoPNG(t)
    testAnalizeFilesPNGtoJPG(t)
}

func TestConvert(t *testing.T) {
	t.Parallel()

    testConvertJPGtoPNG(t)
    testConvertPNGtoJPG(t)
}

type analizeOrConverFilesTestCase struct {
    dir 		string
    should_fail	bool
}

func testAnalizeFilesJPGtoPNG(t *testing.T) {
    var conv = imgconv.NewConverter(imgconv.JPG, imgconv.PNG)
    var tests = []analizeOrConverFilesTestCase{
        {"1", false},
        {"2", true},
        {"no_such_dir", true},
        {"3", false},
    }

    for _, test := range(tests) {
        var err = conv.AnalizeFiles("../testdata/jpg/" + test.dir)
        if err != nil && test.should_fail == false { t.Error(test.dir, err.Error()); continue }
        if err == nil && test.should_fail == true { t.Error(test.dir, "No error"); continue }
    }
}

func testAnalizeFilesPNGtoJPG(t *testing.T) {
    var conv = imgconv.NewConverter(imgconv.PNG, imgconv.JPG)
    var tests = []analizeOrConverFilesTestCase{
        {"1", false},
        {"2", true},
        {"3", false},
    }

    for _, test := range(tests) {
        var err = conv.AnalizeFiles("../testdata/png/" + test.dir)
        if err != nil && test.should_fail == false { t.Error(test.dir, err.Error()); continue }
        if err == nil && test.should_fail == true { t.Error(test.dir, "No error"); continue }
    }
}

func testConvertJPGtoPNG(t *testing.T) {
    var conv = imgconv.NewConverter(imgconv.JPG, imgconv.PNG)
    var tests = []analizeOrConverFilesTestCase{
        {"4", false},
        {"5", false},
        {"no_such_dir", true},
        {"6", false},
        {"7", false},
    }

    for _, test := range(tests) {
        var err = conv.Convert("../testdata/jpg/" + test.dir)
        if err != nil && test.should_fail == false { t.Error(test.dir, err.Error()); continue }
        if err == nil && test.should_fail == true { t.Error(test.dir, "No error"); continue }
    }
}

func testConvertPNGtoJPG(t *testing.T) {
    var conv = imgconv.NewConverter(imgconv.PNG, imgconv.JPG)
    var tests = []analizeOrConverFilesTestCase{
        {"4", false},
        {"5", false},
        {"6", false},
    }

    for _, test := range(tests) {
        var err = conv.Convert("../testdata/png/" + test.dir)
        if err != nil && test.should_fail == false { t.Error(test.dir, err.Error()); continue }
        if err == nil && test.should_fail == true { t.Error(test.dir, "No error"); continue }
    }
}
