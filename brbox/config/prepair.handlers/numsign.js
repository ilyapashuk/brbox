const digits = "0123456789"
function isdig(chr) {
return digits.includes(chr)
}

let ns = charForDots("3456")

let dig = false

let chars = text.split("")

chars.map(function(char) {
let cdig = isdig(char)
if (cdig) {
if (! dig) {
dig = true
return ns + char
}
} else {
dig = false
}
return char
}).join("")