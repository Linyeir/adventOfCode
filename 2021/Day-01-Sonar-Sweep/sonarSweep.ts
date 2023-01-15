import { readFileSync } from 'fs';

console.log('--- AOC 21 --- Day 1: Sonar Sweep ---');

const input = readFileSync('./input.txt', 'utf-8').split("\n").map(Number)

firstPuzzle(input)
secondPuzzle(input)

function firstPuzzle(input: number[]) {
    let counter = 0;

    for (let i = 1; i < input.length; i++) {
        if (input[i] > input[i - 1]) {
            counter++
        }
    }
    console.log("--- Part One ---")
    console.log(counter)
}

function secondPuzzle(input: number[]){
    var counter = 0;

    for (var i = 0; i < input.length - 3; i++) {
        var firstWindow = input[i] + input[i+1] + input[i+2];
        var secondWindow = input[i+1] + input[i+2] + input[i+3];
        if (firstWindow < secondWindow) {
            counter++
        }
    }

    console.log("--- Part Two ---")
    console.log(counter)
}