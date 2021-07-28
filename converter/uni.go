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

// braille unicode module

package converter
import . "github.com/ilyapashuk/go-braille"

func unireader(d []byte) (BraillePage, error) {
return PageFromUnicode(string(d))
}
func uniwriter(p BraillePage) ([]byte, bool, error) {
return []byte(p.ToUnicode()), true, nil
}
func init() {
Readers["buni"] = unireader
Writers["buni"] = uniwriter
}