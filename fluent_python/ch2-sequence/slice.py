l = list(range(10))
print(l)

del l[5:7]
print(l)

l[3::2] = [11, 22, 33]
print(l)

#l[2:5] = 100 右边必须是可迭代的
#print(l)
    
l[2:5] = [100]
print(l)
