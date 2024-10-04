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

type CreateBlogData struct {
    Year     *int
    Month    *int
    Day      *int
    Title    string
    Body     string
    IsShow   bool
}

func NewCreateBlogData(year *int, month *int, day *int, title string, body string, isShow bool) CreateBlogData {
    return CreateBlogData{
        Year:   year,
        Month:  month,
        Day:    day,
        Title:  title,
        Body:   body,
        IsShow: isShow,
    }
}

type UpdateBlogData struct {
    UUID     string
    Title    string
    Body     string
    IsShow   bool
}

func NewUpdateBlogData(uuid string, title string, body string, isShow bool) UpdateBlogData {
    return UpdateBlogData{
        UUID:   uuid,
        Title:  title,
        Body:   body,
        IsShow: isShow,
    }
}