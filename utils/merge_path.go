package utils

func MergePaths(paths ...string) string {
	str := ""
	if len(paths) == 0 || paths[0] == "" {
		str = "./"
	}
	for i := range paths {
		if i != 0 {
			paths[i] = removeInitialSlash(paths[i])
			str = removeEndingSlash(str) + "/"
		}
		str += paths[i]
	}
	return str
}

func removeInitialSlash(s string) string {
	if len(s) > 0 && s[0] == '/' {
		return s[1:]
	}
	return s
}

func removeEndingSlash(s string) string {
	if len(s) > 0 && s[len(s)-1] == '/' {
		return s[:len(s)-1]
	}
	return s
}
