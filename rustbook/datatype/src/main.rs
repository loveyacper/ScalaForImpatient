fn main() {
    // "42" is struct String from crate std, module string.
    // pub fn parse<F>(&self) -> Result<F, <F as FromStr>::Err>
    //
    let guess: u32 = "42".parse().expect("Not a number!");
    let guess2 = "43".parse::<u32>().expect("Not a number!");
    let guess3 = "44".parse::<u32>().unwrap();
    println!("Hello, guess {}", guess);
    println!("Hello, guess2 {}", guess2);
    println!("Hello, guess3 {}", guess3);

    // str
    let ss: &str = "hello str";
    println!("a str : {}", ss);

    // char type
    let c: char = 'z';
    println!("ch is {}", c);

    // tuple
    let tup: (i32, f64, u8) = (500, 6.4, 1);
    let (x, y, z) = tup;
    println!("x {} y {} z {}", x, y, z);

    let five_hundred = tup.0;
    println!("five_hundred {}", five_hundred);

    // array
    let a: [i32; 5] = [1, 2, 3, 4, 5];
    println!("arr_a {:?}", a);

    let b = [3; 5]; // same as [3, 3, 3, 3, 3];
    println!("arr_b {:?}", b);
}
