''' similar to std::bind'''
from operator import mul
from functools import partial

triple = partial(mul, 3)

print(triple(7))
print(tuple(c for c in map(triple, range(0, 5))))

def echo(a, b, c):
    print(a, b, c)

helloworld = partial(echo, b = 'hello ', c = 'world!') 
helloworld('HAHA: ')

