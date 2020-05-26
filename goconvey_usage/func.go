package goconvey_usage

// Leetcode算法题：寻找字符串中连续不重复的最长子串，并返回其长度。

func maxLengthOfNonRepeatSubStr(s string) int {
	//遍历字符串，将每个遍历到的字符作为key，其索引作为value，更新入map。

	lastOccuredMap := make(map[byte]int)
	start := 0  //所遍历的字符其开始位置
	maxLength := 0 //最长不重复子串长度

	for i, ch := range []byte(s) {

		//更新字符起始位
		lastI, exists := lastOccuredMap[ch]
		if exists && lastI >= start {
			start = lastI + 1
		}

		//当遍历到的字符距离start的间隔比之前存的maxLength大时，更新maxLength
		//这是因为当遍历到不重复字符时，没有啥操作，此时就应该更新maxLength
		if i-start+1 > maxLength {
			maxLength = i-start+1
		}

		//写入/更新map
		lastOccuredMap[ch] = i
	}

	return maxLength
}
