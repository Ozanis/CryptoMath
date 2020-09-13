pub fn order(mut x : usize) -> usize {
	let mut i : usize = 32;
	while i > 0 {
		x |= x >> i;
		i >>= 1;
	}
	return x - (x >> 1);
}

//IsBinRadix returns 1 in case if given number is radix of 2 and 0 otherwise in O(1)
pub fn is_radix(x : usize) -> usize {
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
pub fn reverse(mut x : usize, mut len : usize) -> usize {
	let mut y : usize = 0;
	for i in 0..len {
		y <<= 1;
		y |= x & 1;
		x >>= 1;
	}
	return y;
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

//IsBinRadix returns 1 in case if given number is radix of 2 and 0 otherwise in O(1)
pub fn is_radix(x : u16) -> u16 {
	return !(x & (x - 1));
}

//log_bin finds log with  base 2 of 16 bit number with O(log(N)) complexity
pub	fn log(mut x : u16) -> u16 {
	const mask : [u16; 4] = [0x2, 0xC, 0xF0, 0xFF00];
	const shift : [u16; 4] = [1, 2, 4, 8];
	let mut i : usize = 4;
	let mut log2 : u16 = 0;
	for i in 4..=0 {
  		if x & mask[i] > 0 {
    		x >>= shift[i];
			log2 |= shift[i];
		  } 
		  i -= 1;
		}
	return log2
}

//Invert implements bit inversion with O(log(N)) required by FFT
pub fn reverse(mut x : u16) -> u16 {
	x = ((x & 0xaaaa) >> 1) | ((x & 0x5555) << 1);
	x = ((x & 0xcccc) >> 2) | ((x & 0x3333) << 2);
	x = ((x & 0xf0f0) >> 4) | ((x & 0x0f0f) << 4);
	x = ((x & 0xff00) >> 8) | ((x & 0x00ff) << 8);
	(x >> 16) | (x << 16)
}

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

pub fn inverse_bit(mut x : u16) -> u16 {
	let mut mask : u16 = 1 << log_bin(x);
	mask |= mask - 1;
	return x ^ mask
}*/
