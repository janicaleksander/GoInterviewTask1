package types

const (
	FUNPROB = "Problem with function"
	EAPI    = "External api problem"
	API     = "Problem"
	UNS     = "Unsupported operation"
)

type ErrorAPI struct {
	Error string `json:"error"`
}
