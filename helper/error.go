package helper

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func ErrorRequestMessage(err error) string {
	var message = "Error : " + err.Error()
	return message
}

func ErrorDuplicateMessage(err error) string {
	return "record already exists"
}

func ErrorForeignMessage(err error) string {
	return "A foreign key constraint fails"
}
