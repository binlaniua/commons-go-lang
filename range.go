package commons_go_lang

//
//
// Range is simple for loop, when callback return true, the loop break
//
//
func Range(start, end, step int64, callback func(int64) bool) {
	for ; start < end; start += step {
		isInterrupt := callback(start)
		if isInterrupt {
			return
		}
	}
}

//
//
// RangeEnd is start from 1, and step is 1
//
//
func RangeExt(end int64, cb func(int64) bool) {
	Range(1, end, 1, cb)
}
