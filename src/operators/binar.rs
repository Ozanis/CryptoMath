use big_hex::BigHex;

trait Binar {
    fn inv(&mut self) -> &BigHex;
	fn or(&mut self,  bigHex: &BigHex) -> &BigHex;
    fn xor(&mut self,  bigHex: &BigHex) -> &BigHex;
	fn and(&mut self,  bigHex: &BigHex) -> &BigHex;
}

impl Binar for BigHex {	
    fn inv(&mut self) -> &BigHex {
        for i in 0..self.vector.len() {
			self.vector[i] = !self.vector[i]
		}
		return self;		
	}

	fn or(&mut self,  bigHex: &BigHex) -> &BigHex {
        for i in 0..self.vector.len() {
			self.vector[i] |= self.bigHex[i]
		}
		return self;		
	}

    fn xor(&mut self,  bigHex: &BigHex) -> &BigHex {
        for i in 0..self.vector.len() {
			self.vector[i] ^= self.bigHex[i]
		}
		return self;		
	}

	fn and(&mut self,  bigHex: &BigHex) -> &BigHex {
        for i in 0..self.vector.len() {
			self.vector[i] &= self.bigHex[i]
		}
		return self;		
	}
    
}
