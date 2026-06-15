package util

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

func TestRetryUtil(t *testing.T){
	testCases:=[]struct{
		name string
		times int
		steep time.Duration
		fn func() error
	}{
		{name: "count to 5 5 times",times: 5,steep: 5,fn:func() error {
			for i:=0;i<5;i++{
				fmt.Print(i)
			}
			return nil
		}},
		{name: "count to 5 5 times with returning error",times: 5,steep: 5,fn: func() error {
			for i:=0;i<5;i++{
				fmt.Print(i)
			}
			return errors.New("Successfully returnned an error")
		}},	
	}
	for _,tt:=range testCases{
		t.Run(tt.name,func(t *testing.T) {
			Retry(tt.times,tt.steep,tt.fn)
		})
	}
}
