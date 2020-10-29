# Multiples of 3 and 5
# Problem 1
# If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.

# Find the sum of all the multiples of 3 or 5 below 1000.

# import time

# sum = 0;

# # print range(0,1001)

# def multiples(x):
# 	if x % 5 == 0 or x % 3 == 0:
# 		return True
# 	else:
# 		return False

# for i in range(0,1000):
# 	if(multiples(i)):
# 		sum = sum + i
# 	else:
# 		print "%i is not a multiple of 3 or 5" %i

# start = time.time()
# elapsed = time.time() - start
# print "The sum is::: %i ::: Time elapsed:: %f" %(sum , elapsed)
# print elapsed


def findSum(multiples, num):
    """
    docstring
    """
    sum = 0

    for i in range(0, num):
        for j in multiples:
            # Find if it is a multiple
            if (i % j == 0):
                # import pdb; pdb.set_trace()
                sum += i
    return sum