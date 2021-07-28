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
func lstripper(p string, s BraillePage) (BraillePage, error) {
ls,err := strconv.Atoi(p)
if err != nil {
return nil,err
}
res := make(BraillePage, 0, len(s))
for _,v := range s {
if len(v) > ls {
res = append(res, v[:ls])
res = append(res, v[ls:])
} else {
res = append(res, v)
}
}
return res,nil
}
func init() {
Filters["lstrip"] = lstripper
}