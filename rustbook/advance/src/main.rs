#![allow(unused_variables)]
fn main() {
    unsafe {
        test_unsafe_rawp();
    }
}

unsafe fn test_unsafe_rawp() {
    let mut num = 5;

    let r1 = &num as *const i32; // const int* const r1 = &num;
    let r2 = &mut num as *mut i32; // int* const r2 = &num;
    *r2 += 1;
    println!("r1, r2, num {}, {}, {}", *r1, *r2, num);
}

