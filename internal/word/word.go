package word

import (
	"strings"
	"unicode"
)

func ToUpper(s string)string {
	return  strings.ToUpper(s)
}

func ToLower(s string)string {
	return strings.ToLower(s)
}

//下划线单词转为大写驼峰单词
func UndersoreToUpperCamelCase(s string) string{
	s = strings.Replace(s,"_","",-1)
	s = strings.Title(s)
	return strings.Replace(s," ","",-1)
}

//下划线单词转小写驼峰单词
func UnderScoreToLowerCamelCase(s string) string {
	s = UndersoreToUpperCamelCase(s)
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}

//驼峰单词转下划线单词
func CamelCaseToUnderscore(s string) string {
	var output []rune
	for i ,r := range  s {
		if i == 0 {
			output = append(output,unicode.ToLower(r))
			continue
		}
		//这个用来判断是否是大写字母是的话 就加一个_ 然后在加字母
		if unicode.IsUpper(r) {
			output = append(output,'_')
		}
		output = append(output,unicode.ToLower(r))
	}
	return  string(output)
}

