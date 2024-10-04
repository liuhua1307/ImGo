package domain

type Dynamic struct {
	ID          uint64
	AuthorId    uint64
	Title       string
	Info        string
	ImageUrl    []string
	CommentList []uint64
}

type DynamicUseCase interface {
	AsyncPushDynamic(info Dynamic) error
}

type DynamicRepository interface {
	PushDynamic(info Dynamic) error
}
