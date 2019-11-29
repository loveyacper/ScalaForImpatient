#[allow(unused_variables)]
fn main() {
    //test_vec();
    test_hashmap();
}


fn test_hashmap() {
    use std::collections::HashMap;

    let mut scores = HashMap::new();

    scores.insert(String::from("Blue"), 10);
    scores.insert(String::from("Yellow"), 50);

    let team_name = String::from("Blue");
    let score = scores.get(&team_name);
    if let Some(s) = score {
        println!("The score {}", s);
    }

    for (key, value) in &scores {
        println!("{}: {}", key, value);
    }

    ////

    let field_name = String::from("Favorite color");
    let field_value = String::from("Blue");

    let mut map = HashMap::new();
    let mut what = map.insert(&field_name, &field_value);
    what = map.insert(&field_name, &field_value);
    if let Some(v) = what {
        println!("The insert res {}", v);
    }
}

fn test_vec() {
    //let v : Vec<i32>  = vec![1, 2, 3];
    let mut v: Vec<i32> = Vec::new();
    for i in 1..4 {
        v.push(i);
    }

    for i in &mut v {
        *i += 100;
    }

    let third: & mut i32 = & mut v[2];
    *third += 1000;

    println!("The third element is {}", third);
    println!("The 3rd element is {}", v[2]);

    match v.get(2) {
        Some(third) => println!("The third element is {}", third),
        None => println!("There is no third element."),
    }
    println!("Hello, world!");
}
