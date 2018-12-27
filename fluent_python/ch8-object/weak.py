import weakref

def bye(): 
    print('Goodbye...')

a_set = {0, 1}

wref = weakref.ref(a_set)
fin = weakref.finalize(a_set, bye)

print('wref:', wref)
print('fin.alive:', fin.alive)

print('wref()', wref())

del a_set

print('wref', wref)
print('wref()', wref())
