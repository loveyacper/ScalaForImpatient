# -*- coding:utf-8 -*-
'''函数对象有个 __defaults__ 属性，它的值是一个元组，里面保存着定位参数和关键字
参数的默认值。仅限关键字参数的默认值在 __kwdefaults__ 属性中。然而，参数的名
称在 __code__ 属性中，它的值是一个 code 对象引用，自身也有很多属性。'''

def clip(text:str, max_len:'int > 0'=80, *tupleargs) -> str:
    """在max_len前面或后面的第一个空格处截断文本 """
    end = None
    if len(text) > max_len:
        space_before = text.rfind(' ', 0, max_len)
        if space_before >= 0:
            end = space_before
        else:
            space_after = text.rfind(' ', max_len)
            if space_after >= 0:
                end = space_after

    if end is None: # 没找到空格
        end = len(text)
         
    return text[:end].rstrip()

print(clip.__defaults__)
print(clip.__kwdefaults__)
print(clip.__code__)
print(clip.__code__.co_varnames)
print(clip.__code__.co_argcount)

def fuck(a, *, keyonly):
    pass

print("use inspect module")
from inspect import signature
sig = signature(clip)
print(sig)
for name, param in sig.parameters.items(): 
    print(param.kind, ':', name, '=', param.default)

sig = signature(fuck)
print(sig)
for name, param in sig.parameters.items(): 
    print(param.kind, ':', name, '=', param.default)
