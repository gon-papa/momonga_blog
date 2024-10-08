package types

var ImageExtension []string = []string{
	"jpg",
	"jpeg",
	"png",
	"gif",
}

func IsAllowedExtension(extension string) bool {
	for _, ext := range ImageExtension {
		if ext == extension {
			return true
		}
	}
	return false
}
