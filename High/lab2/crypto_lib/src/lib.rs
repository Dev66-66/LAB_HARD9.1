use pyo3::prelude::*;

/// XOR-шифрование строки с использованием ключа.
#[pyfunction]
fn encrypt(data: String, key: u8) -> PyResult<String> {
    let result: String = data
        .chars()
        .map(|c| {
            let b = c as u8;
            (b ^ key) as char
        })
        .collect();
    Ok(result)
}

/// Расшифрование (в случае XOR оно идентично шифрованию).
#[pyfunction]
fn decrypt(data: String, key: u8) -> PyResult<String> {
    encrypt(data, key)
}

/// Python-модуль на Rust.
#[pymodule]
fn crypto_lib(_py: Python, m: &PyModule) -> PyResult<()> {
    m.add_function(wrap_pyfunction!(encrypt, m)?)?;
    m.add_function(wrap_pyfunction!(decrypt, m)?)?;
    Ok(())
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_xor_logic() {
        let original = "Hello".to_string();
        let key = 42;
        let encrypted = encrypt(original.clone(), key).unwrap();
        let decrypted = decrypt(encrypted, key).unwrap();
        assert_eq!(original, decrypted);
    }
}
