package lib

// 获取有效字段 GetValidFields
//
// fields    默认字段
//
// newFields 新字段
func GetValidFields(fields []string, newFields []string) []string {
	var validFiedls []string
	for _, v := range newFields {
		for _, v1 := range fields {
			if v == v1 {
				validFiedls = append(validFiedls, v1)
				break
			}
		}
	}
	return validFiedls
}
