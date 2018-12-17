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

func Test_ErrorLevel(t *testing.T){
	e := errors.New("error")
	err := ErrorLevel(4,1, e)
	err1 := ErrorLevel(4,1, err)
	if !reflect.DeepEqual(err, err1) {
		t.Fatal(err1)
	}

}