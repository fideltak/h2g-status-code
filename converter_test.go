package converter

import (
	"testing"

	"google.golang.org/grpc/codes"
)

func TestHttpStatusToGrpcCode2XX(t *testing.T) {
	if HttpCodeToGrpcStaus(200) != codes.OK {
		t.Fatal("Error")
	}
	if HttpCodeToGrpcStaus(201) != codes.OK {
		t.Fatal("Error")
	}
	if HttpCodeToGrpcStaus(202) != codes.OK {
		t.Fatal("Error")
	}
	if HttpCodeToGrpcStaus(203) != codes.OK {
		t.Fatal("Error")
	}
	if HttpCodeToGrpcStaus(204) != codes.OK {
		t.Fatal("Error")
	}
	if HttpCodeToGrpcStaus(205) != codes.OK {
		t.Fatal("Error")
	}
	if HttpCodeToGrpcStaus(206) != codes.OK {
		t.Fatal("Error")
	}
	if HttpCodeToGrpcStaus(226) != codes.OK {
		t.Fatal("Error")
	}
	if HttpCodeToGrpcStaus(208) != codes.OK {
		t.Fatal("Error")
	}
}

func TestHttpStatusToGrpcCode4XX(t *testing.T) {
	if HttpCodeToGrpcStaus(400) != codes.InvalidArgument {
		t.Fatal("Error")
	}
	if HttpCodeToGrpcStaus(401) != codes.Unauthenticated {
		t.Fatal("Error")
	}
	if HttpCodeToGrpcStaus(403) != codes.PermissionDenied {
		t.Fatal("Error")
	}
	if HttpCodeToGrpcStaus(404) != codes.NotFound {
		t.Fatal("Error")
	}
	if HttpCodeToGrpcStaus(408) != codes.Canceled {
		t.Fatal("Error")
	}
	if HttpCodeToGrpcStaus(409) != codes.AlreadyExists {
		t.Fatal("Error")
	}
	if HttpCodeToGrpcStaus(429) != codes.ResourceExhausted {
		t.Fatal("Error")
	}
}

func TestHttpStatusToGrpcCode5XX(t *testing.T) {
	if HttpCodeToGrpcStaus(500) != codes.Internal {
		t.Fatal("Error")
	}
	if HttpCodeToGrpcStaus(501) != codes.Unimplemented {
		t.Fatal("Error")
	}
	if HttpCodeToGrpcStaus(503) != codes.Unavailable {
		t.Fatal("Error")
	}
	if HttpCodeToGrpcStaus(504) != codes.DeadlineExceeded {
		t.Fatal("Error")
	}
}

func TestHttpStatusToGrpcCodeOther(t *testing.T) {
	if HttpCodeToGrpcStaus(999) != codes.Unknown {
		t.Fatal("Error")
	}
}
