package request

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

import (
	"github.com/abrie/censusgeocoder/internal/service"
)

type Benchmarks struct {
}

func (params Benchmarks) BuildHttpRequest(ctx context.Context, service *service.Service) (*http.Request, error) {
	url := strings.Join([]string{service.Endpoint, "benchmarks"}, "/")

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Failed to build HTTP Request: %v", err)
	}

	q := req.URL.Query()
	q.Add("format", FormatJSON)

	req.URL.RawQuery = q.Encode()

	return req, nil
}
