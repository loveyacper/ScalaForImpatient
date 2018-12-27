from mirror_gen import looking_glass
import sys

with looking_glass(age = 32, name = "bert") as enterResult:
    sys.stdout.write(enterResult)
    sys.stdout.write("\n")
    sys.stdout.write("It is reversed")
    sys.stdout.write("\n")
    a = 5 / 0

sys.stdout.write("It is normal\n")
