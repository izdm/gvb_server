package utils

// 判断key是否存在于列表中
func InList(key string, list []string) bool {
	for _, v := range list {
		if v == key {
			return true
		}
	}
	return false
}
