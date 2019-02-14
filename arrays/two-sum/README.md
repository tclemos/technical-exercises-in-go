# Exercise: Two sum

## Given an array of integers, return indices and values of the two numbers such that they add up to a specific target.

## Constraints
- The input must have exactly one solution.
- The collection can have repeated numbers.
- You can not use the same element twice.

## Example 1: 
> Given nums = [2, 7, 11, 15], target = 9,
> 
> Because nums[0] + nums[1] = 2 + 7 = 9,
> return 0, 1, 2, 7.

## Example 2:
> Given nums = [9, 4, 3, 7, 5, 2, 0], target = 8,
> 
> Because nums[2] + nums[4] = 3 + 5 = 8,
> return 2, 4, 3, 5.

# Usage
All the numbers should be provided via command line arguments.

- The first number will be consider as the __Target__, in the following exemple it is the number __9__.
- The rest of the numbers will be considered as the number collection to found the two numbers that matches the target when added, in the following example they are the numbers __2, 7, 11__ and __15__.

> C:\\my-go-workspace\\bin\\> __twosum.exe 9 2 7 11 15__

# Executing tests

To run all the tests from this project just execute the following command inside of the two-sum project folder:
> C:\\my-go-workspace\\src\\github.com\\tclemos\\technical-exercises-in-go\\arrays\\two-sum\\> __go test ./...__




