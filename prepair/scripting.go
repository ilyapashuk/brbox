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

package prepair

import "github.com/dop251/goja"
import "os"
import "brbox"
import "strings"
import "fmt"
import "path/filepath"
import "brbox/configuration"
type Stringer interface {
String() string
}
func ConstructStringsBuilder(c goja.ConstructorCall, r *goja.Runtime) *goja.Object {
sb := new(strings.Builder)
val := r.ToValue(sb)
return val.ToObject(r)
}
func PrepairRuntime() *goja.Runtime {
r := goja.New()
r.Set("getenv",os.Getenv)
r.Set("listHandlers",ListHandlers)
r.Set("callHandler",CallHandler)
r.Set("setenv",os.Setenv)
r.Set("unsetenv",os.Unsetenv)
r.Set("lookupEnv",os.LookupEnv)
r.Set("StringsBuilder",ConstructStringsBuilder)
r.Set("charForDots",func(d string) (string,error) {
r,err := CharForDots(d)
if err != nil {
return "",err
}
return string(r),nil
})
r.Set("print",fmt.Println)
return r
}

func scriptHandler(t string, opts []string) string {
script := opts[0]
r := PrepairRuntime()
r.Set("text",t)
val,err := r.RunString(script)
if err != nil {
panic(err)
}
res := val.Export()
switch res.(type) {
case string:
return res.(string)
case Stringer:
return res.(Stringer).String()
}
panic("invalid script return type")
}

func scriptFileHandler(t string, opts []string) string {
script,err := brbox.ReadInputFile(opts[0])
if err != nil {
panic(err)
}
return scriptHandler(t, []string{script})
}

type ScriptHandler struct {
FileName string
p *goja.Program
}
func (c *ScriptHandler) Handle(t string, opts []string) string {
if c.p == nil {
filedata,err := brbox.ReadInputFile(c.FileName)
if err != nil {
panic(err)
}
prog,err := goja.Compile(filepath.Base(c.FileName), filedata, true)
if err != nil {
panic(err)
}
c.p = prog
}
r := PrepairRuntime()
r.Set("text",t)
r.Set("opts",opts)
val,err := r.RunProgram(c.p)
if err != nil {
panic(err)
}
res := val.Export()
switch res.(type) {
case string:
return res.(string)
case Stringer:
return res.(Stringer).String()
}
panic("invalid script return type")
}
func init() {
Handlers["script"] = HandlerFunc(scriptHandler)
Handlers["scriptfile"] = HandlerFunc(scriptFileHandler)
hd := filepath.Join(configuration.ConfDir,"prepair.handlers")
fl,err := os.ReadDir(hd)
if err != nil {
if ! os.IsNotExist(err) {
panic(err)
}
}
for _,f := range fl {
if f.IsDir() {
continue
}
if filepath.Ext(f.Name()) == ".js" {
hn := strings.TrimSuffix(f.Name(), filepath.Ext(f.Name()))
Handlers[hn] = &ScriptHandler{FileName: filepath.Join(hd,f.Name())}
}
}
}