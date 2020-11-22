package test

import (
	"fmt"
	"github.com/Jecced/go-tools/src/commutil"
	"github.com/Jecced/go-tools/src/encrypt"
	"testing"
)

func TestEncryption(t *testing.T) {
	seed := int64(2112)
	fmt.Println("种子:", seed)

	str := "the weather are very nice, enjoy it."

	data := []byte(str)
	fmt.Println("原文", string(data))
	{
		// 顺序打乱
		data = encrypt.BytesSeedSwap(data, seed, true)
		// 所有位置, 固定+5
		data = encrypt.BytesUpdateUp(data, 5)
		// 种子方式 第三位置取反
		data = encrypt.BytesSeekSkip(data, seed, encrypt.ByteFlip(3))
		// base64
		str = commutil.EncodeBase64(&data)
		fmt.Println("密文", str)
	}

	{
		// base64
		data, _ = commutil.DecodeBase64(str)
		// 种子方式, 第三位置取反
		data = encrypt.BytesSeekSkip(data, seed, encrypt.ByteFlip(3))
		// 所有位置, 固定-5
		data = encrypt.BytesUpdateUp(data, -5)
		// 逆序打乱
		data = encrypt.BytesSeedSwap(data, seed, false)
		fmt.Println("还原", string(data))
	}

}
