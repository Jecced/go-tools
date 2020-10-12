package encry

import (
	"github.com/Jecced/go-tools/src/randutil"
)

// 字节序列打乱
// @param data 数据
// @param seed 种子
// @param direction 打乱方向
// @return []byte 返回打乱后的字节数组
func BytesSeedSwap(data []byte, seed int64, direction bool) []byte {

	length := len(data)

	r := randutil.Random(seed)
	amount := r.NextInt(length/2, length)
	index := make([]int, 0, amount*2)

	for i := 0; i < amount*2; i++ {
		index = append(index, r.NextInt(0, length))
	}

	for i, l, c := 0, len(index)/2, 0; i < l; i++ {
		c = i
		if direction {
			c = l - i - 1
		}
		data[index[c]], data[index[c+1]] = data[index[c+1]], data[index[c]]
	}
	return data
}

// 修改字节方式的通用类型
// 操作一个二进制的通用定义方法, 入参一个byte, 返回修改后的byte
// @param b 将要修改的字节数据
// @return 修改后的字节数据
type ModifyByte func(b byte) byte

// 某多个位置二进制进行翻转
// @param mask 将一个二进制的第几位进行翻转, warn: mask取值范围是0~7
// @return 返回一个自定义修改函数 ModifyByte
func ByteFlips(masks ...uint8) ModifyByte {
	return func(value byte) byte {
		for _, mask := range masks {
			value = byteFlip(mask, value)
		}
		return value
	}
}

// 某1个位置二进制进行翻转
// @param mask 将一个二进制的第几位进行翻转, warn: mask取值范围是0~7
// @return 返回一个自定义修改函数 ModifyByte
func ByteFlip(mask uint8) ModifyByte {
	return ByteFlips(mask)
}

// 将某一个位置二进制进行翻转
func byteFlip(mask uint8, value byte) byte {
	flag := byte(1 << mask)
	if 0 == value&flag {
		return value | flag
	}
	return value & ^flag
}

// 某1个位置的二进制进行加入offset
// @param offset 将一个二进制的值 + offset, warn: 取值范围在-128~127之间
// @return 返回一个自定义修改函数 ModifyByte
func ByteAdd(offset int8) ModifyByte {
	return func(b byte) byte {
		return b + uint8(offset)
	}
}

// 自定义种子跳跃
// 每隔多少个字节[由随机数控制, 取值为1~10], 进行一次操作, 每次操作是 自定义的 m ModifyByte 操作
// @param data 数据
// @param seed 种子
// @param m 自定义字节操作
func BytesSeekSkip(data []byte, seed int64, m ModifyByte) []byte {
	r := randutil.Random(seed)
	for index, length := 0, len(data); index < length; index += r.NextInt(1, 10) {
		data[index] = m(data[index])
	}
	return data
}

// 固定跳跃修正
// 每隔多少个字节[固定值 @param x], 进行一次操作, 每次操作是 自定义的 m ModifyByte 操作
// @param data 数据
// @param x    固定跳跃多少个位置进行一次操作
// @param m 自定义字节操作
func BytesFixedSkip(data []byte, x int, m ModifyByte) []byte {
	if x <= 1 {
		x = 1
	}
	for index, length := 0, len(data); index < length; index += x {
		data[index] = m(data[index])
	}
	return data
}

// 所有位置进行更新
// 所有二进制, 都进行一次操作, 每次操作是 自定义的 m ModifyByte 操作
// @param data 数据
// @param m 自定义字节操作
func BytesUpdate(data []byte, m ModifyByte) []byte {
	return BytesFixedSkip(data, 1, m)
}

// 种子跳跃
// 每隔多少个字节[由随机数控制, 取值为1~10], 进行一次操作, 每次操作是固定的 += offset
// @param data 数据
// @param seed 种子
// @param offset 取到的 byte 进行 += offset
func BytesSeedSkipUp(data []byte, seed int64, offset int8) []byte {
	return BytesSeekSkip(data, seed, ByteAdd(offset))
}

// 固定跳跃
// 每隔多少个字节[固定值 @param x], 进行一次操作, 每次操作是固定的 += offset
// @param data 数据
// @param x    固定跳跃多少个位置进行一次操作
// @param offset 取到的 byte 进行 += offset
func BytesFixSkipUp(data []byte, x int, offset int8) []byte {
	return BytesFixedSkip(data, x, ByteAdd(offset))
}

// 所有位置进行跳跃
// 所有二进制, 都进行一次操作, 每次操作是固定的 += offset
// @param data 数据
// @param offset 取到的 byte 进行 += offset
func BytesUpdateUp(data []byte, offset int8) []byte {
	return BytesFixSkipUp(data, 1, offset)
}
