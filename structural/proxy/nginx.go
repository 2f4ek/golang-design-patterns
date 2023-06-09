package proxy

type Nginx struct {
	letsGoChat        *LetsGoChat
	maxAllowedRequest int
	rateLimiter       map[string]int
}

var maxAllowedRequests = 3

func NewNginxServer() *Nginx {
	return &Nginx{
		letsGoChat:        &LetsGoChat{},
		maxAllowedRequest: maxAllowedRequests,
		rateLimiter:       make(map[string]int),
	}
}

func (n *Nginx) HandleRequest(url, method string) (int, string) {
	allowed := n.checkRateLimiting(url)
	if !allowed {
		return 403, "Access denied"
	}

	return n.letsGoChat.HandleRequest(url, method)
}

func (n *Nginx) checkRateLimiting(url string) bool {
	if n.rateLimiter[url] == 0 {
		n.rateLimiter[url] = 1
	}
	if n.rateLimiter[url] > n.maxAllowedRequest {
		return false
	}
	n.rateLimiter[url] = n.rateLimiter[url] + 1
	return true
}
