package controllers_test

import (
	"context"
	"go18/book/v3/config"
	"go18/book/v3/controllers"
	"go18/book/v3/models"
	"testing"
)

func TestGetBook(t *testing.T) {
	book, err := controllers.Book.GetBook(context.Background(), controllers.NewGetBookRequest("1"))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(book)

}

func TestCreateBook(t *testing.T) {
	book, err := controllers.Book.CreateBook(context.Background(), &models.BookSpec{
		Title:  "unit test for go controller obj",
		Author: "will",
		Price:  99.99,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(book)
}

func init() {
	// 执行配置的加载
	err := config.LoadConfigFromYaml("C:/Users/flyfl/Desktop/go_18/book/v3/application.yaml")
	if err != nil {
		panic(err)
	}
}
