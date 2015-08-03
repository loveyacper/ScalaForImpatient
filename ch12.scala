// 12.1
def   values(fun : (Int)=> Int, low : Int, high : Int) = { 
    (low to high).map(x => (x, fun(x)))
}

println(values(x=> x*x, -2, 2)) 

// 12.2
def  maxValue[T <% Ordered[T]](arr : Array[T]) = { 
    arr.reduceLeft((a : T, b : T) => if (a > b) a else b)
}

println(maxValue(Array(1, 3, 5, 2)))

// 12.3
def fact(value : Int) = { 
    if (value <= 0)
        1   
    else
        (1 to value).reduceLeft((a : Int, b : Int) => a * b)
}
println(fact(0))
println(fact(3))

// 12.4
def fact2(value : Int) = { 
    (1 to value).foldLeft(1)((a : Int, b : Int) => a * b)
}
println(fact2(0))
println(fact2(3))

// 12.5
def largest(fun : (Int)=>Int, inputs : Seq[Int]) = inputs.map(fun(_)).max

println(largest(x => 10 * x - x * x, 1 to 10))

// 12.6
def largestParam(func : (Int)=>Int, inputs : Seq[Int]) = { 
    inputs.map((m) => (m, func(m))).reduceLeft((x,y)=>if (x._2 > y._2) x else y)._1
}
    
println(largestParam(x => 10 * x - x * x, 1 to 10))

// 12.7
def  adjustToPairs(fun : (Int, Int) => Int) = (p : (Int, Int)) => fun(p._1, p._2)

val pairs = (1 to 5) zip (6 to 10) 
println(pairs.map(adjustToPairs(_ + _)))

// 12.10 Make a "keyword" unless, similar to !if 
def unless(cond: => Boolean)(block: => Unit) { 
    if (!cond) { 
        block 
    }   
}   
    
unless (5 < 3) { 
    println("You can not see this line unless 5 < 3")
}
