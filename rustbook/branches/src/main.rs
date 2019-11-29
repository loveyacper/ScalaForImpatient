
struct User {
    username: String,
    active: bool,
}


fn build_user(username: String) -> User {
    User {
        username,
        active: true,
    }
}

// Tuple structs are useful when you want to give the whole tuple a name and
// make the tuple be a different type from other tuples

#[allow(unused_variables)]
fn main() {
    let mut bert : User = User {
        username: String::from("someusername123"),
        active: true,
    };
    let young : User = build_user(String::from("young"));
    bert.username = String::from("bertyoung");
    println!("user {}, {}", bert.username, bert.active);
    println!("young user {}, {}", young.username, young.active);

    let hello : String = String::from("Hello, world!");
    let ref_hello : &String = &hello;
    println!("the ref value is: {}", ref_hello);
    println!("the ref value is: {}", ref_hello);

    let a : [i32; 3] = [1,2,3];
    for element in a.iter().rev() {
        println!("the value is: {}", element);
    }
    let slice : &str = first_word(&hello);
    println!("the first is: {}", slice);
}

fn first_word(s: &String) -> &str {
    let bytes = s.as_bytes();

    for (i, &item) in bytes.iter().enumerate() {
        if item == b' ' {
            return &s[0..i];
        }
    }

    &s[..]
}

