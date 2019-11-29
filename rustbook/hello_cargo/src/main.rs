#[allow(unused_variables)]
fn main() {
    let x = 6;
    println!("The value of x is: {}", x);
    another_function(3,4);
}

fn another_function(x: i32, y: i32) {
    println!("The value of x is: {}", x);
    println!("The value of y is: {}", y);
}
