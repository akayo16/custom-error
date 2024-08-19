package customerror

import (
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"strconv"
)

// Result: http.StatusCode From net/http
func ConvertGRPCErrorToStatusStruct(e error) *status.Status {

	st, ok := status.FromError(e)
	if !ok {
		return nil
	}

	return st
}

func ConvertGRPCStatusCodeToHTTPStatusCode(e error) (int, *CustomError) {
	st, ok := status.FromError(e)
	if !ok {
		return http.StatusInternalServerError, NewCustomError(
			nil,
			"Error Parse StatusCode From gRPC Response",
			fmt.Sprintf("Error Parse StatusCode From gRPC Response: %v", e),
			strconv.Itoa(http.StatusInternalServerError),
			"custom-error.ConvertGRPCStatusCodeToHTTPStatusCode",
		)
	}

	httpCode, exists := complianceWithGRPCStatusCodesAndHTTPStatusCodes[st.Code()]
	if !exists {
		httpCode = http.StatusInternalServerError
	}

	return httpCode, nil
}

func ConvertHTTPStatusCodeToGRPCStatusCode(e error) (codes.Code, *CustomError) {
	st, ok := status.FromError(e)
	if !ok {
		return http.StatusInternalServerError, NewCustomError(
			nil,
			"Error Parse StatusCode From HTTP Response",
			fmt.Sprintf("Error Parse StatusCode From HTTP Response: %v", e),
			strconv.Itoa(http.StatusInternalServerError),
			"custom-error.ConvertHTTPStatusCodeToGRPCStatusCode",
		)
	}

	httpCode, exists := complianceWithHTTPtatusCodesAndGRPCStatusCodes[int(st.Code())]
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

var complianceWithHTTPtatusCodesAndGRPCStatusCodes = map[int]codes.Code{
	http.StatusContinue:                      codes.OK,
	http.StatusSwitchingProtocols:            codes.OK,
	http.StatusProcessing:                    codes.Unknown,
	http.StatusEarlyHints:                    codes.OK,
	http.StatusOK:                            codes.OK,
	http.StatusCreated:                       codes.OK,
	http.StatusAccepted:                      codes.OK,
	http.StatusNonAuthoritativeInfo:          codes.OK,
	http.StatusNoContent:                     codes.OK,
	http.StatusResetContent:                  codes.OK,
	http.StatusPartialContent:                codes.OK,
	http.StatusMultiStatus:                   codes.Unknown,
	http.StatusAlreadyReported:               codes.AlreadyExists,
	http.StatusIMUsed:                        codes.OK,
	http.StatusMultipleChoices:               codes.Unknown,
	http.StatusMovedPermanently:              codes.Unknown,
	http.StatusFound:                         codes.Unknown,
	http.StatusSeeOther:                      codes.Unknown,
	http.StatusNotModified:                   codes.Unknown,
	http.StatusUseProxy:                      codes.Unknown,
	http.StatusTemporaryRedirect:             codes.Unknown,
	http.StatusPermanentRedirect:             codes.Unknown,
	http.StatusBadRequest:                    codes.InvalidArgument,
	http.StatusUnauthorized:                  codes.Unauthenticated,
	http.StatusPaymentRequired:               codes.PermissionDenied,
	http.StatusForbidden:                     codes.PermissionDenied,
	http.StatusNotFound:                      codes.NotFound,
	http.StatusMethodNotAllowed:              codes.Unimplemented,
	http.StatusNotAcceptable:                 codes.InvalidArgument,
	http.StatusProxyAuthRequired:             codes.Unauthenticated,
	http.StatusRequestTimeout:                codes.DeadlineExceeded,
	http.StatusConflict:                      codes.Aborted,
	http.StatusGone:                          codes.NotFound,
	http.StatusLengthRequired:                codes.InvalidArgument,
	http.StatusPreconditionFailed:            codes.FailedPrecondition,
	http.StatusRequestEntityTooLarge:         codes.InvalidArgument,
	http.StatusRequestURITooLong:             codes.InvalidArgument,
	http.StatusUnsupportedMediaType:          codes.InvalidArgument,
	http.StatusRequestedRangeNotSatisfiable:  codes.OutOfRange,
	http.StatusExpectationFailed:             codes.FailedPrecondition,
	http.StatusTeapot:                        codes.Unknown,
	http.StatusMisdirectedRequest:            codes.FailedPrecondition,
	http.StatusUnprocessableEntity:           codes.InvalidArgument,
	http.StatusLocked:                        codes.FailedPrecondition,
	http.StatusFailedDependency:              codes.FailedPrecondition,
	http.StatusTooEarly:                      codes.FailedPrecondition,
	http.StatusUpgradeRequired:               codes.Unimplemented,
	http.StatusPreconditionRequired:          codes.FailedPrecondition,
	http.StatusTooManyRequests:               codes.ResourceExhausted,
	http.StatusRequestHeaderFieldsTooLarge:   codes.InvalidArgument,
	http.StatusUnavailableForLegalReasons:    codes.PermissionDenied,
	http.StatusInternalServerError:           codes.Internal,
	http.StatusNotImplemented:                codes.Unimplemented,
	http.StatusBadGateway:                    codes.Unavailable,
	http.StatusServiceUnavailable:            codes.Unavailable,
	http.StatusGatewayTimeout:                codes.DeadlineExceeded,
	http.StatusHTTPVersionNotSupported:       codes.Unimplemented,
	http.StatusVariantAlsoNegotiates:         codes.Unknown,
	http.StatusInsufficientStorage:           codes.ResourceExhausted,
	http.StatusLoopDetected:                  codes.Internal,
	http.StatusNotExtended:                   codes.Unknown,
	http.StatusNetworkAuthenticationRequired: codes.Unauthenticated,
}
