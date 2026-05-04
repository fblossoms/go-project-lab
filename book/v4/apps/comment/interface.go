package comment

import "context"

// GetService 让其返回符合Service接口的对象
// 严禁外部依赖内部同时内部依赖内部（死循环）
// 通过ioc容器来进行注册

// APP_NAME 当前以服务名称作为实现名称
const (
	APP_NAME = "comment"
)

// Service 评论的业务定义
type Service interface {
	// AddComment 为书籍添加评论
	AddComment(ctx context.Context, request *AddCommentRequest) (*Comment, error)
}

type AddCommentRequest struct {
	BookId  uint
	Comment string
}
