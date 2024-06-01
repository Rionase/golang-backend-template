package contains

func Contains(text string, slice []string) bool {
	for _, item := range slice {
		if item == text {
			return true
		}
	}
	return false
}
