package resource

import "testing"

func TestDiskInfos(test *testing.T) {
	diskInfos := NewDiskInfos()
	if diskInfos == nil {
		test.Error("NewDiskInfos() returned nil")
	}

	err := diskInfos.GetDiskInfos()
	if err != nil {
		test.Error(err)
	}

	diskInfos.Dump()
}
