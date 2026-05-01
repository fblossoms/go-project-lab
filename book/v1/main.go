package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type BookSet struct {
	// 统计总数
	Total int64 `json:"total"`
	// book清单
	Items []*Book `json:"items"`
}
type Book struct {
	// 主键定义
	Id uint `json:"id" gorm:"primaryKey;column:id"`

	BookSpec
}

type BookSpec struct {
	Title  string  `json:"title" gorm:"column:title;type:varchar(200)" validate:"required"`
	Author string  `json:"author" gorm:"column:author;type:varchar(200);index" validate:"required"`
	Price  float64 `json:"price" gorm:"column:price" validate:"required"`
	IsSale *bool   `json:"is_sale" gorm:"column:is_sale"`
}

// 映射到books表
func (b *Book) TableName() string {
	return "books"
}

func SetupDatabase() *gorm.DB {
	dsn := "root:@tcp(localhost:3306)/go18?parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败")
	}
	db.AutoMigrate(&Book{})
	return db
}

var db = SetupDatabase()

func main() {
	server := gin.Default()

	// Book Restful API
	// List of Books
	server.GET("/api/books", func(ctx *gin.Context) {
		set := &BookSet{}
		// 常见：/api/books?page_number=1&page_size=20
		// 实现前后端分页
		pn, ps := 1, 20
		pageNumber := ctx.Query("page_number")
		if pageNumber != "" {
			pnInt, err := strconv.ParseInt(pageNumber, 10, 64) // 10进制64位
			if err != nil {
				ctx.JSON(400, gin.H{"code": 400, "message": err.Error()}) // 400表示请求不合法并把报错直接丢给用户，但是这样写不太友好
				return
			}
			pn = int(pnInt)
		}

		pageSize := ctx.Query("page_size")
		if pageSize != "" {
			psInt, err := strconv.ParseInt(pageSize, 10, 64)
			if err != nil {
				ctx.JSON(400, gin.H{"code": 400, "message": err.Error()}) // 400表示请求不合法并把报错直接丢给用户，但是这样写不太友好
				return
			}
			ps = int(psInt)
		}

		query := db.Model(&Book{})
		// 关键词过滤查询
		kws := ctx.Query("keywords")
		if kws != "" {
			// WHERE title LIKE %kws%
			query = query.Where("title Like ?", "%"+kws+"%x") // ? 避免恶意注入
		}

		//bookList := []Book{}
		// SELECT * FROM books
		// 通过sql的OFFSET LIMIT来实现分页：OFFSET (page_number -1) * page_size, LIMIT page_size
		// 2	OFFSET 20, 20
		// 3	OFFSET 2 * 20, 20
		// 4	OFFSET 3 * 60, 20
		offset := (pn - 1) * ps
		//err = db.Offset(int(offset)).Find(&bookList)
		err := query.Count(&set.Total).Offset(int(offset)).Limit(int(ps)).Find(&set.Items).Error
		if err != nil {
			ctx.JSON(500, gin.H{"code": 500, "message": err.Error()}) // 400表示请求不合法并把报错直接丢给用户，但是这样写不太友好
			return
		}

		// 获取总数（总共多少页）
		ctx.JSON(200, set)
	})

	// Create new Book
	// 常放在Body: HTTP Entity
	server.POST("/api/books", func(ctx *gin.Context) {
		//payload, err := io.ReadAll(ctx.Request.Body) // 绕开框架直接处理HTTP的Request对象，但是记得使用defer关闭
		//if err != nil {                              // 若有异常，返回给用户
		//	ctx.JSON(400, gin.H{"code": 400, "message": err.Error()}) // 400表示请求不合法并把报错直接丢给用户，但是这样写不太友好
		//	return
		//}
		//defer ctx.Request.Body.Close()
		// 假设{"title": "安徒生童话"}

		bookSPecInstance := &BookSpec{}
		//// 通过json的结构体标签
		//// bookInstance.Title = "安徒生童话"
		//err = json.Unmarshal(payload, bookInstance)
		//if err != nil {
		//	ctx.JSON(400, gin.H{"code": 400, "message": err.Error()}) // 400表示请求不合法并把报错直接丢给用户，但是这样写不太友好
		//	return
		//}

		// 获取到bookInstance实例，等价于上方冗余版本的代码
		err := ctx.BindJSON(bookSPecInstance) // 注意传的是指针，修改指针才能操作数据，修改值可能改不了要改的数据
		if err != nil {
			ctx.JSON(400, gin.H{"code": 500, "message": err.Error()}) // 400表示请求不合法并把报错直接丢给用户，但是这样写不太友好
			return
		}
		// 参数是否为空
		//if bookSPecInstance.Author == "" {
		//	ctx.JSON(400, gin.H{"code": 500, "message": err.Error()}) // 400表示请求不合法并把报错直接丢给用户，但是这样写不太友好
		//	return
		//}
		// 有没有能够检查某个字段是否必须填
		// Gin 集成 validator这个库，通过struct tag validate 来表示这个字段是否允许为空：validate:"required"表示不允许为空

		bookInstance := &Book{
			BookSpec: *bookSPecInstance,
		}
		// 数据入库（待完善省略，会用到ORM）
		err = db.Save(bookInstance).Error
		if err != nil {
			ctx.JSON(400, gin.H{"code": 400, "message": err.Error()})
			return
		}

		// 返回响应
		ctx.JSON(200, bookSPecInstance)
	})

	// Get book by ISBN
	server.GET("/api/books/:bn", func(ctx *gin.Context) { // bn就是URI里的路径变量
		// URI的获取
		bnStr := ctx.Param("bn")
		bn, err := strconv.ParseInt(bnStr, 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{"code": 400, "message": err.Error()}) // 400表示请求不合法并把报错直接丢给用户，但是这样写不太友好
			return
		}

		bookInstance := &Book{}
		// 从数据库中获取一个对象
		err = db.Where("id = ?", bn).Take(bookInstance).Error
		if err != nil {
			ctx.JSON(400, gin.H{"code": 400, "message": err.Error()}) // 400表示请求不合法并把报错直接丢给用户，但是这样写不太友好
			return
		}

		ctx.JSON(200, bookInstance)
	})

	// Update book
	server.PUT("/api/books/:bn", func(ctx *gin.Context) {
		bnStr := ctx.Param("bn")
		bn, err := strconv.ParseInt(bnStr, 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{"code": 400, "message": err.Error()}) // 400表示请求不合法并把报错直接丢给用户，但是这样写不太友好
			return
		}

		fmt.Println(bn)

		// 读取body里的参数
		bookInstance := &Book{
			Id: uint(bn),
		}

		err = ctx.BindJSON(bookInstance) // 注意传的是指针，修改指针才能操作数据，修改值可能改不了要改的数据
		if err != nil {
			ctx.JSON(400, gin.H{"code": 400, "message": err.Error()}) // 400表示请求不合法并把报错直接丢给用户，但是这样写不太友好
			return
		}

		err = db.Where("id = ?", ctx.Param("bn")).Updates(bookInstance).Error
		if err != nil {
			ctx.JSON(400, gin.H{"code": 400, "message": err.Error()}) // 400表示请求不合法并把报错直接丢给用户，但是这样写不太友好
			return
		}
	})

	// Delete book
	server.DELETE("api/books/:bn", func(ctx *gin.Context) {
		bnStr := ctx.Param("bn")
		bn, err := strconv.ParseInt(bnStr, 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{"code": 400, "message": err.Error()}) // 400表示请求不合法并把报错直接丢给用户，但是这样写不太友好
			return
		}

		err = db.Where("id = ?", ctx.Param("bn")).Delete(&Book{}).Error // 避免删除全表
		if err != nil {
			ctx.JSON(400, gin.H{"code": 400, "message": err.Error()}) // 400表示请求不合法并把报错直接丢给用户，但是这样写不太友好
			return
		}
		ctx.JSON(http.StatusNoContent, "ok")
		fmt.Println(bn)
	})

	err := server.Run("localhost:8080")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
