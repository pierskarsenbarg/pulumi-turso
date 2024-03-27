package pkg

import (
	"fmt"

	"github.com/pulumi/pulumi/sdk/go/common/resource"
)

func buildName(name string) (string, error) {
	return resource.NewUniqueHex(fmt.Sprintf("%s-", name), 8, 0)
}

func sliceCompare(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
