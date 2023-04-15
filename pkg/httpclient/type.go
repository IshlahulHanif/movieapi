package httpclient

type (
	Request struct {
		ApiUrl      string
		QueryParams map[string]string
		Method      string
	}
)
