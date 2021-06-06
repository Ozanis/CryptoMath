pub trait ForwardTransfrom<T> {
	fn forward_transform(&mut self, x : &T, core : &T) -> &T;
}

pub trait TransformX<T> {
	fn transform_x(&mut self, x : &T, core : &T) -> &T;
}

pub trait TransformY<T> {
	fn transform_y(&mut self, x : &T, core : &T) -> &T;
}

pub trait InverseTransfrom<T> {
	fn inverse_transform(&mut self, x : &T, core : &T) -> &T;
}

pub trait Carry<T> {
	fn carry(&mut self) -> &T;
}