use crate::operator::group::*;
use crate::operator::logical::*;
use crate::operator::modular::*;
use crate::operator::multiplicative::*;


pub type NativeVector = Vec<u8>;


impl Add<NativeVector> for NativeVector {	
    fn add(&mut self, x : &NativeVector, y : &NativeVector) -> &NativeVector {
        let mut carry : u8 = 0;
        for i in 0..y.len() {
			self[i] = carry ^ y[i] ^ x[i];
			carry = (carry + x[i] + y[i]) >> 1;
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


impl Mul <NativeVector> for NativeVector {	
    fn mul(&mut self, x : &NativeVector, y : &NativeVector) -> &NativeVector {
		let x_length : usize = x.len()-1;
		let y_length : usize = y.len()-1;
		let total_length : usize = x_length + y_length;
		let mut weight : u8 = 0;
		
		for i in 0..y.len() {
			weight = (weight >> 1) + x[i];
			self[i] = weight & y[i];
		}
		
		self[y.len()] = weight >> 1;

		weight = 0;
		for i in 1..y.len()-1 {
			weight += x[i];
			self[total_length-i] = weight & y[i];
		}

		weight = 0;
		for i in y.len()..total_length {
			self[i] = (self[i] + weight) & 1;
			weight = self[i] >> 1;
		}

		return self;
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