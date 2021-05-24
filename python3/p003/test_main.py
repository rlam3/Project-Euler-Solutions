from p003.main import find_largest_prime_factor


class TestFindSum(object):
    """
    This function is to test main.py's functions
    """

    def find_largest_prime_factor_13195(self):
        assert find_largest_prime_factor(13195) == 29

    def find_largest_prime_factor_600851475143(self):
        assert find_largest_prime_factor(600851475143) == 266333
