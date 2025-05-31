package errors

type UserError interface {
	error
	UserMessage() string
	Status() int
}
