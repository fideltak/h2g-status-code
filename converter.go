package converter

import (
	"net/http"

	"google.golang.org/grpc/codes"
)

//https://grpc.github.io/grpc/core/md_doc_http-grpc-status-mapping.html
//https://github.com/grpc-ecosystem/grpc-gateway/blob/master/runtime/errors.go

func HttpCodeToGrpcStaus(httpStatusCode int) codes.Code {
	//code := http.StatusText(httpStatusCode)
	switch httpStatusCode {
	case http.StatusOK: //200
		return codes.OK //0

	case http.StatusCreated: //201
		return codes.OK //0

	case http.StatusAccepted: //202
		return codes.OK //0

	case http.StatusNonAuthoritativeInfo: //203
		return codes.OK //0

	case http.StatusNoContent: //204
		return codes.OK //0

	case http.StatusResetContent: //205
		return codes.OK //0

	case http.StatusPartialContent: //206
		return codes.OK //0

	case http.StatusMultiStatus: //207
		return codes.OK //0

	case http.StatusAlreadyReported: //208
		return codes.OK //0

	case http.StatusIMUsed: //226
		return codes.OK //0

	//No concept for 3XX in GRPC

	case http.StatusBadRequest: //400
		return codes.InvalidArgument //3

	case http.StatusUnauthorized: //401
		return codes.Unauthenticated //16

	case http.StatusForbidden: //403
		return codes.PermissionDenied //7

	case http.StatusNotFound: //404
		return codes.NotFound //5

	case http.StatusRequestTimeout: //408
		return codes.Canceled //1

	case http.StatusConflict: //409
		return codes.AlreadyExists //6

	case http.StatusTooManyRequests: //429
		return codes.ResourceExhausted //8



	case http.StatusInternalServerError: //500
		return codes.Internal //13

	case http.StatusNotImplemented: //501
		return codes.Unimplemented //12

	case http.StatusServiceUnavailable: //503
		return codes.Unavailable //14

	case http.StatusGatewayTimeout: //504
		return codes.DeadlineExceeded //4
	}

	return codes.Unknown //2
}
