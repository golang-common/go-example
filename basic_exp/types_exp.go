/**
 * @Author: daipengyuan
 * @Description:
 * @File:  types_exp
 * @Version: 1.0.0
 * @Date: 2021/6/10 16:48
 */

package main

// Store 定义一个接口，实现该接口的条件为满足如下三个方法
type Store interface {
	CheckTaxCert() string      // 检查税务证
	CheckBusinessCert() string // 检查工商证
	Service() string           // 提供服务
}

// 张三的店子，满足了开店的条件(实现了接口)
type ZhangSanStore struct{}

func (z ZhangSanStore) CheckTaxCert() string {
	return "张三有税务证"
}

func (z ZhangSanStore) CheckBusinessCert() string {
	return "张三有工商证"
}

func (z ZhangSanStore) Service() string {
	return "你吃了一顿火锅"
}

// 李四的店子，也满足了开店的条件(实现了接口)
type LiSiStore struct{}

func (z LiSiStore) CheckTaxCert() string {
	return "李四的税务证过期了"
}

func (z LiSiStore) CheckBusinessCert() string {
	return "李四没有工商证"
}

func (z LiSiStore) Service() string {
	return "你买到一包假烟"
}

// 王五的店子，只有税务证，其它啥也没有
type WangWuStore struct {}

func (w WangWuStore)CheckTaxCert() string {
	return "王五有税务证"
}
//
//// 这个方法描述了你逛街的一天
//func main() {
//	// 一天，你去逛街，准备逛商店(满足以上3个条件)
//	var store Store
//	// 你来到一家店，走进去，原来是张三的店
//	store = ZhangSanStore{}
//	// 你看了看墙上的工商税务证，没啥问题，应该很正规
//	fmt.Println(store.CheckTaxCert())      // Output:张三有税务证
//	fmt.Println(store.CheckBusinessCert()) // Output:张三有工商证
//	// 你在店里享受了服务
//	fmt.Println(store.Service()) // Output: 你吃了一顿火锅
//
//	// 出来后，你换了一家店，李四的店
//	store = LiSiStore{}
//	// 你看了看墙上的工商税务证，感觉有点不对劲
//	fmt.Println(store.CheckTaxCert())      // Output:李四的税务证过期了
//	fmt.Println(store.CheckBusinessCert()) // Output:李四没有工商证
//	// 虽然有疑问，但是你还是消费了一把，果然不对劲
//	fmt.Println(store.Service()) // Output: 你买到一包假烟
//}
