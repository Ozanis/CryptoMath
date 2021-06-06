use crate::operator::group::*;
use crate::operator::logical::*;
use crate::operator::modular::*;
use crate::operator::multiplicative::*;
use crate::operator::convolution::*;

pub type NativeVector = Vec<u8>;

impl Equals<NativeVector> for NativeVector {
	fn equals(&self, x : &NativeVector) -> bool {
		if self.len() != x.len() {
			return false;
		}
		for i in 0..self.len() {
			if self[i] != x[i] {
				return false;
			}
		}
		return true;
	}
}

impl Greater<NativeVector> for NativeVector {
	fn greater(&self, x : &NativeVector) -> bool {
		if self.len() > x.len() {
			return true;
		}
		if self.len() < x.len() {
			return false;
		}
		let mut i : usize = self.len() - 1;
		while i > 0 {
			if self[i] > x[i] {
				return true;
			}
			if self[i] < x[i] {
				return false;
			}
			i -= 1;
		}
		return false;
	}
}

impl GreaterEquals<NativeVector> for NativeVector {
	fn greater_equals(&self, x : &NativeVector) -> bool {
		if self.len() < x.len() {
			return false;
		}
		if self.len() > x.len() {
			return true
		}
		let mut i : usize = self.len() - 1;
		while i > 0 {
			if self[i] > x[i] {
				return true
			}
			if self[i] < x[i] {
				return false
			}
			i -=1;
		}
		return true
	}
}

impl Lesser<NativeVector> for NativeVector {
	fn lesser(&self, x : &NativeVector) -> bool {
		if self.len() < x.len() {
			return true;
		}
		if self.len() > x.len() {
			return false;
		}
		let mut i : usize = self.len() - 1;
		while i > 0 {
			if self[i] < x[i] {
				return true;
			}
			if self[i] > x[i] {
				return false;
			}
			i -= 1;
		}
		return false;
	}
}

impl LesserEquals<NativeVector> for NativeVector {
	fn lesser_equals(&self, x : &NativeVector) -> bool {
		if self.len() > x.len() {
			return false;
		}
		let mut i : usize = self.len() - 1;
		while i > 0 {
			if self[i] > x[i] {
				return false;
			}
			if self[i] < x[i] {
				return true;
			}
			i -= 1;
		}
		return true;
	}
}

impl Add<NativeVector> for NativeVector {	
    fn add(&mut self, x : &NativeVector, y : &NativeVector) -> &NativeVector {
        for i in 0..y.len() {
			self[i+1] = (self[i] + x[i] + y[i]) >> 1;
			self[i] ^= y[i] ^ x[i];
		}
		return self;
    }
}

impl Sub <NativeVector> for NativeVector {	
    fn sub(&mut self, x : &NativeVector, y : &NativeVector) -> &NativeVector {
		let mut borrow : u8 = 0;
		for i in 0..y.len() {
			borrow ^= x[i];
			self[i] = y[i] ^ borrow;
			borrow = ((borrow < y[i]) | (borrow > x[i])) as u8;
		}
		return self;
    }
}

impl TransformX<NativeVector> for NativeVector {
	fn transform_x(&mut self, x : &NativeVector, y : &NativeVector) -> &NativeVector {
		let mut w_x : u8 = 0;
		for i in 0..x.len() {
			w_x += y[i];
			self[i] = match x[i] > 0 {
				true => w_x,
				_ => 0
			}
		}
		return self
	}
}

impl TransformY<NativeVector> for NativeVector {
	fn transform_y(&mut self, x : &NativeVector, y : &NativeVector) -> &NativeVector {
		let mut w_y : u8 = 0;
		let mut pos : usize = self.len()-2;
		for i in (1..y.len()).rev() {
			w_y += y[i];
			self[pos] = match x[i] > 0 {
				true => w_y,
				_ => 0
			};
			pos -= 1;
		}
		return self
	}
}

impl Carry<NativeVector> for NativeVector{
	fn carry(&mut self) -> &NativeVector {
		let mut carry : u8 = 0;
		for i in 0..self.len() {
			carry = carry + self[i];
			self[i] = match carry > 0 {
				true => carry & 1,
				_ => 0
			};
			carry >>= 1;
		}
		return self
	}
}

impl ForwardTransfrom<NativeVector> for NativeVector {
	fn forward_transform(&mut self, x : &NativeVector, y : &NativeVector) -> &NativeVector{
		self.transform_x(x,y);
		self.transform_y(x,y);
		return self
	}
}

impl Mul <NativeVector> for NativeVector {	
    fn mul(&mut self, x : &NativeVector, y : &NativeVector) -> &NativeVector {
		self.forward_transform(x, y);
		self.carry();
		return self
	}
}

impl Sqr <NativeVector> for NativeVector {
	fn sqr(&mut self, x : &NativeVector) -> &NativeVector {
		return self;			
    }
}

/*
impl Exp<NativeVector> for NativeVector {
	fn exp(&self, &pow : Vec<u8>, &modulo : Vec<u8>) {
		for i in 1..pow.len() {
			if i & 1 == 1 {
				res = res * self % modulo;
			}
			self = self.sqr() % modulo;
		}
		return res;
	}

}*/

impl Rem <NativeVector> for NativeVector{
    fn rem(&mut self, x : &NativeVector, modulo : &NativeVector) -> &NativeVector {
		return self;
    }
}