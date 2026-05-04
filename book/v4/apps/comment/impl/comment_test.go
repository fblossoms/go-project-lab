package impl_test

import (
	"go18/book/v4/apps/comment"
	"testing"
)

func TestAddComment(t *testing.T) {
	ins, err := svc.AddComment(ctx, &comment.AddCommentRequest{
		BookId:  10,
		Comment: "评论测试",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)

}
