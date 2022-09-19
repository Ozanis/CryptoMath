pub trait GCD<T> {
    fn gcd(&mut self, x : &T, y : &T) -> &T;
}

pub trait Modulo<T> {
    fn modulo(&mut self) -> &T;
}

pub trait SqrMod<T> {
	fn sqr_mod(&mut self, modulo : &T) -> &T;
}

pub trait MultMod<T> {
	fn sqr_mod(&mut self, x : &T, modulo : &T) -> &T;
}

pub trait InvMod<T> {
	fn inv_mod(&mut self, modulo : &T) -> &T;
}

pub trait PowMod<T> {
	fn pow_mod(&mut self, x : &T, pow : &T, modulo : &T) -> &T;
}