const fs = require('fs');

fs.readFile('./input.txt', 'utf8', (err, data) => {
    if (err) {
        console.error(err)
        return
    }
    const lines = data.split('\n');
    firstpuzzle(lines)
    secondpuzzle(lines)
})

function firstpuzzle(lines) {
    var horizontal = 0;
    var depth = 0;

    lines.forEach(element => {
        var number = Number(element.charAt(element.length - 1))
        if (/forward/.test(element)) {
            horizontal += number
        } else if (/down/.test(element)) {
            depth += number
        } else if (/up/.test(element)) {
            depth -= number
        }
    });
    console.log(horizontal * depth)
}

function secondpuzzle(lines) {
    var horizontal = 0;
    var aim = 0;
    var depth = 0;

    lines.forEach(element => {
        var number = Number(element.charAt(element.length - 1))
        if (/forward/.test(element)) {
            horizontal += number
            depth += aim*number
        } else if (/down/.test(element)) {
            aim += number
        } else if (/up/.test(element)) {
            aim -= number
        }
    });
    console.log(horizontal * depth)
}