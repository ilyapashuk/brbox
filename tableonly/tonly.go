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

package tableonly

import "brbox"
import "github.com/ilyapashuk/go-braille"
import "strings"

func TableOnly(s string) string {
rl := brbox.LoadDosTable()
backt := rl.ToBackTable()
field,_ := braille.FieldFromString("123456")
sdots := backt[field]
brt := rl.ToForwardTable()
mapper := func(r rune) rune {
if r == '\n' || r == ' ' {
return r
}
if _,ok := brt[r]; ok {
return r
} else {
return sdots
}
}
return strings.Map(mapper, s)
}

func tonly(args []string) {
indata,err := brbox.ReadInputFile(args[0])
if err != nil {
panic(err)
}
err = brbox.WriteOutputFile(args[1], TableOnly(indata), true)
if err != nil {
panic(err)
}
}
func init() {
brbox.Subcommands["tableonly"] = tonly
}