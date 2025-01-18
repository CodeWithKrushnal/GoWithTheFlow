# Error Handling Program for Invalid slice Indices

In the given code, the accessSlice function accepts a slice and index.

If the value is present in the slice at that index, the program should print the following.

`item <index>, value <value present at the index>`

But if the index does not hold any value,
it will lead to an index out of range panic in our program.


So in order to safeguard our program from panicking, add a condition to handle the condition if

`index > lengthOfSlice - 1`

and return an error from the accessSlice function if the above condition is met.
The error message should be the following :

`length of the slice should be more than index`

Complete the given program to return an error from the accessSlice function and handle the returned error inside the main function to print the message.

## Example Test Case :
- Input: 3
- Output:
  - item 3
  - value 6