package resource

import "testing"

func TestMemoryInfo(test *testing.T) {
	memoryInfo := NewMemoryInfo()
	if memoryInfo == nil {
		test.Error("NewMemoryInfo() returned nil")
	}

	err := memoryInfo.GetMemoryInfo()
	if err != nil {
		test.Error(err)
	}

	memoryInfo.Dump()
}
