const fs = require('fs');

fs.readFile('./input.txt', 'utf8', (err, data) => {
    if (err) {
        console.error(err)
        return
    }
    const numbers = data.split('\n').map(Number);
    firstpuzzle(numbers)
    secondpuzzle(numbers)
})


function firstpuzzle(numbers) {
    var c = 0;

    for (var i = 0; i < numbers.length - 1; i++) {
        if (numbers[i] < numbers[i + 1]) {
            c++
        }
    }

    console.log(c)
}

function secondpuzzle(numbers){
    var c = 0;

    for (var i = 0; i < numbers.length - 3; i++) {
        var smallergroup = numbers[i] + numbers[i+1] + numbers[i+2];
        var largergroup = numbers[i+1] + numbers[i+2] + numbers[i+3];
        if (smallergroup < largergroup) {
            c++
        }
    }

    console.log(c)
}