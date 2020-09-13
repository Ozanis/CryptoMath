mod tests{
	
	use lazy_static::lazy_static;
	use crate::vector::vector::*;

	lazy_static!{
		static ref TEST_X: [NativeVector; 4] = [
			NativeVector{vector: vec![1, 1, 0, 1, 0, 0, 0, 0]},
			NativeVector{vector: vec![1, 1, 0, 1, 1, 0, 0, 0]},
			NativeVector{vector: vec![1, 0, 1, 0, 1, 0, 1, 0]},
			NativeVector{vector: vec![1, 1, 0, 1, 1, 0, 1, 0]}
		];

		static ref TEST_Y: [NativeVector; 4] = [
			NativeVector{vector: vec![0, 1, 1, 0, 0, 0, 0, 0]},
			NativeVector{vector: vec![1, 1, 0, 1, 1, 0, 0, 0]},
			NativeVector{vector: vec![1, 1, 1, 1, 0, 0, 0, 0]},
			NativeVector{vector: vec![1, 1, 0, 1, 1, 0, 1, 0]}
			];
		}
	
	#[test]
	fn test_add() {
		let expected: [NativeVector; 4] = [
			NativeVector{vector: vec![1, 0, 0, 0, 1, 0, 0, 0]},
			NativeVector{vector: vec![0, 1, 1, 0, 1, 1, 0, 0]},
			NativeVector{vector: vec![0, 0, 1, 0, 0, 1, 1, 0]},
			NativeVector{vector: vec![0, 1, 1, 0, 1, 1, 0, 1]}
		];
		let mut got : NativeVector;
		for i in 0..4 {
			got = NativeVector{vector: TEST_X[i].vector.clone()};
			assert_eq!(got.add(&TEST_Y[i]).vector, expected[i].vector);
			dbg!(i, &TEST_X[i].vector, &TEST_Y[i].vector);
		}
	}

	#[test]
	fn test_sub() {
		let expected: [NativeVector; 4] = [
			NativeVector{vector: vec![1, 0, 1, 0, 0, 0, 0, 0]},
			NativeVector{vector: vec![0, 0, 0, 0, 0, 0, 0, 0]},
			NativeVector{vector: vec![0, 1, 1, 0, 1, 0, 1, 0]},
			NativeVector{vector: vec![0, 0, 0, 0, 0, 0, 0, 0]}
		];
		let mut got : NativeVector;
		for i in 0..4 {
			got = NativeVector{vector: TEST_X[i].vector.clone()};
			assert_eq!(got.sub(&TEST_Y[i]).vector, expected[i].vector);
			dbg!(i, &TEST_X[i].vector, &TEST_Y[i].vector);
		}
	}
/*
	#[test]
	fn test_rem() {
		let expected: [NativeVector; 4] = [
			NativeVector{vector: vec![1, 0, 0, 1]},
			NativeVector{vector: vec![0, 0, 0, 0, 0, 0]},
			NativeVector{vector: vec![0, 1, 0, 1, 1, 0, 1, 0]},
			NativeVector{vector: vec![0, 1, 1, 0, 1, 1]}
		];
		let mut got : NativeVector;
		for i in 0..4 {
			got = NativeVector{vector: TEST_X[i].vector.clone()};
			assert_eq!(got.rem(&TEST_Y[i]).vector, expected[i].vector);
			dbg!(i, &TEST_X[i].vector, &TEST_Y[i].vector);
		}
	}*/
/*
	#[test]
	fn test_div() {
		let expected: [NativeVector; 4] = [

		];
		let mut got : NativeVector;
		for i in 0..4 {
			got = TEST_X[i].div(TEST_Y[i]);
			println!("{} : {:?} / {:?} is {:?}", i, TEST_X[i].vector, TEST_Y[i].vector, got.vector);
			assert_eq!(got, expected[i]);
			}
		}*/
	
	#[test]
	fn test_eq() {
		let expected: [bool; 4] = [
			false, true, false, true
			];
		let mut got : bool;
		for i in 0..4 {
			got = TEST_X[i].eq(&TEST_Y[i]);
			assert_eq!(got, expected[i]);
			dbg!(i, &TEST_X[i].vector, &TEST_Y[i].vector, got);
		}
	}

	#[test]
	fn test_gt() {
		let expected: [bool; 4] = [
			true, false, false, false
		];
		let mut got : bool;
		for i in 0..4 {
			got = TEST_X[i].gt(&TEST_Y[i]);
			assert_eq!(got, expected[i]);
			dbg!(i, &TEST_X[i].vector, &TEST_Y[i].vector, got);
			}
		}
	

	#[test]
	fn test_gte() {
		let expected: [bool; 4] = [
			true, true, false, true
		];
		let mut got : bool;
		for i in 0..4 {
			got = TEST_X[i].gte(&TEST_Y[i]);
			println!("{} : {:?} >= {:?} is {}", i, TEST_X[i].vector, TEST_Y[i].vector, got);
			assert_eq!(got, expected[i]);
		}
	}

	#[test]
	fn test_ls() {
		let expected: [bool; 4] = [
			false, false, true, false
		];
		let mut got : bool;
		for i in 0..4 {
			got = TEST_X[i].ls(&TEST_Y[i]);
			println!("{} : {:?} < {:?} is {}", i, TEST_X[i].vector, TEST_Y[i].vector, got);
			assert_eq!(got, expected[i]);
		}
	}

	#[test]
	fn test_lse() {
		let expected: [bool; 4] = [
			false, true, true, true
		];
		let mut got : bool;
		for i in 0..4 {
			got = TEST_X[i].lse(&TEST_Y[i]);
			println!("{} : {:?} <= {:?} is {}", i, TEST_X[i].vector, TEST_Y[i].vector, got);
			assert_eq!(got, expected[i]);
		}
	}

/*
    #[test]
    fn test_fast_walsh_hadamard() {
	    let expected : [NativeVector; 4] = [
			NativeVector{vector: vec![1, 0, 1, 0]},
    		NativeVector{vector: vec![1, 1, 1, 1, 1, 1, 1, 1]},
			NativeVector{vector: vec![4, 2, 0, -2, 0, 2, 0, 2]},
			NativeVector{vector: vec![4, 2, 0, -2, 0, 2, 0, 2]}
		];
		let mut got : NativeVector;
		for i in 0..4 {
			got = test_x[i].fast_walsh_hadamard();
			println!("{} : FWHT of {:?} is {:?}", i, test_x[i].vector, got.vector);
			assert_eq!(got, expected[i]);
		}
    }

    #[test]
    fn test_radix() {

    }	

    #[test]
    fn test_convolution() {
		let expected : [NativeVector; 4] = [
			NativeVector{vector: vec![1, 0, 1, 0]},
    		NativeVector{vector: vec![1, 1, 1, 1, 1, 1, 1, 1]},
			NativeVector{vector: vec![4, 2, 0, -2, 0, 2, 0, 2]},
			NativeVector{vector: vec![4, 2, 0, -2, 0, 2, 0, 2]}
		];
		let mut got : NativeVector;
		for i in 0..4 {
			got = test_x[i].convolution(test_y[i]);
			println!("{} : {:?} ^ {:?} is {:?}", i, test_x[i].vector, test_y[i].vector[i], got.vector);
			assert_eq!(got, expected[i]);
		}
	}*/


	/*
	#[test]
	fn test_mul() {
		let expected: [NativeVector; 4] = [
			NativeVector{vector: vec![1, 0, 0, 1, 1, 1, 0]},
			NativeVector{vector: vec![1, 0, 1, 1, 0, 1, 1, 0, 0, 1, 0, 0]},
			NativeVector{vector: vec![1, 0, 0, 1, 1, 1, 1, 1, 0, 1, 1, 0, 0, 0, 0, 0]},
			NativeVector{vector: vec![1, 0, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 0]},
		];
		let mut got : NativeVector;
		for i in 0..4 {
			got = TEST_X[i].mul(&TEST_Y[i]);
			println!("{} : {:?} * {:?} is {:?}", i, TEST_X[i].vector, TEST_Y[i].vector, got.vector);
			assert_eq!(got, expected[i]);
			}
		}

	#[test]
	fn test_sqr() {
		let expected: [bool; 4] = [
			NativeVector{vector: vec![1, 0, 1, 0, 1, 0, 0, 1]},
			NativeVector{vector: vec![1, 0, 1, 1, 0, 1, 1, 0, 0, 1, 0, 0]},
			NativeVector{vector: vec![1, 1, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0]},
			NativeVector{vector: vec![1, 0, 1, 1, 1, 0, 0, 1, 1, 0, 1, 0, 0, 1]}
		];
		let mut got : NativeVector;
		for i in 0..4 {
			got = TEST_X[i].sqr();
			println!("{} : {:?} ^2  is {:?}", i, TEST_X[i].vector, got.vector);
			assert_eq!(got, expected[i]);
		}
	}*/
/*
    #[test]
    fn test_exp() {
		let expected: [NativeVector; 4] = [
				NativeVector{vector: vec![1, 0, 0, 1, 0, 0, 1, 1, 0, 1, 0, 0, 1, 1, 0, 1, 0, 1, 1, 1, 0, 0, 1]},
				NativeVector{vector: vec![1, 0, 0, 1, 0, 0, 1, 1, 0, 1, 0, 0, 1, 1, 0, 1, 0, 1, 1, 1, 0, 0, 1]},
				NativeVector{vector: vec![1, 0, 0, 1, 0, 0, 1, 1, 0, 1, 0, 0, 1, 1, 0, 1, 0, 1, 1, 1, 0, 0, 1]},
				NativeVector{vector: vec![1, 0, 0, 1, 0, 0, 1, 1, 0, 1, 0, 0, 1, 1, 0, 1, 0, 1, 1, 1, 0, 0, 1]},
			];
		let test_mod: [NativeVector; 4] = [
			NativeVector{vector: vec![1, 0, 0, 1, 0, 0, 1, 1, 0, 1, 0, 0, 1, 1, 0, 1, 0, 1, 1, 1, 0, 0, 1]},
			NativeVector{vector: vec![1, 0, 0, 1, 0, 0, 1, 1, 0, 1, 0, 0, 1, 1, 0, 1, 0, 1, 1, 1, 0, 0, 1]},
			NativeVector{vector: vec![1, 0, 0, 1, 0, 0, 1, 1, 0, 1, 0, 0, 1, 1, 0, 1, 0, 1, 1, 1, 0, 0, 1]},
			NativeVector{vector: vec![1, 0, 0, 1, 0, 0, 1, 1, 0, 1, 0, 0, 1, 1, 0, 1, 0, 1, 1, 1, 0, 0, 1]},
		];
		let mut got : NativeVector;
		for i in 0..4 {
			got = TEST_X[i].exp(TEST_Y[i], test_mod);
			println!("{} : {:?} ^ {:?}  is {:?}", i, TEST_X[i].vector, TEST_Y[i].vector, got.vector);
			assert_eq!(got.vector, expected[i].vector);
		}         
    }
*/
}