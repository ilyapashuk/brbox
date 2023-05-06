const digits = "0123456789"

function buildchobj(lets) {
let obj = Object.create(null)
for (let chr of lets) {
obj[chr] = true
}
return obj
}
let digobj = buildchobj(digits)
function isdig(chr) {
return (chr in digobj)
}

const ns = charForDots("3456")

function handle(text, opts) {
let dig = false
let chars = text.split("")

return chars.map(function(char) {
let cdig = isdig(char)
if (cdig) {
if (! dig) {
dig = true
return ns + char
}
} else {
if (char === ',') {
return char
}
dig = false
}
return char
}).join("")
}