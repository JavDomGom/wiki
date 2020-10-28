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
