package proxy

type LetsGoChat struct {
}

func (a *LetsGoChat) HandleRequest(url, method string) (int, string) {
	if url == "/user" && method == "POST" {
		return 201, "Create user endpoint"
	}

	if url == "/user/login" && method == "POST" {
		return 200, "User login endpoint"
	}

	return 404, "Endpoint not found"
}
