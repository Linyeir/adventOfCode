import { readFileSync } from 'fs';

console.log('--- AOC 21 --- Day 2: Dive! ---');

const input = readFileSync('./input.txt', 'utf-8').split('\n')

firstpuzzle(input)
secondpuzzle(input)


function firstpuzzle(input: string[]) {
    var horizontal = 0;
    var depth = 0;

    input.forEach(line => {
        let parts = line.split(' ')
        var number = Number(parts[1])
        switch (parts[0]) {
            case "forward":
                horizontal += number;
                break;
            case "down":
                depth += number
                break;
            case "up":
                depth -= number
                break

        }
    });
    console.log("--- Part One ---")
    console.log("Product of the final horizontal position and depth: ", horizontal * depth)
}

function secondpuzzle(input: string[]) {
    var horizontal = 0;
    var aim = 0;
    var depth = 0;

    input.forEach(line => {
        let parts = line.split(' ')
        var number = Number(parts[1])

        switch (parts[0]) {
        case "forward":
            horizontal += number
            depth += aim * number
            break;
        case "down":
            aim += number
            break;
        case "up":
            aim -= number
            break
        }
    });
    console.log("--- Part Two ---")
    console.log("Product of the final horizontal position and depth: ", horizontal * depth)
}