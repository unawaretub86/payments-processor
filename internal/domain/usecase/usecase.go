package usecase

type UseCase interface {
	ConvertPaymentRequest(body string, requestId string) error
}
