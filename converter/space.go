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

package converter
import . "github.com/ilyapashuk/go-braille"
import "strconv"
import "fmt"
func spacer(p string, s BraillePage) (BraillePage, error) {
ls,err := strconv.Atoi(p)
if err != nil {
return nil,err
}
res := make(BraillePage, len(s))
for i,r := range s {
if len(r) > ls {
return res,InTextError{i, fmt.Errorf("line has length %v, but maximum line length set to %v", len(r), ls)}
}
if len(r) == ls {
res[i] = r
}
if len(r) < ls {
delta := ls - len(r)
zs := make(BrailleRow, delta)
res[i] = append(r, zs...)
}
}
return res,nil
}
func init() {
Filters["space"] = spacer
}