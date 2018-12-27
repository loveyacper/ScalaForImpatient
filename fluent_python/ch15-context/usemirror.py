from mirror import LookingGlass
import sys

with LookingGlass() as enterResult:
    sys.stdout.write(enterResult)
    sys.stdout.write("\n")
    sys.stdout.write("It is reversed")
    sys.stdout.write("\n")
    #a = 5 / 0

sys.stdout.write("It is normal\n")
