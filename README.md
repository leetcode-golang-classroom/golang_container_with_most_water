# golang_container_with_most_water

You are given an integer array `height` of length `n`. There are `n` vertical lines drawn such that the two endpoints of the `ith` line are `(i, 0)` and `(i, height[i])`.

Find two lines that together with the x-axis form a container, such that the container contains the most water.

Return *the maximum amount of water a container can store*.

**Notice** that you may not slant the container.

## Examples

**Example 1:**

![https://s3-lc-upload.s3.amazonaws.com/uploads/2018/07/17/question_11.jpg](https://s3-lc-upload.s3.amazonaws.com/uploads/2018/07/17/question_11.jpg)

```
Input: height = [1,8,6,2,5,4,8,3,7]
Output: 49
Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.

```

**Example 2:**

```
Input: height = [1,1]
Output: 1

```

**Constraints:**

- `n == height.length`
- `2 <= n <= 105`
- `0 <= height[i] <= 104`

## 解析

給定一個非負整數陣列 height

其中每個 height[i] 代表 index i 的格版高度

假設 每兩個 height 之間相隔為單位 1 

往中間注水 假設格版本身很薄 所以不計

則兩個 height 所能裝的水的面積 剛好是由比較小的 height * 兩個 height 的 X軸相查距離 = min(height[i], height[j]) *(j -i) if i < j


![](https://i.imgur.com/Uqn879I.png)


要求寫一個演算法來計算再給定的 height 中所能注入最多面積是多少

由於每個面積都是由兩個 height 來圍出

所以可以透過 two pointer 然後逐步更新最短的 height 的 point 來找尋下一個更多面積的可能值

直到兩個 height 只差 1 

而再這之間不斷把當下最大的做紀錄

初始化 lp =0, rp = len(height), maxArea = 0

當 lp < rp 做以下運算

   area = min(height[lp], height[rp]) *(rp-lp)

 如果 area > maxArea,  則更新 更新 maxArea = area

 如果 height[lp] > height[rp], 更新 rp -= 1 否則更新 lp += 1

![](https://i.imgur.com/0ovM7Ly.png)

## 程式碼
```go
package sol

func maxArea(height []int) int {
	mArea := 0
	lp, rp := 0, len(height)-1
	var min = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	for lp < rp {
		area := min(height[lp], height[rp]) * (rp - lp)
		if mArea < area {
			mArea = area
		}
		if height[lp] > height[rp] {
			rp--
		} else {
			lp++
		}
	}
	return mArea
}

```
## 困難點

1. 要看出與面積相關的計算方式
2. 要看出如何找下一個可能的更大的面積

## Solve Point

- [x]  初始化 rp =0, lp = len(height) - 1, maxArea = 0
- [x]  當 rp < lp 做以下運算
- [x]  area = min(height[lp], height[rp]) *(rp-lp)
- [x]  if area > maxArea 更新 maxArea = area
- [x]  當 height[lp] > height[rp] 更新 rp -= 1 否則更新 lp += 1
- [x]  回傳 maxArea