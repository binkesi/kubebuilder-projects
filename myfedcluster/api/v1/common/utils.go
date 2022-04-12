package common

import (
	"fmt"
	"strings"
)

func PropagatedVersionName(kind, resourceName string) string {
	return fmt.Sprintf("%s%s", PropagatedVersionPrefix(kind), resourceName)
}

func PropagatedVersionPrefix(kind string) string {
	return fmt.Sprintf("%s-", strings.ToLower(kind))
}
