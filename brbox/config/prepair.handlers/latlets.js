// проставление признаков латинских малых и латинских больших букв
// вспомогательные функции
const latlets = "abcdefghijklmnopqrstuvwxyz"
function buildchobj(lets) {
let obj = Object.create(null)
for (let chr of lets) {
obj[chr] = true
}
return obj
}
let latobj = buildchobj(latlets)
Object.keys(latobj).forEach(function(key) {
latobj[key.toUpperCase()] = true
})
function isLatlet(l) {
return (l in latobj)
}
function isupper(l) {
return l === l.toUpperCase()
}


let pm = charForDots("6")
let pb = charForDots("46")

function handle(line,opts) {
let latm = false
// строка разбивается на символы
let chars = line.split("")
return chars.map(function(char) {
if (char === ' ') {
return char
}
if (! isLatlet(char)) {
latm = false
return char
} else {
if (! latm) {
latm = true
return (isupper(char) ? pb : pm) + char
} else {
return char
}
}
}).join("")
}