package responsecode

const (
	Unknown = iota
	OK
	NotFound
	DuplicatedFeedback
)

var messages map[int]string

func Message(code int) string {
	if messages != nil {
		return messages[code]
	}

	messages = map[int]string{
		0: "Unknown error",
		1: "OK",
		2: "Not found",
		3: "Duplicated feedback",
	}

	return messages[code]
}
