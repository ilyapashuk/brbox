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

package brbox
import (
"os"
"io"
"io/ioutil"
"bytes"
"strings"
"golang.org/x/text/encoding/charmap"
_ "embed"
"runtime"
"github.com/ilyapashuk/go-braille/translation"
)

var BomSequence = []byte("\xef\xbb\xbf")

var Subcommands map[string]func(args []string) = make(map[string]func(args []string))
// some commonly used functions
// эта функция осуществляет считывание поданного на вход текстового файла
// если файл имеет bom, он считывается как utf8 с удалением bom, иначе он считывается как файл в кодировке cp1251
// если вместо имени файла введён символ дефиса, то считывание осуществляется с stdin, файл в этом случае должен представлять собой utf8 без bom с окончаниями строк в стиле unix
func ReadInputFile(fn string) (string,error) {
var t string
if fn == "-" {
res,err := io.ReadAll(os.Stdin)
if err != nil {
return "",err
}
// text from stdin is always bomless unicode with unix style line endings
t = string(res)
return t,nil
}
res,err := ioutil.ReadFile(fn)
if err != nil {
return "",err
}
// if this file begins with bom sequence, this is utf8 and can be read directly with stripping bom
if bytes.HasPrefix(res, BomSequence) {
t = string(bytes.Trim(res, "\xef\xbb\xbf"))
} else {
// otherwise treat this as an ansii text file (cp1251)
t,err = charmap.Windows1251.NewDecoder().String(string(res))
if err != nil {
return "",err
}
}
t = strings.ReplaceAll(t, "\r", "")
return t,nil
}
// запись выходного текстового файла. всегда осуществляется в utf8, проставка bom настраиваема. также поддерживает отправку данных в stdout при помощи дефиса
func WriteOutputFile(fn string, data string, bom bool) error {
if fn == "-" {
// bomless unicode with unix line endings will go to stdout
_,err := os.Stdout.Write([]byte(data))
if err != nil {
return err
}
} else {
// if on windows add win style endings
var ddata string = data
if runtime.GOOS == "windows" {
ddata = strings.ReplaceAll(ddata,"\n","\r\n")
}
if bom {
ddata = string(append(BomSequence, []byte(ddata)...))
}
err := os.WriteFile(fn, []byte(ddata), 0644)
if err != nil {
return err
}
}
return nil
}
var DosTable translation.RuleList
// загружает таблицу bxt
func LoadDosTable() translation.RuleList {
if DosTable != nil {
return DosTable
}
table,err := os.ReadFile(os.Getenv("BRBOX_BXT_TABLE"))
if err != nil {
panic(err)
}
ts := string(table)
ts = strings.ReplaceAll(ts, "\r", "")
td := strings.Split(ts, "\n")
rl,err := translation.ParseRuleList(td)
if err != nil {
panic(err)
}
DosTable = rl
return rl
}