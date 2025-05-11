import time

def fib(n):
    if n < 2:
        return n
    return fib(n-1) + fib(n-2)

def main():
    time_start = time.time()
    for i in range(45):
        print(fib(i))
        print(time.time() - time_start)
    time_end = time.time()
    elapsed_time = time_end - time_start
    print(f"{elapsed_time:.9f}")

main()