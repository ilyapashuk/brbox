// тире в диалогах
// вариации тире уже должны быть заменены дефисами на момент вызова этого обработчика

function handle(line,opts) {
if (line.startsWith("- ")) {
return "-" + line.slice(2)
} else {
return line
}
}