package util

// GetOptQuery returns def val if q is empty
func GetOptQuery(q string, def string) string {
	if q == "" {
		return def
	}
	return q
}
