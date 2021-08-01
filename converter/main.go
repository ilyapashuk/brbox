// Copyright 2021 ilyapashuk<ilusha.paschuk@gmail.com>

//    This file is part of Brbox.

//    Brbox is free software: you can redistribute it and/or modify
//    it under the terms of the GNU General Public License as published by
//    the Free Software Foundation, either version 3 of the License, or
//    (at your option) any later version.

//    Brbox is distributed in the hope that it will be useful,
//    but WITHOUT ANY WARRANTY; without even the implied warranty of
//    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//    GNU General Public License for more details.

//    You should have received a copy of the GNU General Public License
//    along with Brbox.  If not, see <https://www.gnu.org/licenses/>.

//this module provides an ability to convert electronic braille to different formats and process it with various filters

package converter
import "strconv"
import "brbox"
import "github.com/ilyapashuk/go-braille"
import "flag"
import "os"
import "path/filepath"
import "errors"
import "fmt"
import "strings"
var Readers map[string]func([]byte) (braille.BraillePage, error) = make(map[string]func([]byte) (braille.BraillePage, error))
var Writers map[string]func(braille.BraillePage) ([]byte, bool, error) = make(map[string]func(braille.BraillePage) ([]byte, bool, error))
var Filters map[string]func(params string, s braille.BraillePage) (braille.BraillePage, error) = make(map[string]func(params string, s braille.BraillePage) (braille.BraillePage, error))
func fatal(err error) {
if err == nil {
return
}
panic(err)
}

func init() {
brbox.Subcommands["conv"] = Converter
}
func Converter(args []string) {
cmdline := flag.NewFlagSet("conv", flag.ExitOnError)
cmdline.Usage = func() {
fmt.Println("usage: conv [options] <infile> [outfile]")
cmdline.PrintDefaults()
}
inform := cmdline.String("inform", "", "format of the input data")
outform := cmdline.String("outform", "", "format of the output data")
outext := cmdline.String("outext", "", "an extension, without dot, to automaticly generate output file name")
filt := cmdline.String("f", "", "process data with filters")
checkpaper := cmdline.Bool("checklinesize",false,"check weather lines of text are not overfilling linesize and refuse to continue if it is not true")
cmdline.Parse(args)
infile := cmdline.Arg(0)
indata,err := brbox.ReadInputFile(infile)
fatal(err)
var inf string
if *inform != "" {
inf = *inform
} else {
if infile == "-" {
fatal(errors.New("you must set input data format"))
}
ext := filepath.Ext(infile)
inf = string([]rune(ext)[1:])
}
if _,ok := Readers[inf]; ! ok {
fatal(errors.New("unknown input format"))
}
inpage,err := Readers[inf]([]byte(indata))
fatal(err)
if *filt != "" {
filtl := strings.Split(*filt, ",")
for _,f := range filtl {
ff := strings.Split(f, ":")
if len(ff) == 2 {
var err error
inpage,err = Filters[ff[0]](ff[1], inpage)
fatal(err)
} else {
var err error
inpage,err = Filters[ff[0]]("", inpage)
fatal(err)
}
}
}
if *checkpaper {
lss := os.Getenv("BRBOX_LINE_SIZE")
ls,_ := strconv.Atoi(lss)
for i,row := range inpage {
if len(row) > ls {
fatal(fmt.Errorf("line %d: linesize is %d, but must not be grater %d", i+1, len(row), ls))
}
}
}
var outfile string
if *outext != "" {
outfile = strings.TrimSuffix(infile, filepath.Ext(infile)) + "." + *outext
} else {
outfile = cmdline.Arg(1)
}
var outf string
if *outform != "" {
outf = *outform
} else {
if outfile == "-" {
fatal(errors.New("you must set output data format"))
}
ext := filepath.Ext(outfile)
outf = string([]rune(ext)[1:])
}
if _,ok := Writers[outf]; ! ok {
fatal(errors.New("unknown output format"))
}
outdata,bom,err := Writers[outf](inpage)
fatal(err)
fatal(brbox.WriteOutputFile(outfile, string(outdata), bom))
}

