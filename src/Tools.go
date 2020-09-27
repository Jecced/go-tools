package gts

import "github.com/Jecced/go-tools/src/translate"

// 使用谷歌翻译, 英文翻译中文
func TranslateEn2Cn(text string) string {
	return translate.GoogleTranslate(text)
}
