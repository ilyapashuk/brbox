// тире в диалогах
// вариации тире уже должны быть заменены дефисами на момент вызова этого обработчика

text.split("\n").map(function(line) {
if (line.startsWith("- ")) {
return "-" + line.slice(2)
} else {
return line
}
}).join("\n")