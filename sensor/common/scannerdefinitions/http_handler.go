package scannerdefinitions

import (
	"io"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
	"github.com/stackrox/rox/pkg/clientconn"
	"github.com/stackrox/rox/pkg/httputil"
	"github.com/stackrox/rox/pkg/mtls"
	"github.com/stackrox/rox/pkg/set"
	"github.com/stackrox/rox/pkg/utils"
	"google.golang.org/grpc/codes"
)

const scannerDefsPath = "/api/extensions/scannerdefinitions"

var (
	headersToProxy = set.NewFrozenStringSet("If-Modified-Since", "Accept-Encoding")
)

// scannerDefinitionsHandler handles requests to retrieve scanner definitions
// from Central.
type scannerDefinitionsHandler struct {
	centralClient *http.Client
}

// NewDefinitionsHandler creates a new scanner definitions handler.
func NewDefinitionsHandler(centralEndpoint string) (http.Handler, error) {
	client, err := clientconn.NewHTTPClient(mtls.CentralSubject, centralEndpoint, 0)
	if err != nil {
		return nil, errors.Wrap(err, "instantiating central HTTP transport")
	}
	return &scannerDefinitionsHandler{
		centralClient: client,
	}, nil
}

func (h *scannerDefinitionsHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	// Validate request.
	if request.Method != http.MethodGet {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	// Prepare the Central's request, proxy relevant headers and all parameters.
	// No need to set Scheme nor Host, as the client will already do that for us.
	centralURL := url.URL{
		Path:     scannerDefsPath,
		RawQuery: request.URL.RawQuery,
	}
	centralRequest, err := http.NewRequestWithContext(
		request.Context(), http.MethodGet, centralURL.String(), nil)
	if err != nil {
		httputil.WriteGRPCStyleErrorf(writer, codes.Internal, "failed to create request: %v", err)
		return
	}
	// Proxy relevant headers.
	for _, headerName := range headersToProxy.AsSlice() {
		for _, value := range request.Header.Values(headerName) {
			centralRequest.Header.Add(headerName, value)
		}
	}
	// Do request, copy all response headers, and body.
	resp, err := h.centralClient.Do(centralRequest)
	if err != nil {
		httputil.WriteGRPCStyleErrorf(writer, codes.Internal, "failed to contact central: %v", err)
		return
	}
	defer utils.IgnoreError(resp.Body.Close)
	for k, vs := range resp.Header {
		for _, v := range vs {
			writer.Header().Add(k, v)
		}
	}
	writer.WriteHeader(resp.StatusCode)
	_, err = io.Copy(writer, resp.Body)
	if err != nil {
		httputil.WriteGRPCStyleErrorf(writer, codes.Internal, "failed write response: %v", err)
		return
	}
}
