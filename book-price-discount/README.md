# Coding Dojo
## Potter
This is a solution to the kata found at http://codingdojo.org/kata/Potter/.

### Explanation

For my first solution I obtained the distinct values in the array and acquired the discount for that number of books. I removed those distinct values from the array and continued to get distinct values until there are no books remaining in the array to price. With this solution I found 87% of the assertions in the test cases were passing.

The two assertions for edge cases were failing. This is due to a mechanism where the results are compared using initial groups of five or four books.

To complete these final assertions I wrote functions that acquire the specific number of values for a group. I can pass in a 4 or 5 to get groups of that corresponding size of distinct books. This allowed me to do two separate pricings for groups of 4 and 5. Comparing the results reveals the least expensive pricing solution. This solution brought 100% success in all the unit tests.