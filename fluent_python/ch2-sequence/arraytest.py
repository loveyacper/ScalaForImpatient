from array import array
from random import random
from os import remove

size = 10**5
floats = array('d', (random() for i in range(size)))
print(floats[-1])

fp = open('floats.bin', 'wb')
floats.tofile(fp)
fp.close()

floats2 = array('d')
fp = open('floats.bin', 'rb')
floats2.fromfile(fp, size)
fp.close()

print(floats2[-1])
assert (floats2 == floats)

remove('floats.bin')
