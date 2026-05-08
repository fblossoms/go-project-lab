package impl

import (
	"fmt"
	"go18/book/v4/apps/comment"

	"github.com/infraboard/mcube/v2/ioc"
)

func init() {
	ioc.Controller().Registry(&CommentServiceImpl{
		MaxCommentPerBook: 100, // 一本书最多支持100条评论
	})
}

// 如何知道有没有实现呢？通过类型约束
// var _ book.Service = &BookServiceImpl{}
// &BookServiceImpl 的 nil（空指针）的对象，减少内存开销
var _ comment.Service = (*CommentServiceImpl)(nil)

// CommentServiceImpl BookServiceImpl Book业务的具体实现
// 作为具体实现就会有很多种情况
// 可以添加多种实现：例如项目后续更新可能使用到多种技术栈，就会产生多种情况
type CommentServiceImpl struct {
	ioc.ObjectImpl

	// Comment 最大限制
	MaxCommentPerBook int `toml:"max_comment_per_book"`
}

func (s *CommentServiceImpl) Init() error {
	// 当前已经读取了配置文件
	fmt.Println(s.MaxCommentPerBook)
	return nil
}

func (s *CommentServiceImpl) Name() string {
	return comment.APP_NAME
}
