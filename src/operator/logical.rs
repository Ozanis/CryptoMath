pub trait Equals<T> {
	fn equals(&self, x : &T) -> bool;
}

pub trait Greater<T> {
	fn greater(&self, x : &T) -> bool;
}

pub trait GreaterEquals<T> {
	fn greater_equals(&self, x : &T) -> bool;
}

pub trait Lesser<T> {
	fn lesser(&self, x : &T) -> bool;
}

pub trait LesserEquals<T> {
	fn lesser_equals(&self, x : &T) -> bool;
}