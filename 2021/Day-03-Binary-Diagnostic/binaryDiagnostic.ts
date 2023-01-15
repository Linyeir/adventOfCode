import { readFileSync } from 'fs';

console.log('--- AOC 21 --- Day 3: Binary Diagnostic ---');

const input = readFileSync('./input.txt', 'utf-8').split('\n')

firstpuzzle(input)
secondpuzzle(input)

function firstpuzzle(input: string[]) {
    let sums = new Array(input[0].length).fill(0)

    input.forEach((line) => {
        line.split('').forEach((symbol, index) => {
            sums[index] += parseInt(symbol)
        })
    })

    let gammaRateString = "";
    sums.forEach(number => {
        gammaRateString += ((number < input.length / 2) ? "0" : "1")
    })

    let gamma = parseInt(gammaRateString, 2);

    console.log("--- Part One ---");
    //epsilon is always the binary inverse, so we do a little maths
    console.log(gamma * ((2) ** input[0].length - gamma - 1))

}

function secondpuzzle(input: string[]) {
    //make 2 copies of the data
    var generator = [...input]
    var scrubber = [...input]

    //until there is only one left
    for (let index = 0; generator.length > 1; index++) {
        //
        let referenceBitGenerators = countActiveBits(generator, index) >= generator.length / 2 ? "1" : "0"
        generator = generator.filter(line => line.charAt(index) === referenceBitGenerators)
    }

    for (let index = 0; scrubber.length > 1; index++) {
        let referenceBitScrubber = countActiveBits(scrubber, index) < scrubber.length / 2 ? "1" : "0"
        scrubber = scrubber.filter(line => line.charAt(index) === referenceBitScrubber)
    }

    console.log("--- Part Two ---");
    console.log(parseInt(generator[0], 2) * parseInt(scrubber[0], 2))
}

function countActiveBits(input: string[], position: number): number {
    let counter = 0;
    input.forEach(line => {
        counter += parseInt(line[position])
    })
    return counter
}