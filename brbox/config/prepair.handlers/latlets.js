// проставление признаков латинских малых и латинских больших букв
// вспомогательные функции
const latlets = "abcdefghijklmnopqrstuvwxyz"
function isLatlet(l) {
return latlets.includes(l.toLowerCase())
}
function isupper(l) {
return l === l.toUpperCase()
}

// главная часть кода
let pm = charForDots("6")
let pb = charForDots("46")


// тут мы обрабатываем построчно

let lines = text.split("\n")

lines.map(function(line) {
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
}).join("\n")