from p003.main import find_largest_prime_factor


class TestFindLargestPrime(object):
    """
    This function is to test main.py's functions
    """

    def test_find_largest_prime_factor_13195(self):
        assert find_largest_prime_factor(13195) == 29

    def xtest_find_largest_prime_factor_600851475143(self):
        assert find_largest_prime_factor(600851475143) == 266333
