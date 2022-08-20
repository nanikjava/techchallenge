The code
```
def multiply(num_a, num_b):
    if num_a == 0 or num_b == 0:
        return 0
    return num_a + multiply(num_a, num_b - 1)
```

As long as function `multiply(..)` is called with non-negative number for `num_b` parameter the function works fine.

A non-positive number send to `num_b` parameter will throw a maximum recursion error like so

```
    print(multiply(1,-2))
```


