// Make a "keyword" unless, similar to !if 

def unless(cond: => Boolean)(block: => Unit) { 
    if (!cond) { 
        block 
    }   
}   
    
unless (5 < 3) { 
    println("You can not see this line unless 5 < 3")
}
