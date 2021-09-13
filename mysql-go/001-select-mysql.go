package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //导入包但不使用，init()
)

type user struct {
	id                   int
	company              string
	company_age          string
	company_age_month    int
	deleted              int
	department           string
	department_id        int
	enter_date           string
	id_card              string
	leader_id            int
	leave_date           string
	level                string
	level_value          int
	post                 string
	post_code            int
	probation_start_date string
	realname             string
	title                string
	user_id              int
}

var db *sql.DB //连接池对象
func initDB() (err error) {
	//数据库
	//用户名:密码啊@tcp(ip:端口)/数据库的名字
	dsn := "root:123456@tcp(127.0.0.1:3306)/salary"
	//连接数据集
	db, err = sql.Open("mysql", dsn) //open不会检验用户名和密码
	if err != nil {
		return
	}
	err = db.Ping() //尝试连接数据库
	if err != nil {
		return
	}
	fmt.Println("连接数据库成功~")
	//设置数据库连接池的最大连接数
	db.SetMaxIdleConns(10)
	return
}

func queryAll() {
	//1.sql语句
	sqlStr := "select id, company, company_age, company_age_month, deleted, department, department_id, enter_date, id_card, leader_id, leave_date, level, level_value, post, post_code, probation_start_date, realname, title, user_id from users;"
	//2.执行
	rows, err := db.Query(sqlStr)
	if err != nil {
		fmt.Printf("%s query failed,err:%v\n", sqlStr, err)
		return
	}
	//3一定要关闭rows
	defer rows.Close()
	//4.循环取值
	for rows.Next() {
		var u1 user
		rows.Scan(&u1.id, &u1.company, &u1.company_age, &u1.company_age_month, &u1.deleted, &u1.department, &u1.department_id, &u1.enter_date, &u1.id_card, &u1.leader_id, &u1.leave_date, &u1.level, &u1.level_value, &u1.post, &u1.post_code, &u1.probation_start_date, &u1.realname, &u1.title, &u1.user_id)
		fmt.Printf("u1:%#v\n", u1)
	}
}

//Go连接Mysql示例
func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed,err%v\n", err)
	}
	//查询多行
	queryAll()
}
