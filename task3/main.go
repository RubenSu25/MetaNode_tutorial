package main

import (
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 题目1：基本CRUD操作
// 假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、 age （学生年龄，整数类型）、
//  grade （学生年级，字符串类型）。

type Student struct {
	Id    int
	Name  string
	Age   int
	Grade string
}

func run1(db *gorm.DB) {
	db.AutoMigrate(&Student{})

	// 编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
	student := Student{Name: "张三", Age: 20, Grade: "三年级"}
	db.Create(&student)

	// 编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
	students := []Student{}
	db.Where("age > ?", 18).Find(&students)

	// 编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
	result := db.Model(&Student{}).Where("name = ?", "张三").Update("grade", "四年级")
	fmt.Printf("result.Row(): %v\n", result.Row())
	fmt.Println(result.Rows())
	fmt.Printf("result.Error: %v\n", result.Error)
	db.Model(Student{}).Where("name = ?", "张三").Update("grade", "四年级")

	s1 := Student{Id: 1}
	db.Model(&s1).Update("grade", "四年级")
	db.Exec("update students set grade =? where name = ?", "四年级", "张三")

	// 编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
	db.Delete(&Student{}, "age < ?", 15)
	db.Where("age < ?", 15).Delete(&Student{})

}

// 假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和 transactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
// 要求 ：
// 编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。

type Account struct {
	Id      int
	Balance float64
}

type Transaction struct {
	Id            int
	Amount        float64
	FromAccountId int
	ToAccountId   int
}

func run2(db *gorm.DB) {

	db.AutoMigrate(&Account{}, &Transaction{})

	a := []Account{{Balance: 100}, {Balance: 100}}

	db.Create(&a)

}

func transfer(aId, bId int, db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {

		a := &Account{Id: aId}
		result := db.Find(&a)
		if result.RowsAffected == 0 {
			return errors.New("账户不存在")
		} else if result.Error != nil {
			return result.Error
		}

		var b Account
		result = db.First(&b, bId)
		if result.Error != nil {
			return result.Error
		}

		if a.Balance < 100 {
			return errors.New("余额不足")
		}

		result = tx.Model(&a).Update("balance", gorm.Expr("balance -100"))
		if result.Error != nil {
			return result.Error
		}
		result = tx.Model(&b).Update("balance", gorm.Expr("balance +100"))
		if result.Error != nil {
			return result.Error
		}

		t := &Transaction{Amount: 100, FromAccountId: a.Id, ToAccountId: b.Id}
		result = tx.Create(t)
		if result.Error != nil {
			return result.Error
		}
		return nil
	})
}

// 使用SQL扩展库进行查询
// 假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
// 要求 ：
// 编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
// 编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。

type Employee struct {
	Id         int     `db:"id"`
	Name       string  `db:"name"`
	Department string  `db:"department"`
	Salary     float64 `db:"salary"`
}

// 实现类型安全映射
// 假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
// 要求 ：
// 定义一个 Book 结构体，包含与 books 表对应的字段。
// 编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。

type Book struct {
	Id     int     `db:"id"`
	Title  string  `db:"title"`
	Author string  `db:"author"`
	Price  float64 `db:"price"`
}

func run3(db *sqlx.DB) error {
	var employees []Employee
	err := db.Select(&employees, "select * from employees where department = ?", "技术部")
	if err != nil {
		return err
	}
	for _, e := range employees {
		fmt.Printf("e: %v\n", e)
	}

	var employee Employee
	err = db.Get(&employee, "select * from employees order by salary desc limit 1")
	if err != nil {
		return err
	}
	fmt.Printf("employee: %v\n", employee)

	var findBooks []*Book
	err = db.Select(&findBooks, "select * from books where price > ？", 50.0)
	if err != nil {
		fmt.Println("查询查询价格大于 50 元的书籍，失败原因是:", err)
	}
	for _, b := range findBooks {
		fmt.Printf("b: %v\n", b)
	}

	return nil
}

// 题目1：模型定义
// 假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
// 要求 ：
// 使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
// 编写Go代码，使用Gorm创建这些模型对应的数据库表。

type User struct {
	Id        int
	Name      string `gorm:"uniqueIndex;size:50;not null"`
	Posts     []Post `gorm:"foreignKey:UserId"`
	PostCount int    `gorm:"default:0"`
}

type Post struct {
	Id            int
	Title         string `gorm:"size:200;not null"`
	UserId        int    `gorm:"index;not null"` // 外键
	CommentStatus string `gorm:"size:20;default:无评论"`
	CommentCount  int
	Comments      []Comment `gorm:"foreignKey:PostId"`
}

type Comment struct {
	Id      int
	Content string `gorm:"type:text;not null"`
	PostId  int    `gorm:"index;not null"` // 外键
	UserId  int    `gorm:"index;not null"` // 外键
}

func createTable(db *gorm.DB) {
	db.AutoMigrate(&User{}, &Post{}, &Comment{})

	db.Create([]User{{Name: "张三"}, {Name: "李四"}, {Name: "王五"}})

	db.Create([]Post{{Title: "webs初识", UserId: 1}, {Title: "gorm初识", UserId: 1}, {Title: "gin初识", UserId: 1}, {Title: "sqlx初识", UserId: 2}})

	db.Create([]Comment{{Content: "webs初识内容评论", PostId: 1, UserId: 2}, {Content: "欢迎大家评论", PostId: 1, UserId: 1},
		{Content: "不错，学到了", PostId: 2, UserId: 2}, {Content: `有问题，第3行哪里不对
		不是那样的，应该是。。。`, PostId: 4, UserId: 1}})
}

// 题目2：关联查询
// 基于上述博客系统的模型定义。
// 要求 ：
// 编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
// 编写Go代码，使用Gorm查询评论数量最多的文章信息。

func run4(userId int, db *gorm.DB) error {
	var users []User
	result := db.Preload("Posts").Preload("Posts.Comments").Where("id = ?", userId).Find(&users)
	if result.Error != nil {
		return result.Error
	}
	for _, user := range users {
		fmt.Printf("user: %v\n", user)
	}

	var post Post
	result = db.Order("comment_count desc").First(&post)
	if result.Error != nil {
		return result.Error
	}
	fmt.Printf("post: %v\n", post)

	var comment = &Comment{UserId: 2, PostId: 4}
	db.Where("user_id =? and post_id=?", 2, 4).Delete(&comment)
	return nil
}

// 题目3：钩子函数
// 继续使用博客系统的模型。
// 要求 ：
// 为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
// 为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。

func (c *Comment) AfterCreate(db *gorm.DB) error {
	var post Post
	result := db.Where("id =?", c.PostId).Find(&post)
	if result.Error != nil {
		return result.Error
	}
	post.CommentCount = post.CommentCount + 1
	post.CommentStatus = "有评论"
	result = db.Save(&post)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (p *Post) AfterCreate(db *gorm.DB) error {
	result := db.Model(&User{}).Where("id = ?", p.UserId).Update("post_count", gorm.Expr("post_count + 1"))
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (c *Comment) AfterDelete(db *gorm.DB) error {
	var commentCount int64
	result := db.Model(&Comment{}).Where("post_id = ?", c.PostId).Count(&commentCount)

	if result.Error != nil {
		return result.Error
	}

	if commentCount != 0 {
		return nil
	}

	result = db.Debug().Model(Post{}).Select("CommentCount,CommentStatus").Where("id = ?", c.PostId).Updates(Post{CommentStatus: "无评论", CommentCount: 0})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func main() {

	dsn := "root:123456@tcp(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// db, err := sqlx.Connect("mysql", dsn)

	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	// run1(db)

	// run2(db)
	// error := transfer(1, 2, db)

	// run3(db)

	// createTable(db)
	run4(1, db)

}
