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

// support of table driven bxt encoding, sea docs for details

package converter

import "brbox"
import . "github.com/ilyapashuk/go-braille"
func rbrfwriter(p BraillePage) ([]byte, bool, error) {
rl := brbox.LoadDosTable()
t := rl.ToBackTable()
res,err := t.TranslateText(p)
if err != nil {
return nil, false, err
}
return []byte(res), true, nil
}
func rbrfreader(d []byte) (BraillePage, error) {
rl := brbox.LoadDosTable()
t := rl.ToForwardTable()
res,err := t.TranslateText(string(d))
if err != nil {
return nil,err
}
return res,nil
}
func init() {
Writers["bxt"] = rbrfwriter
Readers["bxt"] = rbrfreader
}