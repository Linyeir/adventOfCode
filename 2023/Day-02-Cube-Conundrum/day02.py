f = open("input", "r")
text = f.read()

text = text.replace(";", " ;")
text = text.replace(",", "")
text = text.replace(":", "")

games = text.splitlines()

max_cubes = {
    "blue": 14,
    "green": 13,
    "red": 12
}
sum_of_games = 0
sum_of_powers = 0

for game in games:
    min_game_cubes = {
        "blue": 0,
        "green": 0,
        "red": 0
    }
    game = game.split(" ")
    is_possible = True
    for index, word in enumerate(game):
        if word in max_cubes:
            cubes = int(game[index-1])
            if cubes > max_cubes[word]:
                is_possible = False
            if cubes > min_game_cubes[word]:
                min_game_cubes[word] = cubes

    if is_possible:
        sum_of_games += int(game[1])
    sum_of_powers += min_game_cubes["blue"] * min_game_cubes["green"] * min_game_cubes["red"]

print("Solution of part 1:", sum_of_games)
print("Solution of part 2:", sum_of_powers)
