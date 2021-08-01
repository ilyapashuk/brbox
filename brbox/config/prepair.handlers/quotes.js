let clc = charForDots("356")
let qm = false
let chars = text.split("")
chars.map(function(lett) {
if (lett === '"' && (!qm)) {
qm = true
return lett
}
if (lett === '"' && qm) {
qm = false
return clc
}
return lett
}).join("")