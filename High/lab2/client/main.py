import urllib.request
import json
import time
import sys

# Пытаемся импортировать Rust-библиотеку. 
# Если она еще не собрана, используем заглушку для тестирования.
try:
    import crypto_lib
except ImportError:
    print("WARNING: crypto_lib not found. Make sure to build it with 'maturin develop' or 'cargo build'.")
    print("Using a placeholder encryption for demonstration.")
    class CryptoMock:
        @staticmethod
        def encrypt(data, key):
            return f"ENCRYPTED_MOCK({data}, key={key})"
    crypto_lib = CryptoMock()

def call_api(url, method="GET", payload=None):
    """Отправка HTTP-запроса к Go-оркестратору."""
    req = urllib.request.Request(url, method=method)
    if payload:
        req.add_header('Content-Type', 'application/json')
        body = json.dumps(payload).encode('utf-8')
    else:
        body = None
    
    try:
        with urllib.request.urlopen(req, data=body) as resp:
            if resp.status in (200, 204):
                return json.loads(resp.read().decode('utf-8')) if resp.status == 200 else True
    except Exception as e:
        print(f"Error calling {url}: {e}")
    return None

def orchestrate(host="http://localhost:8081"):
    # 1. Создание задачи (имитируем поступление данных)
    print("1. Creating job in Go orchestrator...")
    job = call_api(f"{host}/job/create", method="POST", payload={"data": "Secret Lab Message"})
    if not job:
        print("Failed to create job.")
        return
    
    job_id = job['id']
    raw_data = job['data']
    print(f"Job #{job_id} created with data: '{raw_data}'")

    # 2. Шифрование в Rust
    print("2. Encrypting data with Rust library...")
    key = 42
    encrypted_data = crypto_lib.encrypt(raw_data, key)
    print(f"Encrypted data: '{encrypted_data}'")

    # 3. Обновление статуса в Go
    print("3. Sending results back to Go orchestrator...")
    success = call_api(f"{host}/job/update", method="PATCH", payload={"id": job_id, "encrypted": encrypted_data})
    
    if success:
        print(f"Job #{job_id} completed successfully.")
        
        # Проверяем финальный статус
        final_job = call_api(f"{host}/job/status?id={job_id}")
        print(f"Final Job State: {final_job}")
    else:
        print("Failed to update job.")

if __name__ == "__main__":
    orchestrate()
