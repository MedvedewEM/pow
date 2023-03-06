package components

type Challenger interface {
	TryGenerate() (string, string, error)
	TryVerify(id string, token string) error
}
