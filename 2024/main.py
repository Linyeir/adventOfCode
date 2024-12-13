# AOC helper script


# Imports
import os
import datetime
import argparse
import aocd
import requests
import random
from dotenv import load_dotenv
from bs4 import BeautifulSoup

emojis = ["🎄","🔔","🍴","🎁","👶","🕯️","🎅","👼","🎶","⛪","🤶","❄️","☃️","👪","⛄","🌟","🔥","🍷","🦌","🍪","🥛","🧝‍♀️","🧦","🧝‍♂️","🧑‍🎄","🧝"]
year = 2024

def load_user_session():
    load_dotenv()
    user_session_id = os.getenv('USER_SESSION_ID')    
    return user_session_id

def get_puzzle_name(day):
    url = f"https://adventofcode.com/2024/day/{day}"
    try:
        response = requests.get(url)
        response.raise_for_status() 
        soup = BeautifulSoup(response.text, 'html.parser')

        h2_element = soup.find('h2')

        h2_text = h2_element.text.strip().strip('-').strip()
        _, puzzle_name = h2_text.split(": ", 1)
        puzzle_name  =  puzzle_name.strip()
        print(f"Found <h2>: {h2_text}")
        print(f"extracted: {puzzle_name}")
        return puzzle_name

    except requests.exceptions.RequestException as e:
        print(f"Error fetching the webpage for day {day}: {e}")
        return None

def init_day(day):
    print(f"Initializing day {day}...")

    # Creating all the names and variations
    puzzle_name = get_puzzle_name(day)
    words = puzzle_name.split()
    emoji_of_the_day = random.choice(emojis)
    puzzle_slug = puzzle_name.replace(" ", "-")
    folder_name = f"Day-{day:02d}-{puzzle_slug}"
    formatted_line = f"|   {emoji_of_the_day}    | [Day {day}: {puzzle_name}]({folder_name})  |"
    python_file_name = f"{words[0].lower() + ''.join(word.capitalize() for word in words[1:])}.py"
    folder_path = os.path.join(".", folder_name)
    python_file_path = os.path.join(folder_path, python_file_name)

    print(folder_path)

    # Add to readme
    try:
        with open("2024.md", "a") as readme_file:
            readme_file.write(f"{formatted_line}\n")
        print(f"Appended to readme.md: {formatted_line}")
    except Exception as e:
        print(f"Error writing to readme.md: {e}")

    # Create the folder
    try:
        os.makedirs(folder_path, exist_ok=True)
        print(f"Created folder: {folder_path}")
    except Exception as e:
        print(f"Error creating folder {folder_path}: {e}")
        return

    # Create the Python file with a basic template
    try:
        with open(python_file_path, "w") as python_file:
            python_file.write(f"# Advent of Code 2024 - Day {day}: {puzzle_name}\n")
            python_file.write("\n")
            python_file.write("def part1():\n")
            python_file.write("    # TODO: Implement the solution\n")
            python_file.write("    pass\n")
            python_file.write("\n")
            python_file.write("if __name__ == '__main__':\n")
            python_file.write("    main()\n")
        print(f"Created Python file: {python_file_path}")
    except Exception as e:
        print(f"Error creating Python file {python_file_path}: {e}")

def test_day(day):
    print(f"Testing solution for day {day}...")

def run_day(day):
    print(f"Running solution for day {day}...")

if __name__ == "__main__":
    # Create the argument parser
    parser = argparse.ArgumentParser(description="Advent of Code Automation Script")
    load_user_session()

    # Add arguments
    parser.add_argument(
        '--init', 
        type=int, 
        help="Initialize a specific day's setup."
    )
    parser.add_argument(
        '--test', 
        type=int, 
        help="Test the solution for a specific day."
    )
    parser.add_argument(
        '--run', 
        type=int, 
        help="Run the solution for a specific day."
    )
    
    # Parse the arguments
    args = parser.parse_args()

    print(args)
    
    # Handle the arguments
    if args.init is not None:
        init_day(args.init)
    elif args.test is not None:
        test_day(args.test)
    elif args.run is not None:
        run_day(args.run)
    else:
        print("No valid arguments provided. Use --init, --test, or --run.")
