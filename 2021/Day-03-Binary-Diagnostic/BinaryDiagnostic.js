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
    var sums = new Array(lines[0].length).fill(0)

    lines.forEach(element => {
        for(var i=0; i < element.length; i++){
            if(element.charAt(i) == '1'){
                sums[i] ++
            }
        }
    });

    console.log(sums)

    var gammaBinary = "";
    var epsilonBinary = "";

    sums.forEach(element => {
        gammaBinary += ((element<lines.length/2) ? "0": "1")
        epsilonBinary += ((element<lines.length/2) ? "1": "0")
    });

    console.log(gammaBinary)

    var gamma = parseInt(gammaBinary, 2);
    var epsilon = parseInt(epsilonBinary, 2);
    
    console.log(gamma)
    console.log(epsilon)

    console.log(gamma*epsilon);
}

function secondpuzzle(lines){

}