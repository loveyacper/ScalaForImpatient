// 2.1
def  signum(v : Int) = {
    if (v > 0) 1
    else if (v < 0) -1
    else 0
}
println(signum(0))
println(signum(5))
println(signum(-5))
println()


// 2.4
for (i <- 10 to (0, -1))
    println(i)

println()

// 2.5
def countdown(n : Int) = {
    for (i <- n to (0, -1))
        println(i)
}
countdown(3)
println()

// 2.6
def multiunicode(s : String) = {
    var res = 1
    for (c <- s)
        res *= c
    res
}
println(multiunicode("Hello"))
println()
