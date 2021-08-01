let digits = "0123456789"
function isdig(chr) {
return digits.includes(chr)
}

let ns = charForDots("3456")
let sb = new StringsBuilder()

let dig = false

let chars = text.split("")

for (let char of chars) {
let cdig = isdig(char)
if (cdig) {
if (! dig) {
sb.WriteString(ns)
dig = true
}
} else {
dig = false
}
sb.WriteString(char)
}
sb