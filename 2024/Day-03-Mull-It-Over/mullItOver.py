# Advent of Code 2024 - Day 3: Mull It Over
import re

def part1(data):
    # Regular expression to match valid mul(X,Y) patterns
    pattern = r'mul\((\d{1,3}),(\d{1,3})\)'

    # Find all matches in the input string
    matches = re.findall(pattern, data)

    # Convert matches to tuples of integers and calculate the sum of products
    total_sum = sum(int(x) * int(y) for x, y in matches)

    return total_sum


def part2(data):
     # Regular expression to match valid mul(X,Y), do(), and don't() patterns
    pattern = r'mul\(\d{1,3},\d{1,3}\)|do\(\)|don\'t\(\)'
    
    # Find all matches in the input string in order of appearance
    matches = [match.group() for match in re.finditer(pattern, data)]
    
    dont = False
    activated = []

    for item in matches:
        if item == "don't()":
            dont = True
        elif item == "do()":
            dont = False
        elif not dont:
            activated.append(item)

    total_sum = 0

    for item in activated:
        match = re.match(r'mul\((\d+),(\d+)\)', item)
        if match:
            # Extract X and Y from the mul(X,Y) pattern
            x, y = map(int, match.groups())
            # Multiply the numbers and add to the total sum
            total_sum += x * y

    return total_sum

if __name__ == '__main__':
    with open("input", "r") as file:
        data = file.read()
        
    print(part1(data))
    print(part2(data))
