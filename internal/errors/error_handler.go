package errors

type GeneralError struct {
	Code             int
	Message          string
	DeveloperMessage string
}

func (e GeneralError) Error() string {
	return e.Message
}

func NewGeneralError(code int, message string, developerMessage string) GeneralError {
	return GeneralError{
		Code:             code,
		Message:          message,
		DeveloperMessage: developerMessage,
	}
}
