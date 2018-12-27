def AriProGen(begin, step, end = None):
    result = type(begin + step)(begin)
    forever = end is None
    index = 0
    while forever or result < end:
        yield result
        index += 1
        result = begin + step * index
    pass

if __name__ == '__main__':
    gen = AriProGen(1, 1/3, 2)
    for e in gen:
        print(e)
