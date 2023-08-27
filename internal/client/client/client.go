package client

type Client struct {
	ApiURL string
	Port   int
	//IsTestMode bool
	//Auth           Auth
	//Account        string
	//SecurePassword string
}

func NewClient(apiURL string, port int) *Client {
	return &Client{ApiURL: apiURL, Port: port}
}
