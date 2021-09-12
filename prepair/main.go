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

// main file for this module

package prepair
import "strings"
import "flag"
import "fmt"
import "brbox"
import "path/filepath"

import "brbox/configuration"
import "github.com/mattn/go-shellwords"
var BomSequence = []byte("\xef\xbb\xbf")
// такой формат сохранения множиств символов введен для компактности кода
type SymbolGroupe string
func (c SymbolGroupe) Contains(r rune) bool {
for _,i := range c {
if r == i {
return true
}
}
return false
}
var Digits SymbolGroupe = SymbolGroupe("0123456789")

type Handler interface {
Handle(string, []string) *string
}
type HandlerFunc func(string, []string) *string
func (c HandlerFunc) Handle(st string, opts []string) *string {
return c(st, opts)
}
var Handlers map[string]Handler = make(map[string]Handler)

func CallHandler(t string, hn string, args []string) *string {
res := Handlers[hn].Handle(t,args)
return res
}
func ScriptCallHandler(t,hn string, args []string) interface{} {
res := CallHandler(t,hn,args)
if res == nil {
return false
} else {
return *res
}
}
func ListHandlers() []string {
res := make([]string,0,len(Handlers))
for k,_ := range Handlers {
res = append(res,k)
}
return res
}

// this function will check weather provided rune is a digit
func isdig(s rune) bool {
return Digits.Contains(s)
}

type HandlerChain [][]string
func (c HandlerChain) Handle(t string,_ []string) *string {
res := t
for _,v := range c {
var r *string
if len(v) == 1 {
r = Handlers[v[0]].Handle(res,nil)
} else {
r = Handlers[v[0]].Handle(res,v[1:])
}
if r == nil {
return nil
}
res = *r
}
return &res
}

func Prepair(args []string) {
cmdline := flag.NewFlagSet("prepair", flag.ExitOnError)
cmdline.Usage = func() {
fmt.Println("usage: prepair [options] <infile>")
cmdline.PrintDefaults()
fmt.Println("available handlers:")
for key,_ := range Handlers {
fmt.Println(key)
}
}
scriptname := cmdline.String("script",filepath.Join(configuration.ConfDir, "prepair.script"), "preparation script file")
outext := cmdline.String("outext","","extension for new file")
cmdline.Parse(args)
fn := cmdline.Arg(0)
t,err := brbox.ReadInputFile(fn)
if err != nil {
panic(err)
}
script,err := brbox.ReadInputFile(*scriptname)
if err != nil {
panic(err)
}
scriptl := strings.Split(script, "\n")
shellwords.ParseEnv = true
var hc [][]string
for _,line := range scriptl {
if line == "" {
continue
}
if strings.HasPrefix(line, "#") {
continue
}
args,err := shellwords.Parse(line)
if err != nil {
panic(err)
}
hc = append(hc,args)
}
t = PrepairText(hc,t)
var rfn string
if *outext != "" {
rfn = strings.TrimSuffix(fn, filepath.Ext(fn)) + "." + *outext
} else {
rfn = cmdline.Arg(1)
}
err = brbox.WriteOutputFile(rfn, t, true)
if err != nil {
panic(err)
}
}

func PrepairText(hc [][]string, t string) string {
hcc := HandlerChain(hc)
lines := strings.Split(t,"\n")
res := make([]*string, len(lines))
for i,l := range lines {
res[i] = hcc.Handle(l, nil)
}
sb := new(strings.Builder)
sb.Grow(len(t))
for _,l := range res {
if l != nil {
sb.WriteString(*l)
sb.WriteString("\n")
}
}
return sb.String()
}
func LineWrapCommand(args []string) {
arg := args[0]
ifn := args[1]
ofn := args[2]
indata,err := brbox.ReadInputFile(ifn)
if err != nil {
panic(err)
}
res := PrepairText([][]string{{"linewrap",arg}}, indata)
if err != nil {
panic(err)
}
err = brbox.WriteOutputFile(ofn,res,true)
}
func init() {
brbox.Subcommands["prepair"] = Prepair
brbox.Subcommands["linewrap"] = LineWrapCommand
}