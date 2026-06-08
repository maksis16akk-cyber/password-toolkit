#!/usr/bin/env python3
import random, string, sys, argparse

def check_strength(pwd):
    score = 0
    if len(pwd) >= 8: score += 1
    if any(c.isdigit() for c in pwd): score += 1
    if any(c.isupper() for c in pwd) and any(c.islower() for c in pwd): score += 1
    if any(c in string.punctuation for c in pwd): score += 1
    return ["Weak", "Medium", "Strong", "Very Strong"][score-1] if score>0 else "Weak"

def generate(length=12):
    chars = string.ascii_letters + string.digits + string.punctuation
    return ''.join(random.choice(chars) for _ in range(length))

if __name__ == '__main__':
    parser = argparse.ArgumentParser()
    parser.add_argument('-g', '--generate', type=int, nargs='?', const=12, help='generate password')
    parser.add_argument('-c', '--check', help='check password strength')
    args = parser.parse_args()
    if args.generate:
        pwd = generate(args.generate)
        print(f"Generated: {pwd}\nStrength: {check_strength(pwd)}")
    elif args.check:
        print(f"Strength: {check_strength(args.check)}")
    else:
        print("Use -g [length] or -c <password>")
