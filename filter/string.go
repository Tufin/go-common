package filter

import "strings"

func EndsWith(str string, suffix ...string) bool {

	ret := false
	for _, currSuffix := range suffix {
		if strings.HasSuffix(str, currSuffix) {
			ret = true
			break
		}
	}

	return ret
}
