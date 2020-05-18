# Coding Dojo
## Potter
This is a solution to the [Potter kata](http://codingdojo.org/kata/Potter/){:target="_blank"} found at [Coding Dojo](http://codingdojo.org){:target="_blank"}.

### Explanation

For my first solution, I obtained the distinct values in the array and acquired the discount for that number of books. I aggregated the total price using the discount price obtained multiplied by the number of distinct books found. I removed those distinct values from the array and continued to get distinct values, discount price, and the total price until no books were remaining in the array to price. With this solution, I found that 87% of the assertions in the test cases were passing.

The two assertions for edge cases were failing. These failures are because it requires a mechanism that compares the results from starting with groups of 4 or 5 books.

To complete these final assertions, I wrote functions that acquire a specific number of values for a group. I can pass in a 4 or 5 to get groups of that corresponding size of distinct books. This function allowed me to do two separate pricings for groups of 4 and 5. Comparing the results reveals the least expensive pricing solution. This solution brought 100% success in all the unit tests.

I believe that there could be an optimization early on before doing two separate pricings. In the interest of solving the kata, I've forgone the investigation of that solution. I may revisit this at a later time.
