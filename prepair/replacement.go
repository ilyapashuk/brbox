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

// this handler can replace parts of text with help of the dictionary

package prepair
import "strings"
import "brbox"
func replaceHandler(s string, opts []string) *string {
repl := strings.NewReplacer(opts...)
res := repl.Replace(s)
return &res
}
func fileReplaceHandler(s string, opts []string) *string {
fn := opts[0]
fcs,err := brbox.ReadInputFile(fn)
if err != nil {
panic(err)
}
lines := strings.Split(fcs,"\n")
repl := strings.NewReplacer(lines...)
res := repl.Replace(s)
return &res
}
func init() {
Handlers["replace"] = HandlerFunc(replaceHandler)
Handlers["replacefile"] = HandlerFunc(fileReplaceHandler)
}