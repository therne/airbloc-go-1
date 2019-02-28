package protocol

import (
	"github.com/json-iterator/go"
	"net/url"
	"time"

	"github.com/airbloc/airbloc-go/data"

	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type HttpProtocol struct {
	timeout time.Duration
	client  *fasthttp.Client
}

func NewHttpProtocol(timeout time.Duration, maxConnsPerHost int) *HttpProtocol {
	client := &fasthttp.Client{
		ReadTimeout:     timeout,
		MaxConnsPerHost: maxConnsPerHost,
	}
	return &HttpProtocol{timeout, client}
}

func (http *HttpProtocol) Name() string {
	return "http"
}

func (http *HttpProtocol) Read(uri *url.URL) (bundle *data.Bundle, _ error) {
	request := fasthttp.AcquireRequest()
	request.SetRequestURI(uri.String())

	response := fasthttp.AcquireResponse()
	if err := http.client.DoTimeout(request, response, http.timeout); err != nil {
		return nil, errors.Wrap(err, "failed to estabilish HTTP connection")
	}

	if response.Header.StatusCode() == 404 {
		return nil, ErrNotFound
	}

	if err := json.Unmarshal(response.Body(), &bundle); err != nil {
		return nil, errors.Wrapf(err, "failed to parse bundle from the URL %s", uri.String())
	}

	return bundle, nil
}

type HttpsProtocol struct {
	HttpProtocol
}

func NewHttpsProtocol(timeout time.Duration, maxConnsPerHost int) *HttpsProtocol {
	return &HttpsProtocol{HttpProtocol: *NewHttpProtocol(timeout, maxConnsPerHost)}
}

func (https *HttpsProtocol) Name() string {
	return "https"
}
