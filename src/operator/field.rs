pub trait Inverse<T> {
    fn inverse(&mut self, x : &T) -> &T;
}

pub trait Order<T> {
    fn order(&self) -> u64;
}