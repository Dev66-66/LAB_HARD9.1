import urllib.request
import json
import argparse
import sys

def compute(start, end, workers=None, host="http://localhost:8080"):
    """
    Calls the Go backend HTTP API to perform heavy computation.
    """
    params = f"start={start}&end={end}"
    if workers:
        params += f"&workers={workers}"
    
    url = f"{host}/compute?{params}"
    
    try:
        with urllib.request.urlopen(url) as response:
            if response.status == 200:
                data = json.loads(response.read().decode('utf-8'))
                return data
            else:
                print(f"Error: Received status code {response.status}")
                return None
    except Exception as e:
        print(f"Error connecting to Go backend: {e}")
        return None

def main():
    parser = argparse.ArgumentParser(description="Python client for Go compute backend")
    parser.add_argument("--start", type=int, default=1, help="Start of range")
    parser.add_argument("--end", type=int, default=10, help="End of range")
    parser.add_argument("--workers", type=int, help="Number of workers (goroutines)")
    parser.add_argument("--host", type=str, default="http://localhost:8080", help="Go backend URL")
    
    args = parser.parse_args()
    
    print(f"Requesting computation for range [{args.start}, {args.end}]...")
    result = compute(args.start, args.end, args.workers, args.host)
    
    if result:
        print("--- Computation Results ---")
        print(f"Total Sum: {result.get('result')}")
        print(f"Range: [{result.get('start')}, {result.get('end')}]")
        print(f"Workers: {result.get('workers')}")
    else:
        print("Failed to get result from backend.")
        sys.exit(1)

if __name__ == "__main__":
    main()
