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

}

function secondpuzzle(lines){

}