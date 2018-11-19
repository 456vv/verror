package verror

import (
	"testing"
	"errors"
	"reflect"
)


func Test_Error(t *testing.T){
	e := errors.New("error")
	err := TrackError(e)
	if reflect.DeepEqual(e, err) {
		t.Fatal(err)
	}
//	t.Log(err)
}