package httpclient

type (
	Module struct{}

	Request struct {
		ApiUrl      string
		QueryParams map[string]string
		Method      string
	}
)
