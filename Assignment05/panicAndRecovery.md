# Panic and Recovery Mechanism Program

In the given code, the accessSlice function accepts a slice and index.

If the value is present in the slice at that index, the program should print the following.

`item <index>, value <value at that index in slice>`

But if the index does not hold any value, it will lead to an index out of range panic in our program.

Complete the given code to recover from panic inside the accessSlice function, when index out of range panic occurs.
Also, Print the following after handling panic

`internal error: <recovered panic value`>


## Example Test Case :
- Input: 3
- Output:
  - item 3
  - value 6