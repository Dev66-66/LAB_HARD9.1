import unittest
from main import compute
from http.server import HTTPServer, BaseHTTPRequestHandler
import threading
import json

class MockGoBackend(BaseHTTPRequestHandler):
    def do_GET(self):
        self.send_response(200)
        self.send_header('Content-Type', 'application/json')
        self.end_headers()
        # Mock logic similar to Go backend
        response = {
            "result": 385,
            "start": 1,
            "end": 10,
            "workers": 4
        }
        self.wfile.write(json.dumps(response).encode('utf-8'))

    def log_message(self, format, *args):
        return # Disable logging

def run_mock_server():
    server = HTTPServer(('localhost', 8888), MockGoBackend)
    server.serve_forever()

class TestPythonClient(unittest.TestCase):
    @classmethod
    def setUpClass(cls):
        cls.server_thread = threading.Thread(target=run_mock_server, daemon=True)
        cls.server_thread.start()

    def test_compute_call(self):
        result = compute(1, 10, workers=4, host="http://localhost:8888")
        self.assertIsNotNone(result)
        self.assertEqual(result['result'], 385)
        self.assertEqual(result['start'], 1)
        self.assertEqual(result['end'], 10)

if __name__ == '__main__':
    unittest.main()
