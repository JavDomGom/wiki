# Python

## Encontrar el número de veces que aparece un substring en un string

```Python
def frequencyCount(string, substr):
   count = 0
   pos = 0
   while(True):
       pos = string.find(substr , pos)
       if pos > -1:
           count = count + 1
           pos += 1
       else:
           break
   return count
print("The count is: ", frequencyCount("thatthathat","that"))
```

## Imprimir una cadena en filas de un ancho determinado

```Python
import textwrap

def wrap(string, max_width):
    return '\n'.join(textwrap.wrap(string, width=max_width))

if __name__ == '__main__':
    string, max_width = input(), int(input())
    result = wrap(string, max_width)
    print(result)
```

## Imprimir mapa de strings centrados

```Python
h, w = map(int, input().split())
t_top = 1
t_botom = h-2

for i in range(h):
    mid_h = h//2
    if i < mid_h:
        print(f'{".|."*t_top:-^{w}}')
        t_top += 2
    elif i == mid_h:
        print(f'{"WELCOME":-^{w}}')
    elif i > mid_h:
        print(f'{".|."*t_botom:-^{w}}')
        t_botom -= 2
```

Input:

```
11 33
```

Output:

```
---------------.|.---------------
------------.|..|..|.------------
---------.|..|..|..|..|.---------
------.|..|..|..|..|..|..|.------
---.|..|..|..|..|..|..|..|..|.---
-------------WELCOME-------------
---.|..|..|..|..|..|..|..|..|.---
------.|..|..|..|..|..|..|.------
---------.|..|..|..|..|.---------
------------.|..|..|.------------
---------------.|.---------------
```

## Imprimir un número entero en dec, oct, hex y bin

```Python
n_dec = 17

w = len(bin(n_dec)[2:])

for i in range (1, n_dec+1):
    n_bin = bin(i)[2:]
    n_oct = oct(i)[2:]
    n_hex = hex(i)[2:].upper()
    print(f'{i:>{w}} {n_oct:>{w}} {n_hex:>{w}} {n_bin:>{w}}')
```

```
    1     1     1     1
    2     2     2    10
    3     3     3    11
    4     4     4   100
    5     5     5   101
    6     6     6   110
    7     7     7   111
    8    10     8  1000
    9    11     9  1001
   10    12     A  1010
   11    13     B  1011
   12    14     C  1100
   13    15     D  1101
   14    16     E  1110
   15    17     F  1111
   16    20    10 10000
   17    21    11 10001
```

## Comprensión de listas

You are given three integers `x`, `y` and `z` representing the dimensions of a cuboid along with an integer `n`. Print a list of all possible coordinates given by `(i, j, k)` on a 3D grid where the sum of `i + j + k` is not equal to `n`. Here, `0<=i<=x;0<=j<=y;0<=k<=z`. Please use list comprehensions rather than multiple loops, as a learning exercise.

Example

```
x = 1
y = 1
z = 1
n = 3
```

All permutations of `[i, j, k]` are:

```
[[0, 0, 0], [0, 0, 1], [0, 1, 0], [0, 1, 1], [1, 0, 0], [1, 0, 1], [1, 1, 0], [1, 1, 1]]
```

Print an array of the elements that do not sum to `n = 3`.

```
[[0, 0, 0], [0, 0, 1], [0, 1, 0], [0, 1, 1], [1, 0, 0], [1, 0, 1], [1, 1, 0]]
```

```Python
x = 1
y = 1
z = 1
n = 3

print([[a,b,c] for a in range(0,x+1) for b in range(0,y+1) for c in range(0,z+1) if a + b + c != n ]
```

Output:

```
[[0, 0, 0], [0, 0, 1], [0, 1, 0], [0, 1, 1], [1, 0, 0], [1, 0, 1], [1, 1, 0]]
```

## Imprimir el hash de una tupla

Given an integer `n` and `n` space-separated integers as input, create a tuple `t` of those `n` integers. Then compute and print the result of `hash(t)`.

```Python
n = int(input())
integer_list = map(int, input().split())
integer_list = tuple([int(x) for x in integer_list])
print(hash(integer_list))
```

Input:

```
2
1 2
```

Output:

```
3713081631934410656
```

# Encuentra el subcampeón

Given the participants' score sheet for your University Sports Day, you are required to find the runner-up score. You are given `n` scores. Store them in a list and find the score of the runner-up.

Input Format:

The first line contains `n`. The second line contains an array `A[]` of integers each separated by a space.

Explanation:

Given list is `[2, 3, 6, 6, 5]`. The maximum score is `6`, second maximum is `5`. Hence, we print `5` as the runner-up score.

```Python
n = int(input())
arr = map(int, input().split())
numbers = list(set([x for x in arr]))
numbers.sort(reverse=True)
print(numbers[1])
```

Input:

```
5
2 3 6 6 5
```

Output:

```
5
```

# Listas anidadas

Given the names and grades for each student in a class of `N` students, store them in a nested list and print the name(s) of any student(s) having the second lowest grade.

Note: If there are multiple students with the second lowest grade, order their names alphabetically and print each name on a new line.

Example:

```Python
records = [['chi', 20.0], ['beta', 50.0], ['alpha', 50.0]]
```

The ordered list of scores is `[20.0, 50.0]`, so the second lowest score is `50.0`. There are two students with that score: `['beta', 'alpha']`. Ordered alphabetically, the names are printed as:

```
alpha
beta
```

```Python
records = []
scores = set()
for _ in range(int(input())):
    name = input()
    score = float(input())
    records.append([name, score])
    scores.add(score)

scores = list(scores)
scores.sort()
second_score = scores[1]

lista = [r for r in records if r[1] == scores[1]]
lista.sort()

for x in lista:
    print(x[0])
```
