import pytest

from p001.main import findSum

class TestFindSum(object):
    """
    This function is to test main.py's functions
    """
    def test_findSum(self):
        assert findSum([3, 5], 10) == 23

    def test_findSum(self):
      assert findSum([3, 5], 1000) == 266333