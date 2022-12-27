package singleton_test

import (
	"design-pattern/creational/singleton"
	"testing"
)

func TestGetSomeData(t *testing.T) {
	singleton.GetSomeData()
	singleton.GetSomeData2()
}
