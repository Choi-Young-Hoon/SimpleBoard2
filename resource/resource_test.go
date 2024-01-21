package resource

import "testing"

func TestSystemResource(t *testing.T) {
	systemResource := NewSystemResource()

	systemResource.GetSystemResource()
	systemResource.Dump()
}
