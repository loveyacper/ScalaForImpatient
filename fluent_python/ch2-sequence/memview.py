import array

# short array, 10 bytes
numbers = array.array('h', [-2, -1, 0, 1, 2])
memv = memoryview(numbers)
print('len(memv)', len(memv))

memv_oct = memv.cast('B')
print('memv_oct.tolist()', memv_oct.tolist())
memv_oct[5] = 4

print('numbers', numbers)
