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
