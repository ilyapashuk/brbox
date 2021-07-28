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

// в этом файле находится код-обёртка вокруг библиотеки github.com/speedata/hyphenation, осуществляющей перенос текста по алгоритму TeX

package prepair

import "github.com/speedata/hyphenation"
import "strings"
import "os"
import "unicode"
import "bufio"

func splitword(w string, bp []int) []string {
var res []string
for i,v := range bp {
if i == 0 {
res = append(res, string([]rune(w)[:v]))
} else {
res = append(res, string([]rune(w)[bp[i-1]:v]))
}
}
res = append(res, string([]rune(w)[bp[len(bp)-1]:]))
var fres []string
for _,v := range res {
if v != "" {
fres = append(fres, v)
}
}
return fres
}
var excl map[string][]int = make(map[string][]int)

func AddExcl(w string) {
ww := strings.ReplaceAll(w, "-", "")
excl[ww] = ProcExcl(w)
}
func ProcExcl(w string) []int {
var res []int
for i,v := range []rune(w) {
if v == '-' {
res = append(res, i)
res = append(res, ProcExcl(strings.Replace(w, "-", "", 1))...)
return res
}
}
return res
}

var hyphl bool
var hlang *hyphenation.Lang

var LangLetters SymbolGroupe

func LoadHyphenation() {
pfn := os.Getenv("BRBOX_HYPHENATION_PATTERNS_FILE")
efn := os.Getenv("BRBOX_HYPHENATION_EXCEPTIONS_FILE")
llr := []rune(os.Getenv("BRBOX_LANG_LETTERS"))
lll := make([]rune, 0, len(llr) * 2)
for _,r := range llr {
lll = append(lll, r, unicode.ToUpper(r))
}
LangLetters = SymbolGroupe(string(lll))
pf,err := os.Open(pfn)
if err != nil {
panic(err)
}
l,err := hyphenation.New(pf)
if err != nil {
panic(err)
}
hlang = l
pf.Close()
ef,err := os.Open(efn)
if err != nil {
panic(err)
}
efscanner := bufio.NewScanner(ef)
for efscanner.Scan() {
AddExcl(efscanner.Text())
}
ef.Close()
hyphl = true
}

func IsHyphenatible(s string) bool {
if ! hyphl {
LoadHyphenation()
}
var letm bool = false
var letmn int
for _,r := range s {
if LangLetters.Contains(r) {
if !letm {
letm = true
letmn += 1
if letmn > 1 {
return false
}
}

} else {
letm = false
}
}
return true
}
func TryHyphenate(s string) (res []string, ok bool) {
defer func() {
if recover() != nil {
res = nil
ok = false
}
}()
ok = true
res = Hyphenate(s)
return
}
func Hyphenate(w string) []string {
if ! hyphl {
LoadHyphenation()
}
var res []int
if ex,ok := excl[strings.ToLower(w)]; ok {
res = ex
} else {
res = hlang.Hyphenate(w)
}
if len(res) == 0 {
return []string{w}
} else {
return splitword(w, res)
}

}