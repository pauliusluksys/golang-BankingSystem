package utils

import "database/sql"

//get keys from a map
func GetKeys(m map[string]string) []string {
	j := 0
	keys := make([]string, len(m))
	for k := range m {
		keys[j] = k
		j++
	}
	return keys
}
func CleanMap(filterMap map[string]string) map[string]string {
	var cleanMap = make(map[string]string)
	for index, key := range filterMap {
		if filterMap[index] != "" {
			cleanMap[index] = key
		}
	}
	return cleanMap
}
func TimeToString(time sql.NullTime) string {
	if time.Valid == true {
		return time.Time.String()
	} else {
		return ""
	}
}
