#[derive(Debug)]
struct Rectangle {
    width: u32,
    height: u32,
}

fn main() {
    let rect1 : Rectangle = Rectangle { width: 30, height: 50 };
    println!("The rect1 {:?} {:#?}", rect1, rect1);

    println!(
        "The area of the rectangle is {} square pixels.",
        rect1.area()
    );
}

impl Rectangle {
    fn area(&self) -> u32 {
        self.width * self.height
    }
}

