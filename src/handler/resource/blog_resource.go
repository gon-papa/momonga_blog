package resource

import (
	"momonga_blog/api"
	"momonga_blog/repository/model"
)

// ブログのリストをAPI用にマッピング
func MapBlogsToAPI(blogs []*model.Blog) []api.Blog {
    var blogList []api.Blog

    for _, blog := range blogs {
        blogList = append(blogList, api.Blog{
            UUID:      api.NewOptString(blog.UUID),
            Year:      api.NewOptInt(blog.Year),
            Month:     api.NewOptInt(blog.Month),
            Day:       api.NewOptInt(blog.Day),
            Title:     api.NewOptString(blog.Title),
            Body:      api.NewOptString(blog.Body),
            IsShow:    api.NewOptBool(blog.IsShow),
            CreatedAt: api.NewOptString(blog.CreatedAt.String()),
            UpdatedAt: api.NewOptString(blog.UpdatedAt.String()),
            DeletedAt: api.NewOptString(blog.DeletedAt.String()),
            Tags:      MapTagsToAPI(blog.Tags), // タグもマッピング
        })
    }

    return blogList
}

// タグのリストをAPI用にマッピング
func MapTagsToAPI(tags []model.Tag) []api.Tag {
    var apiTags []api.Tag
	if len(tags) == 0 {
		apiTags = []api.Tag{}
		return apiTags
	}

    for _, tag := range tags {
        apiTags = append(apiTags, api.Tag{
            UUID: api.NewOptString(tag.UUID),
            Name: api.NewOptString(tag.Name),
        })
    }

    return apiTags
}
