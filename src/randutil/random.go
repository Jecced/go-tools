package randutil

// 定义随机数
type Random int64

// 进入下一个种子序列
func (r *Random) next() float64 {
	*r = (*r*9301 + 49297) % 233280
	return float64(*r) / 233280.0
}

// 下一个浮点数
func (r *Random) Next(min, max int64) float64 {
	if min > max {
		min, max = max, min
	}
	rnd := r.next()
	return float64(min) + rnd*float64(max-min)
}

// 下一个整形
func (r *Random) NextInt(min, max int) int {
	return int(r.Next(int64(min), int64(max)))
}

// 下一个Int32
func (r *Random) NextInt32(min, max int32) int32 {
	return int32(r.Next(int64(min), int64(max)))
}

// 下一个Int64
func (r *Random) NextInt64(min, max int64) int64 {
	return int64(r.Next(min, max))
}

// 下一个bool
func (r *Random) NextBool() bool {
	return r.Next(0, 1) > 0.5
}

// 重新设置种子
func (r *Random) SetSeed(seed int64) {
	*r = Random(seed)
}

// 获取种子
func (r *Random) GetSeed() int64 {
	return int64(*r)
}
