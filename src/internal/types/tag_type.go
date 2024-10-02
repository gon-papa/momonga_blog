package types

type CreateTagData struct {
	Name string
}
func NewCreateTagData(name string) CreateTagData {
	return CreateTagData{
		Name: name,
	}
}


func Map[T any, U any](input []T, mapper func(T) U) []U {
	var result []U
	for _, v := range input {
		result = append(result, mapper(v))
	}
	return result
}