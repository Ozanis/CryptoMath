extern crate num;

use num::Unsigned;

// 
pub fn log_bin(mut x : usize) -> usize {
	let mut i : usize = 32;
	while i > 0 {
		x |= x >> i;
		i >>= 1;
	}
	return x - (x >> 1);
}

//is_bin_radix returns 1 if given number divisible by 2 without reminder and returns 0 otherwise in O(1)
pub fn is_bin_radix <T> (x : T) -> T 
where T: PrimInt + Unsigned {
	return !(x & (x - 1));
}

//log_bin finds log with  base 2 of 64 bit number with O(log(N)) complexity
pub	fn log(mut x : usize) -> usize {
	if x == 0 {
		return 0;
	}
	let mask : [usize; 6] = [ 0xFFFFFFFF00000000, 0xFFFF0000, 0xFF00, 0xF0, 0xC, 0x2 ];
	let mut y : usize = (x & (x - 1)) & 1;
	let mut j : usize = 32;
	for i in 0..6 {
		if x & mask[i] != 0 {
			y += j;
			x >>= j;
		};
	  j >>= 1;
	}
	return y
}

//Invert implements bit inversion with O(log(N)) required by FFT
pub fn reverse(mut x : usize, len : usize) -> usize {
	let mut size : usize = len;
	let mut mask : usize = !0;         
	while size > 0 {
		size >>= 1;
  		mask ^= mask << size;
  		x = ((x >> size) & mask) | ((x << size) & !mask);
	}
	return x;
}

/*
//RoudDownBin returns 2 base number rounded down in O(log(N))
pub fn order(mut x : u16) -> u16 {
	x |= x >> 1;
	x |= x >> 2;
	x |= x >> 4;
	x |= x >> 8;
	return x - (x >> 1);
}

pub fn inverse_bit(mut x : u16) -> u16 {
	let mut mask : u16 = 1 << log_bin(x);
	mask |= mask - 1;
	return x ^ mask
}*/
