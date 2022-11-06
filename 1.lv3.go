// 没玩过英雄联盟，用数码宝贝替代了QAQ
// 输入的数字不在范围内没有打印error,实在不想写了
package main

import "fmt"

type DIGIMON interface {
	Say()
	ReleaseSkill()
	UseItems()
}

type Digimon struct {
	Name string
	HP   int
	MP   int
	ATK  int
	DEF  int
}

func (d *Digimon) Attack() {
	fmt.Printf("%v进行了攻击", d.Name)
}

func (d *Digimon) ReleaseSkill(a string) {
	fmt.Printf("%v使用了技能 %v", d.Name, a)
}
func (d *Digimon) UseItems(c string) {
	fmt.Printf("%v使用了道具 %v\n", d.Name, c)
}
func main() {

	for {
		m := map[int]string{
			1: "回血药剂",
			2: "回蓝药剂",
			3: "攻击果实",
			4: "防御果实",
		}
		Numemon := Digimon{Name: "鼻涕兽", HP: 101, MP: 105, ATK: 26, DEF: 15}
		Megadramon := Digimon{Name: "百万龙兽", HP: 115, MP: 126, ATK: 33, DEF: 24}
		DarkTyranomon := Digimon{Name: "黑暗巨龙兽", HP: 108, MP: 128, ATK: 24, DEF: 34}
		fmt.Printf("请输入你的命令\n1.查看数码宝贝列表\n2.查看物品列表\n3.退出程序\n")
		var option, x, y, z int
		fmt.Scan(&option)
		if option == 1 {
			fmt.Printf("请选择要进行操作的数码宝贝\n1.%v\n2.%v\n3.%v\n", Numemon.Name, Megadramon.Name, DarkTyranomon.Name)
			var a, b, c, d, e, f, g int
			fmt.Scan(&a)
			fmt.Printf("1.攻击\n2.释放技能\n3.查看可对其使用的道具列表\n")
			switch a {
			case 1:
				fmt.Scan(&b)
				switch b {
				case 1:
					Numemon.Attack()
				case 2:
					Numemon.ReleaseSkill("大便投掷")
				case 3:
					fmt.Printf("请选择要使用的道具\n1.回血药剂\n2.回蓝药剂\n3.攻击果实\n4.防御果实\n")
					fmt.Scan(&e)
					DarkTyranomon.UseItems(m[e])
					fmt.Printf("是否返回？\n1.返回\n2.退出\n")
				}
				fmt.Scan(&z)
				switch z {
				case 1:
					continue
				case 2:
					return
				}
			case 2:
				fmt.Scan(&c)
				switch c {
				case 1:
					Megadramon.Attack()
				case 2:
					Megadramon.ReleaseSkill("究极斩切")
				case 3:
					fmt.Printf("请选择要使用的道具\n1.回血药剂\n2.回蓝药剂\n3.攻击果实\n4.防御果实\n")
					fmt.Scan(&f)
					DarkTyranomon.UseItems(m[f])
				}
				fmt.Printf("是否返回？\n1.返回\n2.退出\n")
				fmt.Scan(&x)
				switch x {
				case 1:
					continue
				case 2:
					return
				}
			case 3:
				fmt.Scan(&d)
				switch d {
				case 1:
					DarkTyranomon.Attack()
				case 2:
					DarkTyranomon.ReleaseSkill("火焰爆风")
				case 3:
					fmt.Printf("请选择要使用的道具\n1.回血药剂\n2.回蓝药剂\n3.攻击果实\n4.防御果实\n")
					fmt.Scan(&g)
					DarkTyranomon.UseItems(m[g])
				}
				fmt.Printf("是否返回？\n1.返回\n2.退出\n")
				fmt.Scan(&y)
				switch y {
				case 1:
					continue
				case 2:
					return
				}
			}
		} else if option == 2 {
			var a int
			fmt.Printf("1.回血药剂\n2.回蓝药剂\n3.攻击果实\n4.防御果实\n")
			fmt.Printf("是否返回？\n1.返回\n2.退出\n")
			fmt.Scan(&a)
			switch a {
			case 1:
				continue
			case 2:
				return
			}
		} else if option == 3 {
			return
		}
		continue
	}
}
