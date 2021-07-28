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
func init() {
Handlers["script"] = HandlerFunc(scriptHandler)
Handlers["scriptfile"] = HandlerFunc(scriptFileHandler)
}