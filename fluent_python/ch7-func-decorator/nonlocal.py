# 当 count 是数字或任何不可变类型时，count += 1 语句的作用其实与 count = count + 1 一样
# 因此，我们在 averager 的定义体中为 count 赋值了，这会把
# count 变成局部变量。total 变量也受这个问题影响。

def make_averager():
    count = 0
    total = 0
    def averager(new_value):
        nonlocal count, total # important
        count += 1
        total += new_value
        avg = total / count
        print(avg)
        return avg

    return averager

avg = make_averager()
avg(10)
avg(20)
avg(30)
