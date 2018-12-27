import io.Source
import java.io._
import collection.mutable.ArrayBuffer

// 9.1 reverse lines from file
val src = Source.fromFile("test9.1.txt") 
val reverseLines = src.getLines.toArray[String].reverse

val out = new PrintWriter("test9.1.reverse.txt")
reverseLines.map(out.println(_))
out.close

// 9.2 convert tab to spaces
val tabFile = Source.fromFile("test9.2.txt") 
val outSpaces = new PrintWriter("test9.2.space.txt")
tabFile.toList.map{c => if (c == '\t') " " * 4 else c}.foreach(outSpaces.print)
outSpaces.close

// 9.3  filter words those length > 10
Source.fromFile("test9.3.txt").mkString.split("""\s+""").filter(_.length > 8).distinct.map(println)

// 9.10
@SerialVersionUID(100L)
class Person(val name : String) extends Serializable { 
    val friends : ArrayBuffer[Person]  = new ArrayBuffer[Person]
  
    def addFriend(f : Person) = friends += f

    override def toString : String = {
        var str = "name = " + name + ", friends = "
        friends.map(str += _.name)
        str
    }
}

val p = new Person("bert")
p.addFriend(new Person("xuhan"))

val out10 = new ObjectOutputStream(new FileOutputStream("test9.10.txt"))
out10.writeObject(p)
out10.close

val in10 = new ObjectInputStream(new FileInputStream("test9.10.txt"))
val savedP = in10.readObject().asInstanceOf[Person]
println(savedP)

