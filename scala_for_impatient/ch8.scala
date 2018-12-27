// 8.1
class BankAccount(initBalance : Double) {
    private var balance = initBalance

    final def getBalance = balance

    def deposit(amount : Double) = { balance += amount; balance }
    def withdraw(amount : Double) = { balance -= amount; balance }
}

class CheckingAccount(initBalance : Double) extends BankAccount(initBalance) {
    override def deposit(amount : Double) = { if (amount > 1) super.deposit(amount - 1) else  getBalance }
    override def withdraw(amount : Double) = { if (getBalance > 1) super.withdraw(amount + 1) else  getBalance }
}

val check = new CheckingAccount(10)
check.deposit(10)
check.withdraw(10)
println(check.getBalance)

// 8.4
import scala.collection.mutable.ArrayBuffer

abstract class Item {
    def price : Double
    def description : String

    override def toString() = "price = " + price + ", desc = " + description
}

class SimpleItem(override val price : Double, override val description :  String) extends Item {
}

class Bundle extends Item {
    private val items =  new ArrayBuffer[Item]()

    def addItem(it : Item) = {
        items += it
    }

    def clear = items.clear

//  override def price : Double = items.map(_.price).sum
    override def price : Double = {
        (0.toDouble /: items) { (totalPrice, item) => totalPrice + item.price }
    }

    override def description :  String = {
        ("Bundle " /: items) { (desc, item) => desc + item.description + " " }
    }
}

val item = new SimpleItem(10, "simple")
println(item)

val bundle = new Bundle()

bundle.addItem(item)
val item2 = new SimpleItem(20, "simple2")
bundle.addItem(item)
println(bundle)

bundle.clear 
println(bundle)
