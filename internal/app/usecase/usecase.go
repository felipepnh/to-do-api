package usecase

type UseCase interface {
	Handle(input interface{}) (output interface{}, err error)
}
