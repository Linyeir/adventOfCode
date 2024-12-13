# Advent of Code 2024 - Day 2: Red-Nosed Reports

def max_adjacent_difference(arr):
    return max(abs(arr[i] - arr[i + 1]) for i in range(len(arr) - 1))

def steppy_smally(arr):
    return max_adjacent_difference(arr) <=3
    
def is_increasing(arr):
    return all(arr[i] < arr[i + 1] for i in range(len(arr) - 1))

def is_decreasing(arr):
    return all(arr[i] > arr[i + 1] for i in range(len(arr) - 1))

def is_strictly_monotone(arr):
    return is_increasing(arr) or is_decreasing(arr)

def dampener(arr):
    # First up check if the line is normaly okay, because then it is also okay if dampened
    if is_strictly_monotone(arr) and steppy_smally(arr):
        return True
    
    # Then check if if would be okay with any value removed
    for i in range(len(arr)):
        new_arr = arr[:i] + arr[i+1:]
        if is_strictly_monotone(new_arr) and steppy_smally(new_arr):
            return True
    
    return False
     

def part1(data):
    filtered = [line for line in data if steppy_smally(line)]
    return sum([is_strictly_monotone(line) for line in filtered])

def part2(data):
    return sum([dampener(line) for line in data])


if __name__ == '__main__':
    with open("input", "r") as file:
        example_data = [list(map(int, line.split())) for line in file]

    print(part1(example_data))
    print(part2(example_data))
