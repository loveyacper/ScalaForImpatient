import functools
import operator

# sum (list[1])

initial = 0

#my_list = []
my_list = [[1, 2, 3], [40, 50, 60], [9, 8, 7]]

print(functools.reduce(lambda a, b: a + b, (sub[1] for sub in my_list), initial))

print(functools.reduce(operator.add, (sub[1] for sub in my_list), initial))

print(sum(sub[1] for sub in my_list))

# 60 = 2 + 50 + 8
