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
	CreateBlog(blog types.CreateBlogData, tagUuids []string) (*model.Blog, error)
	UpdateBlog(uuid types.Uuid, blog types.UpdateBlogData, tagUuids []string) (*model.Blog, error)
}

type blogUseCase struct {
	repository repository.BlogRepositoryInterface
	tagRepository repository.TagRepositoryInterface
}

var _ BlogUseCaseInterface = &blogUseCase{}

func NewBlogUseCase() BlogUseCaseInterface {
	return &blogUseCase{
		repository: repository.NewBlogRepository(),
		tagRepository: repository.NewTagRepository(),
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

func (buc *blogUseCase) CreateBlog(blog types.CreateBlogData, tagUuids []string) (*model.Blog, error) {
	// タグをUUIDから取得
	tags, err := buc.tagRepository.GetTagsByUuids(tagUuids)
	if err != nil {
		return nil, fmt.Errorf("failed to get tags: %w", err)
	}
	savedBlog, err := buc.repository.CreateBlog(blog, tags)
	if err != nil {
		return nil, fmt.Errorf("failed to create blog: %w", err)
	}

	return savedBlog, nil
}

func (buc *blogUseCase) UpdateBlog(uuid types.Uuid, updateData types.UpdateBlogData, tagUuids []string) (*model.Blog, error) {
	updatedBlog, err := buc.repository.UpdateBlog(uuid, updateData, tagUuids)
	if err != nil {
		return nil, fmt.Errorf("failed to update blog: %w", err)
	}

	return updatedBlog, nil
}