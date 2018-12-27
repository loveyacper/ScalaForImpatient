// 5.3 - 5.4
class Time(val hours : Int, val minutes : Int) {
    def fromZero = hours * 60 + minutes

    def before(other : Time ) : Boolean = {
        if (hours < other.hours)
            return true;
        else if (hours > other.hours)
            return false;
        else
            return minutes < other.minutes;
    }
}

val now = new Time(23, 55)
val before = new Time(23, 35)
println("result (now < before) = " + now.before(before))
println("result (before < now) = " + before.before(now))

println("23:55 from Zero = " + now.fromZero)
println("23:35 from Zero = " + before.fromZero)

// 5.7
class Person(fullname : String) {
    private val names = fullname.split(" ")

    def firstName = names(0)
    def lastName = names(1)
}

val bert = new Person("Bert Young")

println("first name = " + bert.firstName)
println("last name = " + bert.lastName)

// 5.8
class Car(val manu : String, val model: String, val year : Int = -1, val name : String = "") {
    println("main constructor for Car")
}

val car1 = new Car("bmw", "x5")
val car2 = new Car("bmw", "x5", year = 1999)
val car3 = new Car("bmw", "x5", name = "BN666")

