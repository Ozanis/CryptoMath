mod tests{
	
	use lazy_static::lazy_static;
	use crate::field::vector::native::*;
	use crate::operator::group::*;
	use crate::operator::logical::*;

	const TEST_CASES_NUM : usize = 4;

	lazy_static!{
		static ref TEST_X: [NativeVector; TEST_CASES_NUM] = [
			vec![1, 1, 0, 1, 0], 			//01011
			vec![1, 1, 0, 1, 1, 0], 		//011011
			vec![1, 0, 1, 0, 1, 0, 1, 0], 	//01010101
			vec![1, 1, 0, 1, 1, 0, 1, 0]  	//01011011
		];
		
		static ref TEST_Y: [NativeVector; TEST_CASES_NUM] = [
			vec![0, 1, 1, 0, 0],			//00110
			vec![1, 1, 0, 1, 1, 0], 		//011011
			vec![1, 1, 1, 1, 0, 0, 0, 0], 	//00001111
			vec![1, 1, 0, 1, 1, 0, 1, 0] 	//01011011
		];
	}


	#[test]
	fn test_add() {
		let mut test_vector: [NativeVector; TEST_CASES_NUM] = [
			vec![0; 5],
			vec![0; 6],
			vec![0; 8],
			vec![0; 8]
		];

		let expected: [NativeVector; TEST_CASES_NUM] = [
			vec![1, 0, 0, 0, 1], 		  //1011+110=10001
			vec![0, 1, 1, 0, 1, 1], 	  //11011+11011=0110110
			vec![0, 0, 1, 0, 0, 1, 1, 0],	  //1010101+1111=01100100
			vec![0, 1, 1, 0, 1, 1, 0, 1]  //1011011+1011011=10110110
		];
	
		for i in 0..TEST_CASES_NUM {
			dbg!(&test_vector[i]);
			assert_eq!(test_vector[i].add(&TEST_X[i], &TEST_Y[i]), &expected[i]);
		}
	}

	#[test]
	fn test_sub() {
		let mut test_vector: [NativeVector; TEST_CASES_NUM] = [
			vec![0; 5],
			vec![0; 6],
			vec![0; 8],
			vec![0; 8]
		];

		let expected: [NativeVector; TEST_CASES_NUM] = [
			vec![1, 0, 1, 0, 0], 			//1011-110=00101
			vec![0, 0, 0, 0, 0, 0],			//11011-11011=000000
			vec![0, 1, 1, 0, 0, 0, 1, 0],	//1010101-1111=01000110
			vec![0, 0, 0, 0, 0, 0, 0, 0]  	//1011011-1011011=00000000
		];
		
		for i in 0..TEST_CASES_NUM {
			dbg!(&test_vector[i]);
			assert_eq!(test_vector[i].sub(&TEST_X[i], &TEST_Y[i]), &expected[i]);
		}
	}

	#[test]
	fn test_mul() {
		let mut test_vector: [NativeVector; TEST_CASES_NUM] = [
			vec![0; 8],
			vec![0; 10],
			vec![0; 12],
			vec![0; 14]
		];

		let expected: [NativeVector; TEST_CASES_NUM] = [
			vec![0, 1, 0, 0, 0, 0, 0, 1], 					//1011*110=1000010
			vec![1, 0, 0, 1, 1, 0, 1, 1, 0, 1], 			//11011*11011=1011011001
			vec![1, 1, 0, 1, 1, 1, 1, 1, 0, 0, 1, 0], 		//1010101*1111=1110000111001
			vec![1, 0, 0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 1]  //1011011*1011011=10000001011001
		];
		
	
		for i in 0..TEST_CASES_NUM {
			assert_eq!(test_vector[i].mul(&TEST_X[i], &TEST_Y[i]), &expected[i]);
			dbg!(&test_vector[i]);
			}
		}

/*
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
	}

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
/*
	#[test]
	fn test_rem() {
		const TEST_VECTOR_SIZE : usize = 8;
		let test_vector: [NativeVector; TEST_CASES_NUM] = [
			NativeVector{vector: vec![0; TEST_VECTOR_SIZE ]},
			NativeVector{vector: vec![0; TEST_VECTOR_SIZE ]},
			NativeVector{vector: vec![0; TEST_VECTOR_SIZE ]},
			NativeVector{vector: vec![0; TEST_VECTOR_SIZE ]}
		];
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
	
	#[test]
	fn test_eq() {
		let expected: [bool; 4] = [
			false, 	//1011    == 110     = false
			true, 	//11011   == 11011   = true
			false,  //1010101 == 1111    = false
			true 	//1011011 == 1011011 = true
		];

		let mut got : bool;
		for i in 0..TEST_CASES_NUM {
			got = TEST_X[i].equals(&TEST_Y[i]);
			dbg!(&TEST_X[i], &TEST_Y[i]);
			assert_eq!(got, expected[i]);
		}
	}

	#[test]
	fn test_gt() {
		let expected: [bool; 4] = [
			true,   //1011 > 110 = true
			false,  //11011 > 11011 = false
			true,   //1010101 > 1111 = true
			false   //1011011 > 1011011 = false
		];

		let mut got : bool;
		for i in 0..TEST_CASES_NUM {
			got = TEST_X[i].greater(&TEST_Y[i]);
			dbg!(&TEST_X[i], &TEST_Y[i]);
			assert_eq!(got, expected[i]);
			}
		}
	

	#[test]
	fn test_gte() {
		let expected: [bool; 4] = [
			true,  //1011 >= 110 = true
			true,  //11011 >= 11011 = true
			true,  //1010101 >= 1111 = true
			true   //1011011 >= 1011011 = true
		];

		let mut got : bool;
		for i in 0..TEST_CASES_NUM {
			got = TEST_X[i].greater_equals(&TEST_Y[i]);
			dbg!(&TEST_X[i], &TEST_Y[i]);
			assert_eq!(got, expected[i]);
		}
	}

	#[test]
	fn test_ls() {
		let expected: [bool; 4] = [
			false,	//1011 <= 110 = false
			false,	//11011 <= 11011 = false
			false,	//1010101 <= 1111 = false
			false	//1011011 <= 1011011 = false
		];

		let mut got : bool;
		for i in 0..TEST_CASES_NUM {
			got = TEST_X[i].lesser(&TEST_Y[i]);
			dbg!(&TEST_X[i], &TEST_Y[i]);
			assert_eq!(got, expected[i]);
		}
	}

	#[test]
	fn test_lse() {
		let expected: [bool; 4] = [
			false,  //1011 <= 110 = false
			true,  	//11011 <= 11011 = true
			false,	//1010101 <= 1111 = false
			true	//1011011 <= 1011011 = true
		];

		let mut got : bool;
		for i in 0..TEST_CASES_NUM {
			got = TEST_X[i].lesser_equals(&TEST_Y[i]);
			dbg!(&TEST_X[i], &TEST_Y[i]);
			assert_eq!(got, expected[i]);
		}
	}

}