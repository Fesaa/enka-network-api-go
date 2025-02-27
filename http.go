package enkanetworkapigo

import "net/http"

type HttpUserAgentInterceptor struct {
	api         EnkaNetworkAPI
	transporter http.RoundTripper
}

func (uai *HttpUserAgentInterceptor) RoundTrip(request *http.Request) (*http.Response, error) {
	request.Header.Set("User-Agent", uai.api.GetUserAgent())
	return uai.transporter.RoundTrip(request)
}
