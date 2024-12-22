# TODO: blah bla bhla
def fizzbuzz(n):
    """Random ChatGPT generate fizzbuzz code 
    Used just for test cases.
    
    Prints the numbers from 1 to n, 
    replacing multiples of 3 with "Fizz", 
    multiples of 5 with "Buzz", 
    and multiples of both with "FizzBuzz"."""

    for i in range(1, n + 1):
        if i % 3 == 0 and i % 5 == 0:  # Check for divisibility by both 3 and 5
            print("FizzBuzz")
        elif i % 3 == 0:  # Check for divisibility by 3
            print("Fizz")
        elif i % 5 == 0:  # Check for divisibility by 5
            print("Buzz")
        else:
            print(i)  # Print the number itself if not divisible by 3 or 5

# Example usage:
fizzbuzz(15) 