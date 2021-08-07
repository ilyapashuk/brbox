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

// this handler will strip any blank lines or lines containing only spaces

package prepair

import "strings"

func noBlankLinesHandler(line string, _ []string) *string {
ll := line
if line == "" {
return nil
}
if strings.TrimSpace(line) == "" {
return nil
}
return &ll
}
func trimSpaceHandler(s string, _ []string) *string {
res := strings.TrimSpace(s)
return &res
}

func init() {
Handlers["noblank"] = HandlerFunc(noBlankLinesHandler)
Handlers["trimspace"] = HandlerFunc(trimSpaceHandler)
}