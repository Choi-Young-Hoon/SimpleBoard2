package resource

import "testing"

func TestNetworkInfo(test *testing.T) {
	networkInfo := NewNetworkInfo()

	err := networkInfo.GetNetworkInfo()
	if err != nil {
		test.Error(err)
	}

	networkInfo.Dump()
}
