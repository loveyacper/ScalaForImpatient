import time
import functools

def func_id(func, *args, **kwargs):
    arg_lst = []

    if args:
        arg_lst.append(','.join(repr(a) for a in args))

    if kwargs:
        pairs = ['%s=%r' % (k, w) for k, w in kwargs.items()]
        arg_lst.append(', '.join(pairs))

    return func.__name__ + ', '.join(arg_lst)

# decorator
def cache(size = 16):
    lru = {}
    def cache_wrap(func):
        @functools.wraps(func)
        def inner_cache(*args, **kwargs):
            tid = func_id(func, *args, **kwargs)
            #print(tid)
            if tid in lru:
                return lru[tid]
            else:
                res = func(*args)

            if len(lru) < size:
                lru[tid] = res

            return res

        return inner_cache
    return cache_wrap

#decorator另外一种写法
def cache2(func, size):
    lru = {}

    @functools.wraps(func)
    def inner_cache(*args, **kwargs):
        tid = func_id(func, *args, **kwargs)
        #print(tid)
        if tid in lru:
            return lru[tid]
        else:
            res = func(*args)

        if len(lru) < size:
            lru[tid] = res

        return res

    return inner_cache

def cache2_wrapper(size = 16):
    return functools.partial(cache2, size=size)
   
#@cache(20) 使用第一种cache 
@cache2_wrapper(20) #使用第二种cache
def fib_cache(n):
    if n < 2:
        return n
    return fib_cache(n - 1) + fib_cache(n - 2) 

def fib(n):
    if n < 2:
        return n
    return fib(n - 1) + fib(n - 2)

n = 33

# with cache
start = time.time()
fib_cache(n)
used = time.time() - start
print("%-12s used %.6f" % (fib_cache.__name__, used))

# without cache
start = time.time()
fib(n)
used = time.time() - start
print("%-12s used %.6f" % (fib.__name__, used))

