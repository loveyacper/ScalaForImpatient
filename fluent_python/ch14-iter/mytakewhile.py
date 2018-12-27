import itertools

def mytakewhile(cond, gen):
    def inner(g):
        while True:
            try:
                v = next(g)
            except StopIteration:
                return
            except TypeError:
                if index == len(g):
                    return
                    
                print("index = " + index)

                v = g[index]
                index = index + 1

            if v is None:
                return

            if cond(v):
                yield v
            else:
                return

    return inner(gen)


for e in mytakewhile(lambda v : v <= 3, itertools.count(1, .5)):
    print(e) 
