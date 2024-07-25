package customerror

import (
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"strconv"
)

// Result: http.StatusCode From net/http
func convertGRPCErrorCodeToHTTPStatusCode(e error) (int, *CustomError) {

	st, ok := status.FromError(e)
	if !ok {
		return http.StatusInternalServerError, NewCustomError(
			nil,
			"Error Parse StatusCode From gRPC Response",
			fmt.Sprintf("Error Parse StatusCode From gRPC Response: %v", e),
			strconv.Itoa(http.StatusInternalServerError),
		)
	}

	httpCode, exists := complianceWithGRPCStatusCodesAndHTTPStatusCodes[st.Code()]
	if !exists {
		httpCode = http.StatusInternalServerError
	}

	return httpCode, nil
}

var complianceWithGRPCStatusCodesAndHTTPStatusCodes = map[codes.Code]int{
	codes.Canceled:           http.StatusRequestTimeout,
	codes.Unknown:            http.StatusInternalServerError,
	codes.InvalidArgument:    http.StatusBadRequest,
	codes.DeadlineExceeded:   http.StatusGatewayTimeout,
	codes.NotFound:           http.StatusNotFound,
	codes.AlreadyExists:      http.StatusConflict,
	codes.PermissionDenied:   http.StatusForbidden,
	codes.ResourceExhausted:  http.StatusTooManyRequests,
	codes.FailedPrecondition: http.StatusPreconditionFailed,
	codes.Aborted:            http.StatusConflict,
	codes.OutOfRange:         http.StatusBadRequest,
	codes.Unimplemented:      http.StatusNotImplemented,
	codes.Internal:           http.StatusInternalServerError,
	codes.Unavailable:        http.StatusServiceUnavailable,
	codes.DataLoss:           http.StatusInternalServerError,
	codes.Unauthenticated:    http.StatusUnauthorized,
}
