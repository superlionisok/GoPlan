package main

import (
	"CreatePlan/Models"
	"fmt"
	"strconv"
	"time"
)

func main() {
	fmt.Println("start")

	// insert 1
	//TestInsertOne()
	//TestInsertMore()
	//TestFindOne()
	TestFindList()

	fmt.Println("end")
}
func TestInsertOne() {
	var add Models.UserCreatePlan
	add.Name = "cname"
	id, err := Models.Engine.Insert(&add)

	if err != nil {
		fmt.Println("insert one has err.err=", err.Error())
		return
	}

	fmt.Println("insert one id = ", id)

}

func TestInsertMore() {

	var users []Models.UserCreatePlan
	for i := 0; i < 2000; i++ {
		var add Models.UserCreatePlan
		add.Name = "cname" + strconv.Itoa(i)
		add.CreateTime = time.Now()
		add.UpdateTime = time.Now()
		users = append(users, add)
	}

	var t1 = time.Now()
	id, err := Models.Engine.Insert(&users)

	if err != nil {
		fmt.Println("insert more has err.err=", err.Error())
		return
	}
	var t2 = time.Now()
	var timespan = t2.Sub(t1)
	fmt.Println("need time =", timespan.Seconds())
	fmt.Println("insert more id = ", id)

}

func TestFindOne() {
	var id = 1
	user := new(Models.UserCreatePlan)
	has, err := Models.Engine.Id(id).Get(user)

	if err != nil {
		fmt.Println("find one err. err=", err.Error())
		return
	}
	fmt.Println("id = 1 data is have = ", has)
	fmt.Println(user)

}

func TestFindList() {

	//has, err := Models.Engine.Id(id).Get(user)
	users := make([]Models.UserCreatePlan, 0)
	err := Models.Engine.Where(" name like ?", "%cname%").Desc("i_d").Limit(20, 0).Find(&users)
	if err != nil {
		fmt.Println("find one err. err=", err.Error())
		return
	}
	fmt.Println(len(users))
	fmt.Println(users)

	has, err := Models.Engine.Where("name = ?", "cname1").Exist(&Models.UserCreatePlan{})
	if err != nil {
		fmt.Println("Exist one err. err=", err.Error())
		return
	}
	fmt.Println("Exist one is have =", has)

	var user Models.UserCreatePlan
	total, err := Models.Engine.Where("i_d >?", 1).Count(user)

	fmt.Println("Count =", total)

}
