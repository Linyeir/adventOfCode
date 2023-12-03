import re

f = open("input", "r")
text = f.read()
lines = text.splitlines()
scv = 0

for line in lines:        
    digits = re.findall("\d", line)
    calibration_value = int(digits[0]) * 10 + int(digits[-1])
    scv += calibration_value

print("solution for part 1:", scv)

dict_digits = {    
    'oneight': 'oneeight',
    'twone': 'twoone',
    'threeight': 'threeeight',
    'fiveight': 'fiveeight',
    'nineight': 'nineeight',
    'eightwo': 'eighttwo',
    'eighthree': 'eightthree',
    'zero': '0',
    'one': '1',
    'two': '2',
    'three': '3',
    'four': '4',
    'five': '5',
    'six': '6',
    'seven': '7',
    'eight': '8',
    'nine': '9',
}

for key, value in dict_digits.items():
    text = text.replace(key, value)

lines = text.splitlines()
scv = 0

for line in lines:        
    digits = re.findall("\d", line)
    calibration_value = int(digits[0]) * 10 + int(digits[-1])
    scv += calibration_value

print("solution for part 2:", scv)
