package blog

import (
	"fmt"
	"momonga_blog/internal/types"
	"momonga_blog/repository"
)



type BlogUseCaseInterface interface {
	GetBlogList(page types.Page, limit types.Limit) (*types.BlogList, error)
}

type blogUseCase struct {
	repository repository.BlogRepositoryInterface
}

var _ BlogUseCaseInterface = &blogUseCase{}

func NewBlogUseCase() BlogUseCaseInterface {
	return &blogUseCase{
		repository: repository.NewBlogRepository(),
	}
}



func (buc *blogUseCase) GetBlogList(page types.Page, limit types.Limit) (*types.BlogList, error) {
	blogList, err := buc.repository.GetBlogs(page, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to get blog list: %w", err)
	}

	return blogList, nil
}