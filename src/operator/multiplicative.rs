pub trait Sqr<T> {
	fn sqr(&mut self, x : &T) -> &T;
}

pub trait Exp<T> {
	fn exp(&mut self, x : &T, pow : &T) -> &T;
}