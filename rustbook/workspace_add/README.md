# package & crate

  crate是binary或lib。package是提供一组功能的多个crates。package中包含了一个Cargo.toml文件，描述如何去构建crates。
rust编译器从一个特殊的源文件——crate root开始编译；这个源文件也构成了该crate的root module。
package可以包含0或1个lib crate，多个binary crates。package最少要包含一个crate，可以是lib也可以是binary；

  Cargo认为src/main.rs是一个与package同名的binary crate的root。同样的，如果package目录包含src/lib.rs,Cargo认为这个package包含了一个与package同名的lib crate，且src/lib.rs是crate root。Cargo把root文件传递给rustc去构建lib或binary。

# module

  Module便于将crate内的代码组织的更有可读性和复用性。Modules也控制可见性，类似其它编程语言的private/public机制。

# rustc --explain E0432

  An import was unresolved.

Erroneous code example:

```
use something::Foo; // error: unresolved import `something::Foo`.
```

Paths in `use` statements are relative to the crate root. To import items
relative to the current and parent modules, use the `self::` and `super::`
prefixes, respectively. Also verify that you didn't misspell the import
name and that the import exists in the module from where you tried to
import it. Example:

```
use self::something::Foo; // ok!

mod something {
    pub struct Foo;
}
```

Or, if you tried to use a module from an external crate, you may have missed
the `extern crate` declaration (which is usually placed in the crate root):

```
extern crate core; // Required to use the `core` crate

use core::any;
