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
