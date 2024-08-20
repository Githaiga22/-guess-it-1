# Guess-It-1
Overview

The guess-it-1 project is designed to challenge your mathematical and statistical skills by creating a program that predicts the range in which the next number in a sequence will fall. Given a series of numbers as input, the program calculates and outputs a predicted range for the subsequent number, leveraging statistical and probability principles.

### Objectives

The main goal of this project is to apply the math skills you have developed to build a program capable of making accurate predictions based on given numerical data. The program will take a number as standard input and print a range in which the next number should be located.

### Instructions

Your program should read input data where each line represents a number in a sequence. The line numbers represent the x-axis values, and the numbers themselves represent the y-axis values on a graph. Based on this input, your task is to calculate a range for the next expected number in the sequence.

#### Example

Given the following input data:
```bash
189
113
121
114
145
110
...
```
This data represents a graph where:

- The x-axis values correspond to the line numbers (0, 1, 2, 3, 4, 5, ...).
-  The y-axis values correspond to the input numbers (189, 113, 121, 114, 145, 110, ...).

```bash
$ ./your_program
189  # input number
120 200  # predicted range for the next input (113)
113  # input number
160 230  # predicted range for the next input (121)
121  # input number
110 140  # predicted range for the next input (114)
114  # input number
100 200  # predicted range for the next input (145)
145  # input number
1 99  # predicted range for the next input (110)
110  # input number
100 150  # predicted range for the next input
```
###
Key Considerations

    The ranges are not necessarily correct, as the example provided is for illustration purposes only. You are free to implement your own logic for calculating the ranges.
    The objective is to use the calculations and methodologies learned in the math-skills exercise to guess the number ranges effectively.

## Authors
- allkamau

