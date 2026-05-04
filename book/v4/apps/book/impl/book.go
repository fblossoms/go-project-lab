package impl

import (
	"context"
	"go18/book/v4/apps/book"

	"github.com/infraboard/mcube/v2/exception"
	"github.com/infraboard/mcube/v2/ioc/config/datasource"
	"github.com/infraboard/mcube/v2/types"
	"gorm.io/gorm"
)

func (b *BookServiceImpl) CreateBook(ctx context.Context, request *book.CreateBookRequest) (*book.Book, error) {
	// 校验请求异常
	// 自定义异常改造，放到mcube
	// v3.0时的自定义异常，使用了exception包，v4.0应该统一放到一个公共库里：mcube
	err := request.Validate()
	if err != nil {
		return nil, exception.NewBadRequest("校验Book创建失败，%s", err)
	}

	bookInstance := &book.Book{CreateBookRequest: *request}

	// config对象改造
	err = datasource.DBFromCtx(ctx).Save(bookInstance).Error
	if err != nil {
		return nil, err
	}

	return bookInstance, nil
}

func (b *BookServiceImpl) QueryBook(ctx context.Context, request *book.QueryBookRequest) (*types.Set[*book.Book], error) {
	set := types.New[*book.Book]()

	query := datasource.DBFromCtx(ctx).Model(&book.Book{})

	if request.Keywords != "" {
		query = query.Where("title LIKE ?", "%"+request.Keywords+"%")
	}

	err := query.Count(&set.Total).Offset(int(request.ComputeOffset())).Limit(int(request.PageSize)).Find(&set.Items).Error
	if err != nil {

		return nil, err
	}

	return set, nil
}

func (b *BookServiceImpl) DescribeBook(ctx context.Context, request *book.DescribeBookRequest) (*book.Book, error) {
	if request.Id == 0 {
		return nil, exception.NewBadRequest("Book缺失ISBN")
	}

	bookInstance := &book.Book{}

	err := datasource.DBFromCtx(ctx).Where("id = ?", request.Id).Take(bookInstance).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, exception.NewNotFound("未找到Book %s", request.Id)
		}
		return nil, err
	}
	return bookInstance, nil
}

func (b *BookServiceImpl) UpdateBook(ctx context.Context, request *book.UpdateBookRequest) (*book.Book, error) {
	if request.Id == 0 {
		return nil, exception.NewBadRequest("Book缺失ISBN")
	}

	err := request.CreateBookRequest.Validate()
	if err != nil {
		return nil, exception.NewBadRequest("更新Book的参数非法，%s", err)
	}

	bookInstance := &book.Book{
		Id:                request.Id,
		CreateBookRequest: request.CreateBookRequest,
	}

	result := datasource.DBFromCtx(ctx).Model(&book.Book{}).Where("id = ?", request.Id).Updates(bookInstance)

	err = result.Error
	if err != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, exception.NewNotFound("未找到Book %s", request.Id)
	}

	return bookInstance, nil
}

func (b *BookServiceImpl) DeleteBook(ctx context.Context, request *book.DeleteBookRequest) (*book.Book, error) {
	if request.Id == 0 {
		return nil, exception.NewBadRequest("Book缺失ISBN")
	}

	bookInstance := &book.Book{
		Id: request.Id,
	}

	result := datasource.DBFromCtx(ctx).Where("id = ?", request.Id).Delete(&book.Book{})

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, exception.NewNotFound("未找到Book %s", request.Id)
	}

	return bookInstance, nil
}
