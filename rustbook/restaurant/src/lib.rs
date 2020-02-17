#[cfg(test)]
mod tests {
    #[test]
    fn it_works() {
        assert_eq!(2 + 2, 4);
    }
}


//eat_at_restaurant 和 front_of_house 是兄弟，所以后者不必要pub
mod front_of_house {
    pub mod hosting {
        pub fn add_to_waitlist() {}

        fn seat_at_table() {}
    }

    mod serving {
        fn take_order() {}

        fn serve_order() {}

        fn take_payment() {}
    }
}

pub fn eat_at_restaurant() {
    // Absolute path, 使用crate，将从src/lib.rs开始寻找定义
    crate::front_of_house::hosting::add_to_waitlist();

    //restaurant::front_of_house::hosting::add_to_waitlist();

    // Relative path,因为front_of_house和eat_at_restaurant在同一个crate定义
    front_of_house::hosting::add_to_waitlist();
}

// demo: using super...
//
fn serve_order() {}

mod back_of_house {
    fn fix_incorrect_order() {
        cook_order();
        super::serve_order();
    }

    fn cook_order() {}
}

