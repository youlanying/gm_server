package tools

import "strconv"

// ListsMember 查看int32值是否存在于list中
func ListsMember(elem int32, list []int32) bool {
	for _, v := range list {
		if elem == v {
			return true
		}
	}
	return false
}

// ListsDelete 删除列表中指定元素
func ListsDelete(elem int32, list []int32) []int32 {
	newList := make([]int32, 0)
	for _, v := range list {
		if elem != v {
			newList = append(newList, v)
		}
	}
	return newList
}

// ListStringToListInt32 []string转[]int32
func ListStringToListInt32(values []string) []int32 {
	groupIdList := make([]int32, 0)
	for i := 0; i < len(values); i++ {
		intK, err := strconv.Atoi(values[i])
		if err == nil {
			groupIdList = append(groupIdList, int32(intK))
		}
	}
	return groupIdList
}
