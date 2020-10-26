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
