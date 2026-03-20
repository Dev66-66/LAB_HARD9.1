import unittest
import os
import subprocess
from caller import call_go_tool

class TestPythonCaller(unittest.TestCase):
    def test_call_go_tool(self):
        # Находим путь к бинарнику
        binary = os.path.abspath("../go_tool/go_tool.exe")
        
        # Если бинарник еще не скомпилирован (в CI/CD), пропускаем
        if not os.path.exists(binary):
            self.skipTest("Go binary not found. Build it first.")
        
        # Вызов
        res = call_go_tool(5, binary)
        self.assertEqual(res, 25)

if __name__ == "__main__":
    unittest.main()
