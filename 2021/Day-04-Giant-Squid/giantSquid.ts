import { readFileSync } from 'fs';

console.log('--- AOC 21 --- Day 4: Giant Squid ---');

const input = readFileSync('./input_test.txt', 'utf-8').split('\n\n')

type Coordinates = {
    x: number;
    y: number
}
type BingoNumber = {
    value: number;
    marked: boolean
}

firstpuzzle(input)
// secondpuzzle(input)

function firstpuzzle(input: string[]) {
    let calledNumbers = input.shift()?.split(',').map(Number);
    let cards = 
    input
    console.log(calledNumbers)
    console.log(input)

}