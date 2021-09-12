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

package prepair
import "os"
import "fmt"
import "strings"
import "strconv"
// this wrapper prevents lines with length more then linesize to appear in the output file
type LineCutter struct {
b *strings.Builder
c int
LineSize int
}
func (c *LineCutter) WriteRune(r rune) {
if c.b == nil {
c.b = new(strings.Builder)
}
if r == '\n' {
c.c = 0
c.b.WriteRune('\n')
return
}
if c.c == c.LineSize {
c.WriteRune('\n')
}
c.c += 1
c.b.WriteRune(r)
}
func (c *LineCutter) WriteString(s string) {
for _,v := range s {
c.WriteRune(v)
}
}

// this function is part of our word wrapping implementation

func wordstr(l string, LineSize int, hyph bool) string {
if ! hyphl {
LoadHyphenation()
}
lw := strings.Split(l, " ")
cut := new(LineCutter)
cut.LineSize = LineSize
for _,word := range lw {
var splres []string
var spli int
var letm int = -1
rword := []rune(word)
for i,v := range rword {
cr := LangLetters.Contains(v)
if letm == -1 {
if cr {
letm = 0
continue
}
}
if letm == 0 {
if ! cr {
letm = 1
continue
}
}
if letm == 1 {
if cr {
letm = 0
splres = append(splres,string(rword[spli:i]))
spli = i
}
}
}
splres = append(splres,string(rword[spli:]))
fmt.Println(splres)
for i,ww := range splres {
w := ww
cycle:
if cut.c == LineSize {
cut.WriteRune('\n')
}
srem := LineSize - cut.c
wlen := len([]rune(w))
if wlen >= LineSize && cut.c == 0 {
cut.WriteString(w)
continue
}
ispace := true
if cut.c == 0 || i != 0 {
ispace = false
}
if ispace {
srem -= 1
}
if srem < wlen {
if hyph && srem >= 3 {
if IsHyphenatible(w) {
if hres,ok := TryHyphenate(w); ok {
if len(hres) > 1 {
var ss string
for _,v := range hres {
sss := ss + v
wlen := len([]rune(sss + "-"))
if wlen <= srem {
ss = sss
} else {
break
}
}
if ss != "" {
if ispace {
cut.WriteRune(' ')
}
cut.WriteString(ss + "-")
cut.WriteRune('\n')
w = strings.TrimPrefix(w,ss)
goto cycle
}
}
}
}
}
cut.WriteRune('\n')
goto cycle
} else {
if ispace {
cut.WriteRune(' ')
}
cut.WriteString(w)
}
}
}
return cut.b.String()
}

func lineWrapHandler(line string, opts []string) *string {
lss := os.Getenv("BRBOX_LINE_SIZE")
ls,err := strconv.Atoi(string(lss))
if err != nil {
panic(err)
}
var hyph bool = false
if opts[0] == "1" {
hyph = true
}
if len([]rune(line)) > ls {
res := wordstr(line, ls, hyph)
return &res
} else {
res := line
return &res
}
}

func init() {
Handlers["linewrap"] = HandlerFunc(lineWrapHandler)
}