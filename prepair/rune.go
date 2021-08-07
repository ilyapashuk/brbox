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

// this file provides some handlers for per rune operations

package prepair
import "strings"

func runeDropHandler(s string, opts []string) *string {
sg := SymbolGroupe(opts[0])
mapper := func(r rune) rune {
if sg.Contains(r) {
return -1
}
return r
}
result := strings.Map(mapper, s)
return &result
}

func runeOnceHandler(s string, opts []string) *string {
sg := SymbolGroupe(opts[0])
var orch rune
mapper := func(r rune) rune {
if r == orch {
if sg.Contains(r) {
return -1
}
}
orch = r
return r
}
res := strings.Map(mapper,s)
return &res
}

func init() {
Handlers["dropchr"] = HandlerFunc(runeDropHandler)
Handlers["once"] = HandlerFunc(runeOnceHandler)
}