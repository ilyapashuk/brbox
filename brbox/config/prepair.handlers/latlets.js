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
let upperobj = {}
Object.keys(latobj).forEach(function(key) {
let uk = key.toUpperCase()
latobj[uk] = true
upperobj[uk] = true
})
function isLatlet(l) {
return (l in latobj)
}
function isupper(l) {
return (l in upperobj)
}

function getchardesc(chr) {
if (isLatlet(chr)) {
if (isupper(chr)) {
return 2
} else {
return 1
}
} else {
return 0
}
}

let pm = charForDots("6")
let pb = charForDots("46")

function handle(line,opts) {
let cd = 0
// строка разбивается на символы
let chars = line.split("")
return chars.map(function(char) {
if (char === ' ') {
return char
}
let cdd = getchardesc(char)
if (cdd > cd) {
cd = cdd
switch (cdd) {
case 2:
return pb + char
break
case 1:
return pm + char
break
}
} else {
cd = cdd
return char
}
}).join("")
}