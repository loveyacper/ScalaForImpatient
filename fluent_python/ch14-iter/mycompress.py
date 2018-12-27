import itertools

heights = [160, 162, 164, 168, 170, 175, 180] 
selector = [True, False, False, False, True, True]

def mycompress1(iterV, iterSelect):
    return (d for d, flag in zip(iterV, iterSelect) if flag)

def mycompress(iterV, iterSelect):
    ''' fool impl '''
    index = 0
    for flag in iterSelect:
        v = None
        try:
            if flag:
                v = yield next(iterV)
            else:
                next(iterV)

        except TypeError:
            if index == len(iterV):
                return

            if flag:
                v = iterV[index]

            index = index + 1

        except StopIteration:
            return


        if v is not None:
            yield v


print('itertools.compress')
for e in itertools.compress(heights, selector):
    print(e)

print('\nmycompress')
for e in mycompress(heights, selector):
    print(e)

print('\nmycompress1')
for e in mycompress1(heights, selector):
    print(e)

