# Advent of Code 2024 - Day 5: Print Queue

import math
from graphlib import TopologicalSorter

def subset(set1, set2):
    for _, item in enumerate(set1):
        if not item in set2:
            return False 
    return True

def followsAllRules(rules, update):
    for _, rule in enumerate(rules):
        ## if both values of the rule are in the update
        if subset(rule, update):
            ## if the values of the rule are in the update in the correct order
            if not update.index(rule[0]) < update.index(rule[1]):
                return False

    return True

def part1(rules, sets):
    sum = 0
    # for each line in set
    for _, update in enumerate(sets):
        if followsAllRules(rules, update):
            sum += update[math.floor(len(update)/2)]
    return sum

def part2(rules, sets):
    sum = 0
    print(rules)
    for _, update in enumerate(sets):
        if not followsAllRules(rules, update):
            print(update)
            sorter = TopologicalSorter()
            for first, last in rules:
                if first in update and last in update:
                # only add if both rule elements are in this update
                    sorter.add(first, last)

            sorted_update = list(sorter.static_order())
            sum += sorted_update[math.floor(len(sorted_update)/2)]
    return sum  

if __name__ == '__main__':
    with open("input", "r") as file:
        data = file.read().split("\n\n")

    first_half = [list(map(int, line.split('|'))) for line in data[0].strip().split('\n')]
    second_half = [list(map(int, line.split(','))) for line in data[1].strip().split('\n')]

    # print(part1(first_half, second_half))
    print(part2(first_half, second_half))