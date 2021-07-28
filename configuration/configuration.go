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


package configuration

import "os"
import "path/filepath"
import "brbox"
import "strings"
import "github.com/mattn/go-shellwords"

var ConfDir string

func init() {
shellwords.ParseEnv = true
ourexe,_ := os.Executable()
ConfDir = filepath.Join(filepath.Dir(ourexe), "config")
if res,ok := os.LookupEnv("BRBOX_CONFIG_DIR"); ok {
ConfDir = res
} else {
os.Setenv("BRBOX_CONFIG_DIR", ConfDir)
}
conffile := filepath.Join(ConfDir, "config.txt")
confdata,err := brbox.ReadInputFile(conffile)
if err != nil {
if os.IsNotExist(err) {
return
}
panic(err)
}
conflines := strings.Split(confdata, "\n")
for _,confline := range conflines {
if confline == "" {
continue
}
if strings.HasPrefix(confline, "#") {
continue
}
args,err := shellwords.Parse(confline)
if err != nil {
panic(err)
}
os.Setenv(args[0], args[1])
}
}