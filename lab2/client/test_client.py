import unittest
from main import call_api
from http.server import HTTPServer, BaseHTTPRequestHandler
import threading
import json

class MockOrchestrator(BaseHTTPRequestHandler):
    def do_POST(self):
        self.send_response(200)
        self.send_header('Content-Type', 'application/json')
        self.end_headers()
        self.wfile.write(json.dumps({"id": 1, "data": "test"}).encode())

    def do_PATCH(self):
        self.send_response(204)
        self.end_headers()

    def log_message(self, format, *args):
        return

def run_server():
    server = HTTPServer(('localhost', 8889), MockOrchestrator)
    server.serve_forever()

class TestPythonOrchestrator(unittest.TestCase):
    @classmethod
    def setUpClass(cls):
        cls.server_thread = threading.Thread(target=run_server, daemon=True)
        cls.server_thread.start()

    def test_create_job(self):
        res = call_api("http://localhost:8889/job/create", method="POST", payload={"data": "test"})
        self.assertEqual(res['id'], 1)

    def test_update_job(self):
        res = call_api("http://localhost:8889/job/update", method="PATCH", payload={"id": 1, "encrypted": "xxx"})
        self.assertTrue(res)

if __name__ == "__main__":
    unittest.main()
