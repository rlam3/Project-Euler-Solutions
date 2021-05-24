# https://projecteuler.net/problem=3

# The prime factors of 13195 are 5, 7, 13 and 29.

# What is the largest prime factor of the number 600851475143 ?


def is_prime(number: int) -> bool:
    # if number is negative
    if number < 1 or number % 2 == 0:
        return False

    for x in range(3, number):
        if number % x == 0:
            return True


def find_largest_prime_factor(number: int) -> int:
    l_prime: list[int] = []

    for x in range(2, number):
        if is_prime(x):
            print(number, "is a prime")
            l_prime.append(x)

    return l_prime[-1]
