mod tests{
	use big_hex::applied_algebra::field::large_number::*;
	use itertools::izip;
	use lazy_static::lazy_static;

	const TEST_CASES_NUM : usize = 5;

	lazy_static!{
		static ref TEST_X: [NativeVector; TEST_CASES_NUM] = [
			vec![1, 0, 1, 1], 			//1101
			vec![1, 1, 0, 1], 			//1011
			vec![1, 1, 0, 1, 1], 		//011011
			vec![1, 0, 1, 0, 1, 0, 1], 	//01010101
			vec![1, 1, 0, 1, 1, 0, 1]  	//01011011
		];
		
		static ref TEST_Y: [NativeVector; TEST_CASES_NUM] = [
			vec![1, 0, 0, 1], 			//1001
			vec![0, 1, 1, 0],			//0110
			vec![1, 1, 0, 1, 1], 		//011011
			vec![1, 1, 1, 1, 0, 0, 0], 	//00001111
			vec![1, 1, 0, 1, 1, 0, 1] 	//01011011
		];
	}

	#[test]
	fn test_add() {
		let mut test_vector: [NativeVector; TEST_CASES_NUM] = [
			vec![0; 5],
			vec![0; 5],
			vec![0; 6],
			vec![0; 8],
			vec![0; 8]
		];

		let expected: [NativeVector; TEST_CASES_NUM] = [
			vec![0, 1, 1, 0, 1], 		  	  //1101	+ 1001	  = 10110
			vec![1, 0, 0, 0, 1], 		  	  //1011	+ 110	  = 10001
			vec![0, 1, 1, 0, 1, 1], 	  	  //11011	+ 11011	  = 0110110
			vec![0, 0, 1, 0, 0, 1, 1, 0],	  //1010101 + 1111	  = 01100100
			vec![0, 1, 1, 0, 1, 1, 0, 1]  	  //1011011 + 1011011 = 10110110
		];
	
		for (test, x, y, ex) in izip!(&mut test_vector, TEST_X.iter(), TEST_Y.iter(), &expected)  {
			test.add(x, y);
			println!("#{:?} + {:?}", x, y);
			assert_eq!(test, ex);
		}
	}

	#[test]
	fn test_sub() {
		let mut test_vector: [NativeVector; TEST_CASES_NUM] = [
			vec![0; 4],
			vec![0; 4],
			vec![0; 5],
			vec![0; 7],
			vec![0; 7]
		];

		let expected: [NativeVector; TEST_CASES_NUM] = [
			vec![0, 0, 1, 0], 			//1101	  - 1001	= 00100
			vec![1, 0, 1, 0], 			//1011	  - 110		= 00101
			vec![0, 0, 0, 0, 0],		//11011	  - 11011	= 000000
			vec![0, 1, 1, 0, 0, 0, 1],	//1010101 - 1111    = 01000110
			vec![0, 0, 0, 0, 0, 0, 0]  	//1011011 - 1011011 = 00000000
		];
		
		for (test, x, y, ex) in izip!(&mut test_vector, TEST_X.iter(), TEST_Y.iter(), &expected)  {
			test.sub(x, y);
			println!("#{:?} - {:?}", x, y);
			assert_eq!(test, ex);
		}
	}
}