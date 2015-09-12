// 3.1 
import util.Random

def randArray(n : Int) = { 
    val a = new Array[Int](n)
    println(a.length)
    for (i <- 0 until a.length)
        a(i) = Random.nextInt(n)
    a   
}

println(randArray(5).mkString("(", "-", ")"))

// 3.2 
def swapElem(arr : Array[Int]) = { 
    for (i <- 0 until (arr.length, 2)) {
        if (i + 1 < arr.length) {
            val tmp = arr(i)
            arr(i) = arr(i + 1)
            arr(i + 1) = tmp 
        }   
    }   
}

// 3.3
def yieldSwapElem(a : Array[Int]) = {
    for (i <- 0 until a.length) yield {
        if (i % 2 == 0) {
            if (i + 1 < a.length)
                a(i + 1)
            else
                a(i)
        }
        else {
            a(i - 1)
        }
    }
}

val arr1 = Array(1, 2, 3, 4, 5)
val arr2 = yieldSwapElem(arr1)
println(arr2.mkString("-"))

// 3.4
def  splitArray(a : Array[Int]) = {
    a.filter(_ > 0) ++ a.filter(_ <= 0)
}

val arr3 = Array(1, -2, 3, 0, -5)
val res = splitArray(arr3)
println(res.mkString(","))

// 3.5
def  averageDoubleArr(a : Array[Double]) : Double = {
    a.sum / a.size
}
