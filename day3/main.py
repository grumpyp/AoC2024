from typing import List
import re

def render_input(input = "input.txt") -> list:
    with open(input) as f:
        return f.read()

#part 1
def find_mul(input: str) -> List[tuple]:
    matched = re.findall('mul\(\d{1,},\d{1,}\)', input)
    sanitized = [i.replace('mul(', '').replace(')', '').split(',') for i in matched]
    res = sum([int(i[0]) * int(i[1]) for i in sanitized])

    return res

#part 2
def part2(input: str) -> List[tuple]:
    pat_do = re.compile('do\(\)')
    pat_dont = re.compile("don't\(\)")
    pos = 0
    out_do = []
    out_dont = []
    while m := pat_do.search(input, pos):
        pos = m.start() + 1
        out_do.append(pos)
    
    pos = 0
    while m := pat_dont.search(input, pos):
        pos = m.start() + 1
        out_dont.append(pos)

    print(out_do)
    print(out_dont)

    # cut the string with the do and don't positions
    # until first don't
    input_new = input[:out_dont[0]]

    input_first = find_mul(input_new)
    print(input_first)

    # from first don't to second do
    x = []
    for i in out_do:
        x.append((i, "do"))
    for i in out_dont:
        x.append((i, "dont"))
    
    # sort the list
    x = sorted(x, key=lambda x: x[0])
    print(x)

    endstring = []
    ac = False
    for i in x:
        if i[1] == "do" and not ac:
            ac = True
            endstring.append(i[0])
        elif i[1] == "dont" and ac:
            endstring.append(i[0])
            ac = False 


    print(endstring)
    res = 0
    for i in range(0, len(endstring), 2):
        try:
            res += find_mul(input[endstring[i]:endstring[i+1]])
        except:
            res += find_mul(input[endstring[i]:])


    
    print(res+input_first)
    return res + input_first
        


#find_mul(render_input())
part2(render_input())
