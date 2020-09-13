use crate::vector::vector::*;
use crate::binary::binary::{reverse, log};


impl Convolution<NativeVector> for NativeVector {
	fn radix(mut self) -> NativeVector {
		let len : usize = log(self.vector.len());
		for i in 0..self.vector.len() {
				self.vector[i] = self.vector[reverse(i, len)];
			}
			return self;
		}

	fn fast_walsh_hadamard(mut self) -> NativeVector {
		let mut h : usize = 1;
		let mut u : i8;
		let mut t : i8;
		while h < self.vector.len()  {
			for i in (0..self.vector.len()).step_by(h*2) {
				for j in i..i+h {
					u = self.vector[j];
					t = self.vector[j+h];
					self.vector[j] = u + t;
					self.vector[j+h] = u - t;
				}
			}
			h *= 2;
		}
		return self;
	}

	fn convolution(mut self, x : &NativeVector) -> NativeVector {
		for i in 0..self.vector.len()/2 {
			self.vector[i] &= x.vector[i];
			self.vector[i+1] ^= x.vector[i];
		}
		return self;
	}

}


