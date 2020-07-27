package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"strings"
	"test/internal/word"
)

var (
	str string
	mode int8
)

var wordCmd = &cobra.Command{
	Use: "word",
	Short: "单词格式转换",
	Long: desc,
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		switch mode {
		case ModeUpper:
			content = word.ToUpper(str)
		case ModeLower:
			content = word.ToLower(str)
		case ModeUnderscodeToUpperCamlcase:
			content = word.UndersoreToUpperCamelCase(str)
		case ModeUnderscodeToLowerCamlcase:
			content = word.UnderScoreToLowerCamelCase(str)
		case ModeCamelcaseToUndescore:
			content = word.CamelCaseToUnderscore(str)
		default:
			log.Fatalf("暂时不支持，请htlp word查看")
		}
		log.Printf("输出结果:%s\n",content)
	},
}

func init() {
	//varp 系列的方法 第一个参数是需要绑定的变量，第二个参数是接受该参数的完整的命令标志，第三个参数是对应的短标识，第四个参数为默认值，第五个参数为使用说明
	wordCmd.Flags().StringVarP(&str,"str","s","","请输入单词内容")
	wordCmd.Flags().Int8VarP(&mode,"mode","m",0,"请输入单词转义的模式")
}

const (
	ModeUpper = iota + 1
	ModeLower
	ModeUnderscodeToUpperCamlcase
	ModeUnderscodeToLowerCamlcase
	ModeCamelcaseToUndescore
)

var desc = strings.Join([]string{
	"该子命令支持各种单词格式转换，模式如下",
	"1:全部转为大写",
	"2:全部转为小写",
	"3:下划线转为大写驼峰",
	"4:下划线转为小谢驼峰",
	"5:驼峰转为下划线",
	},"\n")

