# Advent of Code 2024 - Day 4: Ceres Search

def check_for_word(grid, row, col, direction):
    rows = len(grid)
    cols = len(grid[0])
    dr, dc = direction

    for i, char in enumerate("XMAS"):
        new_row = row + i * dr
        new_col = col + i * dc

        # First check bounds
        if not (0 <= new_row < rows and 0 <= new_col < cols):
            return False

        # Check character match
        if grid[new_row][new_col] != char:
            return False

    return True

def check_for_mas(grid, row, col):
    masses = 0
    rows = len(grid)
    cols = len(grid[0])

    # First check bounds
    if row == 0 or col == 0 or row == rows-1 or col == cols-1:
        return 0

    directions = [(-1, -1), (-1, 1), (1, -1), (1, 1)] 

    for dr, dc in directions:
        if grid[row+dr][col+dc] == "M" and grid[row-dr][col-dc] == "S":
            masses += 1
    
    return masses == 2
    


def part1(grid):
    directions = [
        [-1, 0],  # Up
        [1, 0],   # Down
        [0, -1],  # Left
        [0, 1],   # Right
        [-1, -1], # Up-left
        [-1, 1],  # Up-right
        [1, -1],  # Down-left
        [1, 1]    # Down-right
    ]
    number_of_words = 0

    for row_index, row in enumerate(grid):
        for col_index, char in enumerate(row):
            if char == "X":
                # print(f"'X' found at row {row_index}, column {col_index}")
                for _, direction in enumerate(directions):
                    number_of_words += int(check_for_word(grid, row_index, col_index, direction))

    return number_of_words


def part2(grid):
    number_of_words = 0
    for row_index, row in enumerate(grid):
        for col_index, char in enumerate(row):
            if char == "A":
                number_of_words += int(check_for_mas(grid, row_index, col_index))
    return number_of_words

if __name__ == '__main__':
    with open("input", "r") as file:
        data = [list(line.strip()) for line in file.readlines()]

    # print(part1(data))
    print(part2(data))