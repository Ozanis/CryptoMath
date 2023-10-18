mod tests{
	use big_hex::BigHex;

	#[test]
	fn test_encoding() {
		let mut test_string = "1A2B3C4D5E6F";
		let mut big_hex = BigHex::new(Vec::new());
    	big_hex.set_hex_string(&test_string);
    	println!("Large Number: {}", large_num.get_hex_string());
		assert_eq!(&test_string, large_num.get_hex_string());
	}

}