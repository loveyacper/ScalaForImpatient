
#[derive(Debug)]
enum IpAddrKind {
    V4,
    V6,
}

fn main() {
    let four : IpAddrKind = IpAddrKind::V4;
    let six = IpAddrKind::V6;
    println!("Hello, world! {:?}", four);

    let mut some_u8_value = Some(9u8);
    some_u8_value = match some_u8_value {
        Some(3) => {
            println!("three");
            Some(3u8)
        },

        _ => Some(3u8),
    };

    if let Some(3) = some_u8_value {
        println!("three");
    }
}
