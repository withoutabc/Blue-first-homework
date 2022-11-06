package main

import "fmt"

type movie struct {
	Name            string
	Director        string
	Scriptwriter    string
	Protagonist     string
	Type            string
	ProducerCountry string
	Language        string
	ReleaseDate     string
	Time            string
	Mark            float32
}

func main() {
	n := movie{
		Name:         "杀人回忆",
		Director:     "奉俊昊",
		Scriptwriter: "奉俊昊/沈成宝/金光林",
		Protagonist: "宋康昊/金相庆/金雷夏/" +
			"宋在浩/边希峰/高瑞熙/" +
			"朴努植/朴海日/" +
			"全美善/徐永嬅/崔钟律/" +
			"刘承睦/申贤宗/李在应/" +
			"郑仁仙/吴龙/朴真宇/" +
			"朴泰京/沈成宝",
		Type:            "剧情/动作/悬疑/惊悚/犯罪",
		ProducerCountry: "韩国",
		Language:        "韩语/英语",
		ReleaseDate:     "2003-05-02(韩国)",
		Time:            "132分钟",
		Mark:            8.9,
	}
	fmt.Printf("请输入你的命令\n1.查看名字\n2.查看导演\n3.查看编剧\n4.查看主演\n5.查看类型\n6.查看制片国家\n7.查看语言\n8.查看发行日期\n9.查看时长\n10.查看评分\n11.退出程序\n")
	var option int
	for {
		fmt.Scan(&option)
		switch option {
		case 1:
			fmt.Println(n.Name)
		case 2:
			fmt.Println(n.Director)
		case 3:
			fmt.Println(n.Scriptwriter)
		case 4:
			fmt.Println(n.Protagonist)
		case 5:
			fmt.Println(n.Type)
		case 6:
			fmt.Println(n.ProducerCountry)
		case 7:
			fmt.Println(n.Language)
		case 8:
			fmt.Println(n.ReleaseDate)
		case 9:
			fmt.Println(n.Time)
		case 10:
			fmt.Println(n.Mark)
		case 11:
			return
		}
	}
}
