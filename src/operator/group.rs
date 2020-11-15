pub trait Add<T> {
    fn add(&mut self, x : &T, y : &T) -> &T;
}

pub trait Sub<T> {
    fn sub(&mut self, x : &T, y : &T) -> &T;
}

pub trait Mul<T> {
    fn mul(&mut self, x : &T, y : &T) -> &T;
}
