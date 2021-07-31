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

// braille specific handlers

package prepair
import "strings"
import "brbox"
import "github.com/ilyapashuk/go-braille"
import "github.com/ilyapashuk/go-braille/translation"
func commaNoSpaceHandler(s string, _ []string) string {
var iscom bool
mapper := func(r rune) rune {
if r == ',' {
iscom = true
return r
}
if iscom == true && r == ' ' {
iscom = false
return -1
}
iscom = false
return r
}
return strings.Map(mapper,s)
}

func numSignHandler(s string, _ []string) string {
rl := brbox.LoadDosTable()
backt := rl.ToBackTable()
forwt := rl.ToForwardTable()
field,_ := braille.FieldFromString("3456")
ns := backt[field]
sfield,_ := braille.FieldFromString("6")
sdot := backt[sfield]
digs := make(map[braille.BrailleField]struct{}, 10)
for _,digit := range string(Digits) {
digs[forwt[digit]] = struct{}{}
}

sb := new(strings.Builder)
var dig bool
for _,r := range s {
cdig := isdig(r)
if cdig && dig == false {
dig = true
sb.WriteRune(ns)
sb.WriteRune(r)
continue
}
if (! cdig) && dig {
// this is not digit, but we should check weather it can be confused with a digit by comparing their braille translations
if trans,ok := forwt[r]; ok {
if _,okk := digs[trans]; okk {
// this is not a digit, but has translation identical to digit's, so we need to add a 6th dot
sb.WriteRune(sdot)
}
}
dig = false
}
sb.WriteRune(r)
}
return sb.String()
}

func quotesHandler(s string, _ []string) string {
var qm bool
mapper := func(r rune) rune {
if r == '"' {
if ! qm {
qm = true
return r
} else {
qm = false
return '%'
}
} else {
return r
}
}
return strings.Map(mapper,s)
}
func tableOnlyHandler(s string, _ []string) string {
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

func brailleUnicodeHandler(t string, _ []string) string {
rl := brbox.LoadDosTable()
bt := rl.ToBackTable()
mapper := func(r rune) rune {
if braille.IsBrailleUnicode(r) {
field,_ := braille.FieldFromUnicode(r)
return bt[field]
} else {
return r
}
}
return strings.Map(mapper,t)
}
func CharForDots(dots string) (rune,error) {
bf,err := braille.FieldFromString(dots)
if err != nil {
return 0,err
}
rl := brbox.LoadDosTable()
bt := rl.ToBackTable()
if r,ok := bt[bf]; ok {
return r,nil
} else {
return 0,translation.UnmappableDotsError{bf}
}
}
func tableReplaceHandler(t string, opts []string) string {
rpl := make([]string,len(opts))
for i := 0; i < len(opts); i += 2 {
rpl[i] = opts[i]
sub := opts[i+1]
var rsub string
for _,w := range strings.Split(sub," ") {
if strings.HasPrefix(w,"0") {
rsub += string([]rune(w)[1:])
} else {
rs,err := CharForDots(w)
if err != nil {
panic(err)
}
rsub += string(rs)
}
}
rpl[i+1] = rsub
}
rp := strings.NewReplacer(rpl...)
return rp.Replace(t)
}
func fileTableReplaceHandler(t string, opts []string) string {
fn := opts[0]
fcs,err := brbox.ReadInputFile(fn)
if err != nil {
panic(err)
}
lines := strings.Split(fcs,"\n")
return tableReplaceHandler(t,lines)
}
func init() {
Handlers["commaspace"] = HandlerFunc(commaNoSpaceHandler)
Handlers["numsign"] = HandlerFunc(numSignHandler)
Handlers["quotes"] = HandlerFunc(quotesHandler)
Handlers["tableonly"] = HandlerFunc(tableOnlyHandler)
Handlers["buni"] = HandlerFunc(brailleUnicodeHandler)
Handlers["treplace"] = HandlerFunc(tableReplaceHandler)
Handlers["treplacefile"] = HandlerFunc(fileTableReplaceHandler)
}