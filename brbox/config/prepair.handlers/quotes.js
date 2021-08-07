const clc = charForDots("356")

function handle(text, opts) {
let qm = false
let chars = text.split("")
return chars.map(function(lett) {
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
}