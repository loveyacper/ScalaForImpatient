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
    var res : Long = 1
    for (c <- s)
        res *= c.toLong
    res
}
println(multiunicode("Hello"))
println()

// 2.7 
println("Hello".foldLeft(1L)(_*_.toLong))

// 2.9 
def prod(str : String) : Long  = { 
    if (str.length == 0)
        0   
    else if (str.length == 1)
        str(0).toLong
    else
        str(0)*prod(str.substring(1))
}
println(prod("Hello"))

// 2.10
def power(x : Int, n : Int) : Float = { 
    if (n == 0) {
        1   
    }   
    else if (n < 0) {
        1 / power(x, -n) 
    }   
    else {
        if (n % 2 == 0)
            power(x, n / 2) * power(x, n / 2)
        else
            x * power(x, n - 1)
    }   
}

println(power(2, 0)) 
println(power(2, -2))
println(power(2, 2)) 
println(power(2, -3))
println(power(2, 3)) 

