## Golang Gradien Formula

### Problem
Given an array of points where points[i] = [xi, yi] represents a point on the X-Y plane, 
return the maximum number of points that lie on the same straight line.

```
Input: points = [[1,1],[3,2],[5,3],[4,1],[2,3],[1,4]]
Output: 4

Input: points = [[1,1],[2,2],[3,3]]
Output: 3
```

### Solution
This program calculates the maximum number of points from an array that lie on the same line.
maxPoints is the main function that takes a 2D array where each 1D array represents a point in 2D space. 
The function loops over each pair of points and calculates the gradient (or slope) using 
the calculateGradient function. The result is then normalized to a positive zero or positive infinity if 
necessary using normalizeZeroAndInf.
