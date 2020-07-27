package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"strconv"
	"strings"
	"test/internal/timer"
	"time"
)

var (
	calculateTime string
	duration string
	timeCmd = &cobra.Command{
		Use: "time",
		//显示在主命令里面的help
		Short: "时间格式处理",
		//显示在子命令里面的help
		Long: "时间格式处理",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	nowTimeCmd = &cobra.Command{
		Use: "now",
		Short: "获取当前时间",
		Long: "获取当前时间",
		Run: func(cmd *cobra.Command, args []string) {
			nowTime := timer.GetNowTime()
			//format 按照既定的格式初始化
			//unix返回unix时间 即时间戳 该值为UTC 1970年1月1日起经过的秒数
			//如果要用其他的时间格式，可以用内部定义的格式
			/*
			const (
				ANSIC       = "Mon Jan _2 15:04:05 2006"
				UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
				RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
				RFC822      = "02 Jan 06 15:04 MST"
				RFC822Z     = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
				RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
				RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
				RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
				RFC3339     = "2006-01-02T15:04:05Z07:00"
				RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
				Kitchen     = "3:04PM"
				// Handy time stamps.
				Stamp      = "Jan _2 15:04:05"
				StampMilli = "Jan _2 15:04:05.000"
				StampMicro = "Jan _2 15:04:05.000000"
				StampNano  = "Jan _2 15:04:05.000000000"
			)
			 */
			log.Printf("输出结果 %s ,%d",nowTime.Format("2006-01-02 15:04:05"),nowTime.Unix())
		},
	}

	calculateTimeCmd = &cobra.Command{
		Use: "calc",
		Short: "计算所需时间",
		Long: "计算所需时间",
		Run: func(cmd *cobra.Command, args []string) {
			//在这段代码中 展示了三种常用的时间格式处理，分别是 时间戳 2006-01-02 年月日 和2006-01-02 15:04:05  年月日时分秒
			var (
				currentTimer time.Time
				layout = "2006-01-02 15:04:05"
			)
			if calculateTime == "" {
				currentTimer = timer.GetNowTime()
			}else {
				var err error
				//时间处理上 加入了字符串判断 是否又空字符串 如果有的话则按照2006-01-02 15:04:05  没有的话2006-01-02
				//这里如果包含空 那么会返回true  那么!true 就是false
				if !strings.Contains(calculateTime," ") {
					layout = "2006-01-02"
				}
				//如果按照layout的我们做格式化，出现问题就直接取时间戳
				currentTimer,err = time.Parse(layout,calculateTime)
				if err != nil {
					t,_ := strconv.Atoi(calculateTime)
					//取时间戳
					currentTimer = time.Unix(int64(t),0)
				}
			}
			calculateTime,err := timer.GetCalculateTime(currentTimer,duration)
			if err != nil {
				log.Fatalf("time.GetCalcuateTime err : %v\n",err)
			}
			log.Printf("输出结果: %s %d\n",calculateTime.Format(layout),calculateTime.Unix())
		},
	}
)


func init() {
	timeCmd.AddCommand(nowTimeCmd,calculateTimeCmd)
	calculateTimeCmd.Flags().StringVarP(&calculateTime,"calculate","c","",`需要计算的时间，有效单位位置时间戳或已经格式化后的时间`)
	calculateTimeCmd.Flags().StringVarP(&duration,"duration","d","",`持续时间，有效单位为ns,us,ms,s,m,h,`)
}



