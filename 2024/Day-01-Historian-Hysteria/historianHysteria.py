
import os
import pandas as pd
from dotenv import load_dotenv

from aocd import get_data
load_dotenv()
user_session_id = os.getenv('USER_SESSION_ID')

data = get_data(day=1)

def part1(data):
    left, right = unpack(data)
    return sum([abs(a - b) for a, b in zip(left, right)])

def part2(data):
    left, right = unpack(data)
    return sum(a * right.count(a) for a in left)

def unpack(data):
    # Convert the string into two lists
    list1 = []
    list2 = []
    for line in data.strip().split('\n'):
        a, b = map(int, line.split())
        list1.append(a)
        list2.append(b)

    # Step 2: Sort both lists
    list1.sort()
    list2.sort()

    return list1, list2

if __name__ == '__main__':
    print(part1(data))
    print(part2(data))
