package resource

import "momonga_blog/api"

func MapPaginationToAPI(total int64, page int, limit int) api.Pagenation {
	return api.Pagenation{
		Total: total,
		Page:  page,
		Limit: limit,
	}
}