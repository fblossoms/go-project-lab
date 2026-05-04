package impl

import (
	"go18/book/v4/apps/comment"
)

// 如何知道有没有实现呢？通过类型约束
// var _ book.Service = &BookServiceImpl{}
// &BookServiceImpl 的 nil（空指针）的对象，减少内存开销
var _ comment.Service = (*CommentServiceImpl)(nil)

// CommentServiceImpl BookServiceImpl Book业务的具体实现
// 作为具体实现就会有很多种情况
// 可以添加多种实现：例如项目后续更新可能使用到多种技术栈，就会产生多种情况
type CommentServiceImpl struct {
}
