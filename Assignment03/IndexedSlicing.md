# Index Based Reslicing 

Complete the program to return 3 slices of a given array, based on the following conditions.
Given Array : ["qwe", "wer", "ert", "rty", "tyu", "yui", "uio", "iop"]

Hint: Hard-code the given array into your program.

- Input: Two space-separated integers representing index1 and index2.
- Output: Output will contain 3 slices

1. slice containing all the elements from the start of the array to Index1
2. slice containing all the elements from index1 to index2
3. slice containing all the elements from index2 to the end of the array

- Conditions to Handle:
If Either of the input indexes is out of range of the given array, print "Incorrect Indexes" and exit the program

## Example Test Case 1:
- Input: 2 4
- Output:<br>
[qwe wer ert]<br>
[ert rty tyu]<br>
[tyu yui uio iop]

## Example Test Case 2:
- Input: 2 8
- Output: Incorrect Indexes
