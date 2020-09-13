use crate::vector::vector::{Eq,Gt,Gte,Ls,Lse};
use crate::vector::vector::NativeVector;


impl Eq<NativeVector> for NativeVector {
	fn eq(&self, x : &NativeVector) -> bool {
		if self.vector.len() != x.vector.len() {
			return false;
		}
		for i in 0..x.vector.len() {
			if self.vector[i] != x.vector[i] {
				return false;
			}
		}
		return true;
	}
}

impl Gt<NativeVector> for NativeVector {
	fn gt(&self, x : &NativeVector) -> bool {
		if self.vector.len() > x.vector.len() {
			return true;
		}
		if self.vector.len() < x.vector.len() {
			return false;
		}
		for i in 0..x.vector.len() {
			if self.vector[i] > x.vector[i] {
				return true;
			}
			if self.vector[i] < x.vector[i] {
				return false;
			}
		}
		return false;
	}
}

impl Gte<NativeVector> for NativeVector {
	fn gte(&self, x : &NativeVector) -> bool {
		if self.vector.len() < x.vector.len() {
			return false;
		}
		if self.vector.len() > x.vector.len() {
			return true
		}
		for i in 0..x.vector.len() {
			if self.vector[i] > x.vector[i] {
				return true
			}
			if self.vector[i] < x.vector[i] {
				return false
			}
		}
		return true
	}
}

impl Ls<NativeVector> for NativeVector {
	fn ls(&self, x : &NativeVector) -> bool {
		if self.vector.len() < x.vector.len() {
			return true;
		}
		if self.vector.len() > x.vector.len() {
			return false;
		}
		for i in 0..x.vector.len() {
			if self.vector[i] < x.vector[i] {
				return true;
			}
			if self.vector[i] > x.vector[i] {
				return false;
			}
		}
		return false;
	}
}

impl Lse<NativeVector> for NativeVector {
	fn lse(&self, x : &NativeVector) -> bool {
		if self.vector.len() > x.vector.len() {
			return false;
		}
		for i in 0..x.vector.len() {
			if self.vector[i] > x.vector[i] {
				return false;
			}
			if self.vector[i] < x.vector[i] {
				return true;
			}
		}
		return true;
	}
}