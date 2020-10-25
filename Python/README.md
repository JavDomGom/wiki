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
