package impl_test

import (
	"go18/book/v4/apps/book"
	"testing"
)

func TestCreateBook(t *testing.T) {
	req := book.NewCreateBookRequest()
	req.SetIsSale(true)
	req.Title = "Go语言V4.0"
	req.Author = "will"
	req.Price = 99.9
	ins, err := svc.CreateBook(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func TestQueryBook(t *testing.T) {
	req := book.NewQueryBookeRequest()
	ins, err := svc.QueryBook(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}
