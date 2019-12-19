use std::thread;
use std::time::Duration;
use std::sync::mpsc;

fn main() {
    let v = vec![1, 2, 3];

    let handle = thread::spawn(move || {
        println!("Here's a vector: {:?}", v);
    });

    handle.join().unwrap();

    let (tx, rx) = mpsc::channel::<String>();
    //let (tx, rx) = mpsc::channel::<i32>();
    thread::spawn(move || {
        let val = String::from("hi");
        tx.send(val).unwrap();
    });
}

