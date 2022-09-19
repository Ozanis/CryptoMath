pub trait Add<T> {
    fn add(&mut self, x : &T, y : &T) -> &T;
}

pub trait Sub<T> {
    fn sub(&mut self, x : &T, y : &T) -> &T;
}

pub trait Sqr<T> {
	fn sqr(&mut self, x : &T) -> &T;
}

pub trait Mult<T> {
	fn mult(&mut self, x : &T) -> &T;
}

pub trait Pow<T> {
	fn pow(&mut self, x : &T, pow : &T) -> &T;
}