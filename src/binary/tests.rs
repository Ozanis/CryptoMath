#[cfg(test)]
mod tests {
	
	use crate::binary::binary::*;

	static TEST_X : [usize; 16] = [
		0, 1, 2, 3,
		4, 5, 6, 7,
		8, 9, 10, 11,
		121, 128, 200, 255
	];

	#[test]
	fn test_order() {
		let expected : [usize; 16] = [
			0, 1, 2, 2,
			4, 4, 4, 4,
			8, 8, 8, 8,
			64, 128, 128, 128
		];
		let mut got : usize;
		for i in 0..expected.len() {
			got = log_bin(TEST_X[i]);
			dbg!(i, TEST_X[i], got);
			assert_eq!(got, expected[i]);
		}
	}
	
	#[test]
	fn test_log() {
		let expected : [usize; 16] = [
			0, 0, 1, 1,
			2, 2, 2, 2,
			3, 3, 3, 3,
			6, 7, 7, 7
		];
		let mut got : usize;
		for i in 0..expected.len() {
			got = log(TEST_X[i]);
			dbg!(i, TEST_X[i], got);
			assert_eq!(got, expected[i]);
		}
	}

	#[test]
	fn test_reverse() {
		let expected : [usize; 16] = [
			0, 1, 1, 3,
			1, 5, 3, 7,
			1, 9, 5, 13,
			79, 1, 19, 255
		];
		let len : [usize; 16] = [
			1, 1, 2, 2,
			3, 3, 3, 3,
			4, 4, 4, 4,
			7, 8, 8, 8,
		];
		let mut got : usize;
		for i in 0..expected.len() {
			got = reverse(TEST_X[i], len[i]);
			dbg!(i, TEST_X[i], got);
			assert_eq!(got, expected[i]);
		}		
	}
}