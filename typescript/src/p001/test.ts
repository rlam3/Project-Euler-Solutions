// Multiples of 3 and 5
// Problem 1
// If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.

// Find the sum of all the multiples of 3 or 5 below 1000.
const findSum = require('./main');

describe('1: multiples of 3 and 5', () => {
  
  it('should find all multiples for 3 and 5 below 10', () => {
    expect(findSum([3,5],10)).toEqual(23)
  });

  it('should find all multiples for 3 and 5 below 1000', () => {
    expect(findSum([3,5],1000)).toEqual(266333)
  });

});