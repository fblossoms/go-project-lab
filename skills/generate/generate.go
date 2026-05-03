package generate

type BookSet struct {
	Total int64    `json:"total"`
	Items []string `json:"items"`
}

func (b *BookSet) Add(item string) {
	b.Items = append(b.Items, item)
}

type CommentSet struct {
	Total int64 `json:"total"`
	Items []int `json:"items"`
}

func (b *CommentSet) Add(item int) {
	b.Items = append(b.Items, item)
}

func NewSet[T any]() *Set[T] {
	return &Set[T]{}
}

// Set 泛型结构体
// 使用[]来声明类型参数
type Set[T any] struct {
	Total int64 `json:"total"`

	Items []T `json:"items"`
}

func (b *Set[T]) Add(item T) {
	b.Items = append(b.Items, item)
}
