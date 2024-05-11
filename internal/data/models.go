package data

type Models struct {
	Book BookInterface
}

func NewModels() *Models {
	return &Models{
		Book: NewBookModel(),
	}
}
