package web

import (
	"testing"
)

func TestNewRouter(t *testing.T) {
	newRouter := NewRouter()
	if "passwords" != newRouter.GetRoute("passwords").GetName() {
		t.Errorf("Route name is not the same")
	}
	pathRegexp, _ := newRouter.GetRoute("passwords").GetPathRegexp()
	if "^/password/(?P<v0>[^/]+)/(?P<v1>[^/]+)/(?P<v2>[^/]+)/(?P<v3>[^/]+)[/]?$" != pathRegexp {
		t.Errorf("Route name is not the same")
	}

	methods, _ := newRouter.GetRoute("passwords").GetMethods()
	if contains(methods, "GET") == false {
		t.Errorf("Methods Error")
	}
}

func contains(arr []string, element string) bool {
	for _, a := range arr {
		if a == element {
			return true
		}
	}
	return false
}
