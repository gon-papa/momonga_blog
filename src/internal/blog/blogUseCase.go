package blog

import (
	"fmt"
	"momonga_blog/internal/types"
	"momonga_blog/repository"
	"momonga_blog/repository/model"
)



type BlogUseCaseInterface interface {
	GetBlogList(page types.Page, limit types.Limit) (*types.BlogList, error)
	GetBlog(uuid types.Uuid) (*model.Blog, error)
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

func (buc *blogUseCase) GetBlog(uuid types.Uuid) (*model.Blog, error) {
	blog, err := buc.repository.FindBlogByUUID(uuid)
	if err != nil {
		return nil, fmt.Errorf("failed to get blog: %w", err)
	}

	return blog, nil
}