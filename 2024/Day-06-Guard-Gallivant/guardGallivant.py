# Advent of Code 2024 - Day 6: Guard Gallivant

movements = {
    "<": (0, -1),
    "v": (1, 0),
    ">": (0, 1),
    "^": (-1, 0)
}

right_turn = {
    "<": "^",
    "v": "<",
    ">": "v",
    "^": ">"
}

GridPoint = tuple[int, int]
Grid = dict[GridPoint, str] 

def add_points(point1: GridPoint, point2: GridPoint) -> GridPoint:
    return (point1[0] + point2[0], point1[1] + point2[1])

def parse_to_grid(input: list[str])-> Grid:
    grid = {}
    for y, row in enumerate(input):
        for x, char in enumerate(row):
            grid[(x, y)] = char
    return grid

def print_grid(grid: Grid, width: int, height: int):
    for y in range(height):
        line = ""
        for x in range(width):
            line += grid[(x, y)]
        print(line)

def get_start_position(grid: Grid) -> GridPoint:
    return next(point for point, value in grid.items() if value == "^")

def part1(grid):
    guard = get_start_position(grid)



    while True:

        print_grid(grid, 10, 10)
        print(guard)

        next_position = add_points(guard, movements[grid[guard]])

        if next_position not in grid:
            break   

        if grid[next_position] == "#":
            next_position = add_points(guard, movements[right_turn[guard]])
            grid[next_position] = right_turn[guard]
        else:
            grid[next_position] = grid[guard]
        
        grid[guard] = "X"
        guard = next_position

    # Count the entries with the value "#"
    count = sum(1 for value in grid.values() if value == "#")

    return count

if __name__ == '__main__':
    with open("example", "r") as file:
        
        data = file.read().split("\n")
        grid = parse_to_grid(data)

    print_grid(grid, len(data), len(data[0]))    
    print(part1(grid))