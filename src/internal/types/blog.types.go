package types

import "momonga_blog/repository/model"

type BlogList struct {
    Blogs []*model.Blog
    Page  int
    Limit int
    Total int64
}

func NewBlogList(blogs []*model.Blog, page int, limit int, total int64) *BlogList {
    return &BlogList{
        Blogs: blogs,
        Page:  page,
        Limit: limit,
        Total: total,
    }
}