use crate::vector::vector::*;


impl Add<NativeVector> for NativeVector{
    fn add(mut self, x : &NativeVector) -> NativeVector{
        let mut carry : i8 = 0;
        for i in 0..x.vector.len() {
            carry += self.vector[i] + x.vector[i];
            self.vector[i] = carry & 1;
            carry >>= 1;
        }
        return self;
    }
}


impl Sub<NativeVector> for NativeVector {
    fn sub(mut self, x : &NativeVector) -> NativeVector{
        let mut borrow : i8 = 0;
        for i in 0..x.vector.len() {
            self.vector[i] ^= borrow;
            borrow = (self.vector[i] < x.vector[i]) as i8;
            self.vector[i] ^= x.vector[i];
        }
        //self.vector[x.vector.len()-1] ^= borrow;
        return self;
    }
}


/*
impl Rem<NativeVector> for NativeVector{
    fn rem(mut self, x : &NativeVector) -> NativeVector {
        for i in 0.. x.vector.len() {
            self.vector[i] ^= x.vector[i];
        }
        for i in x.vector.len()..self.vector.len() {
            self.vector[i] = 0;
        }
        return self;
    }
}*/


/*
impl Div<NativeVector> for NativeVector {
    fn div(self, x : &NativeVector ) -> NativeVector{
        return self;
    }
}*/

/*
impl Mul<NativeVector> for NativeVector {

	fn mul(self, x : &NativeVector) -> NativeVector {
        return self.fast_walsh_hadamard().convolution(&x.fast_walsh_hadamard()).fast_walsh_hadamard().radix();
    }
}*/


impl Sqr<NativeVector> for NativeVector {
	fn sqr(self) -> NativeVector{
		let squared : &NativeVector = &NativeVector{vector : self.vector.clone()};
		return self.fast_walsh_hadamard().convolution(&squared).radix();
    }
}

/*
impl Exp<NativeVector> for NativeVector {

	fn reduce_exp(mut self, pow : NativeVector, modulo : NativeVector) -> NativeVector {
		return self;
	}

	fn exp(mut self, pow : NativeVector, modulo : NativeVector) -> NativeVector {
		let mut res : NativeVector = NativeVector{vector : self.vector.clone()};
		for i in 1..pow.vector.len() {
			if i & 1 == 1 {
				res = res * self % modulo;
			}
			self = self.sqr() % modulo;
		}
		return res;
	}

}*/