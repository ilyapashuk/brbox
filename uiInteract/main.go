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


package uiInteract

// subcommands in this package is intended to be called automaticly by user interfaces

import (
"os"
"strings"
"brbox"
)

func getTable() []byte {
table,err := os.ReadFile(os.Getenv("BRBOX_BXT_TABLE"))
if err != nil {
panic(err)
}
ts := string(table)
ts = strings.ReplaceAll(ts, "\r", "")
return []byte(ts)
}

func gettable(_ []string) {
t := getTable()
os.Stdout.Write(t)
}

func getconf(args []string) {
n := args[0]
res := os.Getenv(n)
os.Stdout.Write([]byte(res))
}

func getTableMappings(_ []string) {
rl := brbox.LoadDosTable()
backt := rl.ToBackTable()
forw := rl.ToForwardTable()
r := new(strings.Builder)

for i,v := range backt {
r.WriteRune(i.ToUnicode())
r.WriteRune(v)
}
r.WriteRune('\n')
for i,v := range forw {
r.WriteRune(i)
r.WriteRune(v.ToUnicode())
}
rr := r.String()
os.Stdout.Write([]byte(rr))
}

func init() {
brbox.Subcommands["gettable"] = gettable
brbox.Subcommands["getconf"] = getconf
brbox.Subcommands["getTableMappings"] = getTableMappings
}