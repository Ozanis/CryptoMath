pub trait Rem<T> {
    fn rem(&mut self, x : &T, modulo : &T) -> &T;
}

pub trait Gcd<T> {
    fn gcd(&mut self, x : &T, y : &T) -> &T;
}