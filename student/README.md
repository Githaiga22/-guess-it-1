# Guess-It-1
Overview

The guess-it-1 project is designed to challenge your mathematical and statistical skills by creating a program that predicts the range in which the next number in a sequence will fall. Given a series of numbers as input, the program calculates and outputs a predicted range for the subsequent number, leveraging statistical and probability principles.

## How It Works

### Overview

When you input a sequence of numbers, the program calculates a range where the next number is likely to fall. This range is calculated based on the recent numbers you've entered, ensuring that the predictions adapt to new data.

### Mathematical Concepts

To understand how the program works, let's break down the key mathematical concepts used:

#### 1. Mean (Average)
- **What It Is**: The mean is the average value of a set of numbers. It gives you a general idea of where the center of the numbers lies.
- **How It's Used**: In this project, the mean helps us identify the middle point of the recent numbers you've entered.

  **Example**:
  - If you have the numbers [100, 120, 140], the mean would be:
    \[
    \text{Mean} = \frac{100 + 120 + 140}{3} = 120
    \]
  - This mean (120) tells us that the average value of these numbers is 120.

#### 2. Variance
- **What It Is**: Variance measures how spread out the numbers are from the mean. A small variance means the numbers are close to the mean, while a large variance means they are spread out.
- **How It's Used**: In the program, variance helps us understand how much the numbers vary from the average, which is crucial for calculating the range.

  **Example**:
  - For the numbers [100, 120, 140], the variance is calculated as follows:
    1. Find the squared differences from the mean:
       - (100 - 120)² = 400
       - (120 - 120)² = 0
       - (140 - 120)² = 400
    2. Average these squared differences:
       \[
       \text{Variance} = \frac{400 + 0 + 400}{3} = \frac{800}{3} \approx 266.67
       \]
  - A variance of 266.67 tells us that the numbers are somewhat spread out from the mean.

#### 3. Standard Deviation
- **What It Is**: The standard deviation is the square root of the variance. It gives a more understandable measure of how spread out the numbers are.
- **How It's Used**: The standard deviation helps us create a range around the mean where the next number is likely to fall.

  **Example**:
  - Continuing from our variance example, the standard deviation is:
    \[
    \text{Standard Deviation} = \sqrt{266.67} \approx 16.33
    \]
  - This value tells us that most numbers are within 16.33 units of the mean.

#### 4. Confidence Interval
- **What It Is**: A confidence interval is a range that we believe contains the next number with a certain level of confidence. In this case, we use a 99.99% confidence interval.
- **How It's Used**: We calculate the confidence interval by taking the mean and adding/subtracting 3.89 times the standard deviation. This gives us a range where the next number is very likely to fall.

  **Example**:
  - For a mean of 120 and a standard deviation of 16.33, the range is:
    \[
    \text{Lower Bound} = 120 - 3.89 \times 16.33 \approx 55
    \]
    \[
    \text{Upper Bound} = 120 + 3.89 \times 16.33 \approx 185
    \]
  - So, the program would predict that the next number will likely fall between 55 and 185 with a 99.99% confidence.

### Code Explanation

- **Window Size**: We look at the last 10 numbers to make our predictions. This is called the "window size." You can adjust this to make the predictions more or less sensitive to recent changes.
  
- **Mean and Variance**: The program calculates the mean and variance of the numbers in the window. It then calculates the standard deviation from the variance.

- **Confidence Interval**: Finally, the program uses the mean and standard deviation to create a range (lower and upper bounds) where the next number is likely to be, with a 99.99% (3.89)confidence level.

### Running the Program

To run the program, simply input the first two numbers, one at a time. The program will output the predicted range for the next number based on the numbers you’ve entered so far.

```bash
go run main.go
```
### Example Output
```bash
100
120
70 150
134
62 173
145
57 192
```
### Learning Resources

If you want to learn more about the mathematical concepts used in this project, here are some beginner-friendly resources:

1. Mean (Average):
        [Khan Academy](https://www.khanacademy.org/math/statistics-probability/describing-and-visualizing-data)
 2. Variance and Standard Deviation:
        [Khan Academy](https://www.khanacademy.org/math/statistics-probability/summarizing-quantitative-data)
3. Confidence Intervals:
        [Khan Academy](https://www.khanacademy.org/math/statistics-probability/confidence-intervals-one-sample)

- These resources will help you understand the foundational concepts behind this project and how they are applied in statistics and data analysis.
## Authors
- allkamau

