use big_hex::BigHex;

trait Arithmetical {
    fn add(&mut self,  bigHex: &BigHex);
    fn sub(&self) -> String;
}

impl Arithmetical for BigHex {	
    fn add(&mut self, bigHex : &BigHex) -> &BigHex {
        for i in 0..self.vector.len() {
			self.vector[i+1] = (self.vector[i] + big_hex.vector[i]) >> 1;
			self.vector[i] ^= self.vector[i] ^ big_hex.vector[i];
		}
		return self;
    }

    fn sub(&mut self, bigHex : &BigHex) -> &BigHex {
		let mut borrow : u8 = 0;
		for i in 0..y.len() {
			borrow ^= self.vector[i];
			self[i] = bigHex.vector[i] ^ borrow;
			borrow = ((borrow < bigHex.vector[i]) | (borrow > self.vector[i])) as u64;
		}
		return self;
    }
}

// impl Mul for BigHex {	
//     fn mul(&mut self, vector : &BigHex) -> &BigHex {
// 		return self
// 	}
// }

