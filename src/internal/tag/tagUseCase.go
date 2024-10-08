package tag

import (
	"fmt"
	"momonga_blog/internal/types"
	"momonga_blog/repository"
	"momonga_blog/repository/model"
)

type TagUseCaseInterface interface {
	GetTagList(blogUuid *types.Uuid) ([]*model.Tag, error)
	CreateTag(tagName string) error
}

type tagUseCase struct {
	repository repository.TagRepositoryInterface
	blogRepository repository.BlogRepositoryInterface
}

var _ TagUseCaseInterface = &tagUseCase{}

func NewTagUseCase() TagUseCaseInterface {
	return &tagUseCase{
		repository: repository.NewTagRepository(),
		blogRepository: repository.NewBlogRepository(),
	}
}

func (tuc *tagUseCase) CreateTag(tagName string) error {
	exsistTag, _ := tuc.repository.GetByName(tagName)
	if exsistTag != nil {
		return nil
	}

	err := tuc.repository.CreateTag(tagName)
	if err != nil {
		return fmt.Errorf("failed to create tag: %w", err)
	}
	return nil
}

func (tuc *tagUseCase) GetTagList(blogUuid *types.Uuid) ([]*model.Tag, error) {
	if blogUuid == nil {
		// blogUuidがnilの場合は全てのタグを取得
		tags, err := tuc.repository.GetTags()
		if err != nil {
			return nil, fmt.Errorf("failed to get tags: %w", err)
		}
		return tags, nil
	}

	blog, err := tuc.blogRepository.FindBlogByUUID(*blogUuid)
	if err != nil {
		return nil, fmt.Errorf("failed to get blog: %w", err)
	}

	tags, err := tuc.repository.GetTagByBlogUuid(blog.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get tags: %w", err)
	}
	
	return tags, nil
}