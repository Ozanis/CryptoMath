pub struct BigHex {
    vector: Vec<u64>
}

trait Encoding {
    fn new(vector: Vec<u64>) -> Self;
    fn set_hex_string(&mut self,  hex_str: &str);
    fn get_hex_string(&self) -> String;
}


impl Encoding for BigHex {
    fn new(vector: Vec<u64>) -> Self {
        BigHex { vector }
    }

    // Метод для встановлення числа за його рядковим представленням у шістнадцятковій системі
    fn set_hex_string(&mut self, hex_string: &str) {
        // Конвертація рядка з шістнадцятковим числом у вектор беззнакових цілих чисел
        self.vector = hex_string
            .chars()
            .collect::<Vec<char>>()
            .chunks(16) // Розділення рядка на частини по 16 символів (64 біти)
            .map(|chunk| {
                chunk
                    .iter()
                    .collect::<String>()
                    .chars()
                    .collect::<String>()
                    .parse::<u64>()
                    .unwrap()
            })
            .collect::<Vec<u64>>();
    }        

    // Метод для отримання рядкового представлення числа у шістнадцятковій системі
    fn get_hex_string(&self) -> String {
        let hex_digits: Vec<String> = self
            .vector
            .iter()
            .map(|&chunk| format!("{:016X}", chunk).chars().collect::<String>())
            .collect();

        hex_digits.concat().chars().collect::<String>()
    }
}