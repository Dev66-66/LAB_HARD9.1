import subprocess
import os

def call_go_tool(n, binary_path):
    """Вызов Go-бинарника через subprocess."""
    try:
        # В Windows добавляем .exe, если это не сделано
        if os.name == 'nt' and not binary_path.endswith('.exe'):
            binary_path += '.exe'

        result = subprocess.run([binary_path, str(n)], capture_output=True, text=True, check=True)
        return int(result.stdout.strip())
    except subprocess.CalledProcessError as e:
        print(f"Error calling binary: {e.stderr}")
        return None
    except Exception as e:
        print(f"General error: {e}")
        return None

def main():
    # Путь к скомпилированному бинарнику
    binary = os.path.abspath("../go_tool/go_tool.exe")
    
    number = 12
    print(f"Calling Go tool to square {number}...")
    res = call_go_tool(number, binary)
    
    if res is not None:
        print(f"Result from Go: {res}")
    else:
        print("Failed to get result from Go tool.")

if __name__ == "__main__":
    main()
