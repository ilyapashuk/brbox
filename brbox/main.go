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

// go build command should be called hear

package main
import "os"
import "brbox"
import _ "brbox/prepair"
import _ "brbox/converter"
import _ "brbox/configuration"
import _ "brbox/tableonly"
import "fmt"
func main() {
if len(os.Args) < 3 {
fmt.Println("available subcommands:")
for key,_ := range brbox.Subcommands {
fmt.Println(key)
}
return
}
sc := os.Args[1]
if r,ok := brbox.Subcommands[sc]; ok {
r(os.Args[2:])
} else {
os.Stderr.Write([]byte("subcommand " + sc + " not found\n"))
os.Exit(1)
}
}