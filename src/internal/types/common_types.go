package types

type Page int
func NewPage(p int) Page {
	if p < 1 {
		return Page(1)
	}
	return Page(p)
}

func (p Page) ToInt() int {
	return int(p)
}

type Limit int
func NewLimit(l int) Limit {
	if l < 1 {
		return Limit(10)
	}
	return Limit(l)
}

func (l Limit) ToInt() int {
	return int(l)
}