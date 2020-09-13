pub struct Vector<T> {
	pub vector: Vec<T>,
}


pub type NativeVector = Vector<i8>;


pub trait Add<Vector> {
    fn add(self, x : &NativeVector) -> NativeVector;
}

pub trait Sub<Vector> {
    fn sub(self, x : &NativeVector) -> NativeVector;
}

pub trait Rem<Vector> {
    fn rem(self, x : &NativeVector) -> NativeVector;
}

pub trait Div<Vector> {
    fn div(self, x : &NativeVector ) -> NativeVector;
}

pub trait Mul<Vector> {
    fn mul(self, x : &NativeVector ) -> NativeVector;
}

pub trait Convolution<Vector> {
	fn radix(self) -> Vector;
	fn fast_walsh_hadamard(self) -> Vector;
	fn convolution(self, x : &Vector) -> Vector;
}

pub trait Sqr<Vector> {
	fn sqr(self) -> Vector;
}

pub trait Exp<Vector> {
	fn reduce_exp(self, pow : &Vector, modulo : &Vector) -> Vector ;
	fn exp(self, pow : &Vector, modulo : &Vector) -> Vector;
}

pub trait Eq<Vector> {
	fn eq(&self, x : &Vector) -> bool;
}

pub trait Gt<Vector> {
	fn gt(&self, x : &Vector) -> bool;
}

pub trait Gte<Vector> {
	fn gte(&self, x : &Vector) -> bool;
}

pub trait Ls<Vector> {
	fn ls(&self, x : &Vector) -> bool;
}

pub trait Lse<Vector> {
	fn lse(&self, x : &Vector) -> bool;
}