# Python

## Encontrar el número de veces que aparece un sbstring en un string

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
