package main

import (
	"Collection/SysModels"
	"Collection/models"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/antchfx/htmlquery"
	"github.com/astaxie/beego/orm"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"golang.org/x/net/html"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("start")
	DoBeeOrm()
	//DoPost()
	//MyInsertMore()
	//DoCollect()

	//Get2()
	//htmlquerypath()
	fmt.Println("end")

}
func init() {

	//	orm.RegisterDataBase("default", "mysql", "root:123456@/superplan?charset=utf8")
	//	orm.RegisterDataBase("default2", "mysql", "root:123456@/lionjihua?charset=utf8")
	// 需要在init中注册定义的model

}
func DoBeeOrm() {
	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库
	var u models.User
	u.LoginName = "default-user"
	u.LoginPwd = "123"
	u.CreateTime = time.Now()

	ic, err := o.Insert(&u)
	if err != nil {

		fmt.Println("default 插入失败")
		return
	}
	fmt.Println(ic)

	o.Using("default2")

	var o2 SysModels.UserOrder1
	o2.OrderTitle = "ordertest"
	ic2, err := o.Insert(&o2)
	if err != nil {
		fmt.Println("o2 插入失败")
		return
	}
	fmt.Println(ic2)
}

func DoCollect() {

	//https://www.600wcp.com/ssc/zst/ffssc.html?type=1
	var strurl = "https://www.600wcp.com/ssc/zst/ffssc.html?type=1"
	resp, err := http.Get(strurl)
	if err != nil {
		fmt.Println("http.Get err = ", err)
		return
	}
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	doc.Find(".cl-31").Each(func(i int, s *goquery.Selection) {

		fmt.Println(s.Html())
		band := s.Find("div").Text()
		// title := s.Find("i").Text()
		fmt.Printf("Review %d: %s - %s\n", i, band)
	})

}

func HttpGet(url string) (result string, err error) {
	resp, err1 := http.Get(url)
	if err != nil {
		err = err1
		return
	}
	if resp.StatusCode != 200 {
		fmt.Println("err")
	}

	defer resp.Body.Close()
	//读取网页的body内容
	buf := make([]byte, 4*1024)
	for true {
		n, err := resp.Body.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("文件读取完毕")
				break
			} else {
				fmt.Println("resp.Body.Read err = ", err)
				break
			}
		}
		result += string(buf[:n])
	}
	return
}

func Get2() {
	url := "https://www.600wcp.com/ssc/zst/ffssc.html?type=1"
	doc, err := goquery.NewDocument(url)
	if err != nil {

		fmt.Println("get document error:", err)
		return
	}
	doc.Find("div").EachWithBreak(func(i int, s *goquery.Selection) bool {
		//d := s.Eq(0).Find("td")
		// 每个first tr标签下面就只有一个td节点集
		fmt.Println(s.Children().Text())
		// 遍历孩子节点，需要中断跳出，所以用了EachWithBreak
		s.Children().EachWithBreak(func(j int, selection *goquery.Selection) bool {
			//fmt.Println(selection.Text())
			// 获取内容
			// str := selection.Text()
			// currencyIndexList = append(currencyIndexList, util.FindNumbers(str))
			if j == 5 {
				return false
			}
			return true
		})
		if i == 0 {
			return false
		}
		return true
	})
	return

}

func htmlfetch(url string) *html.Node {
	log.Println("Fetch Url", url)
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Http get err:", err)
	}
	if resp.StatusCode != 200 {
		log.Fatal("Http status code:", resp.StatusCode)
	}
	defer resp.Body.Close()
	doc, err := htmlquery.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return doc
}
func htmlquerypath() {
	start := time.Now()
	ch := make(chan bool)
	for i := 0; i < 1; i++ {
		go htmlparseUrls("https://movie.douban.com/top250?start="+strconv.Itoa(25*i), ch)
	}

	for i := 0; i < 10; i++ {
		<-ch
	}

	elapsed := time.Since(start)
	log.Printf("Took %s", elapsed)
}

func htmlparseUrls(url string, ch chan bool) {
	doc := htmlfetch(url)
	pic := htmlquery.Find(doc, `div`)
	nodes := htmlquery.Find(doc, `div`)
	for key, node := range nodes {
		var n = node
		fmt.Println(n)

		num := htmlquery.FindOne(pic[key], `./em[@class=""]/text()`)
		url := htmlquery.FindOne(node, "./a/@href")
		title := htmlquery.FindOne(node, `.//span[@class="title"]/text()`)
		log.Println(htmlquery.InnerText(num),
			strings.Split(htmlquery.InnerText(url), "/")[4],
			htmlquery.InnerText(title))
	}
	time.Sleep(2 * time.Second)
	ch <- true
}

func MoreDo() {
	var strurl = "https://www.600wcp.com/ssc/zst/ffssc.html?type=1"
	resp, err := http.Get(strurl)
	if err != nil {
		fmt.Println("http.Get err = ", err)
		return
	}
	if resp.StatusCode != 200 {
		fmt.Println("err")
	}
	defer resp.Body.Close()

	buf := make([]byte, 4*1024)
	for true {
		n, err := resp.Body.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("文件读取完毕")
				break
			} else {
				fmt.Println("resp.Body.Read err = ", err)
				break
			}
		}

		var result = ""
		result += string(buf[:n])
	}

}

func DoPost() {

	url := "https://www.600wcp.com/ssc/ajaxGetHistory.json?timestamp=1564633192573" //请求地址
	contentType := "application/x-www-form-urlencoded; charset=UTF-8"
	//参数，多个用&隔开
	data := strings.NewReader("pageIndex=1&playGroupId=15&pageSize=2")
	resp, err := http.Post(url, contentType, data)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(body))

	var f interface{}
	err = json.Unmarshal(body, &f)
	if err != nil {
		fmt.Println("序列化失败", err)
	}
	//{
	//    "result": 1,
	//    "sscHistoryList": [
	//        {
	//            "playGroupName": "分分时时彩",
	//            "playGroupId": 15,
	//            "number": "201908010756",
	//            "openCode": "9,8,2,4,2",
	//            "openTime": 1564634100000,
	//            "date": "2019-08-01"
	//        }
	//    ]
	//}
	m := f.(map[string]interface{})

	var r1 = m["result"]
	fmt.Println(r1)
	var r2 = m["sscHistoryList"]
	fmt.Println(r2)
	var datas = r2.([]interface{})
	var d = datas[0].(map[string]interface{})
	var number = d["number"]
	var opencode = d["openCode"]
	var opentime = d["openTime"]

	var a = opentime.(float64)

	fmt.Println(a)

	var aaa = int64(a)
	//string =string[0:10]
	fmt.Println(aaa)
	bbb := aaa / 1000

	var t = time.Unix(bbb, 0)
	fmt.Println(number, opencode, t)
}

func MyInsertMore() {

	db, err := gorm.Open("mysql", "root:123456@tcp(10.10.15.202:3306)/superplan")
	//db, err := gorm.Open("mysql", "root:QAZxsw2@/superplan?charset=utf8&parseTime=True&loc=10.10.15.202:3306")
	//db.LogMode(true)
	if err != nil {
		fmt.Println("connect db err=", err.Error())
		panic("failed to connect database")
	}
	defer db.Close()
	db.SingularTable(true)

	// 查看表

	hava := db.HasTable(&models.SysUser{})
	if hava == true {
		fmt.Println("有表")
	} else {
		var cres = db.CreateTable(&models.SysUser{})
		fmt.Println("建表=", cres)
	}
	db.AutoMigrate(&models.SysUser{})

	// 使用gorm插入一条
	var user models.SysUser
	user.Name = "mytest"

	a1 := db.Create(&user)
	fmt.Println("a1=", a1)
	aa := db.NewRecord(user) // => 创建`user`后返回`false`
	fmt.Println("aa=", aa)
	fmt.Println("add one id=", user.ID)

	// 批量插入
	var users []models.SysUser

	for i := 0; i < 10; i++ {
		var add models.SysUser
		add.Name = "mytest" + strconv.Itoa(i)
		users = append(users, add)
	}
	fmt.Println(users)
	db.Create(&users)
	// var strsql=" ins"

}
