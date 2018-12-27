from functools import singledispatch
from collections import abc

import numbers
import html

@singledispatch
def htmlize(obj):
    print(htmlize.__name__)
    content = html.escape(repr(obj))
    return '<pre>{}</pre>'.format(content)

@htmlize.register(str)
def _(text):
    print("str")
    content = html.escape(text).replace('\n', '<br>\n')
    return '<p>{0}</p>'.format(content)

@htmlize.register(numbers.Integral)
def _(n):
    print("int")
    return '<pre>{0} (0x{0:x})</pre>'.format(n)

@htmlize.register(tuple)
@htmlize.register(abc.MutableSequence)
def _(seq):
    print("tuple/seq")
    inner = '</li>\n<li>'.join(htmlize(item) for item in seq)
    return '<ul>\n<li>' + inner + '</li>\n</ul>'

htmlize(1)
htmlize("hello")
htmlize(("hello", 1))

print(htmlize.__code__)
print(htmlize.__code__.co_freevars)
