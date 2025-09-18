package utils

import "net/url"

func GetValuePointer[T any](v T) *T {
	return &v
}

func GetPointerValue[T any](v *T) T {
	if v != nil {
		return *v
	} else {
		var zero T
		return zero
	}
}

func PathUnescape(raw string) string {
	res, err := url.PathUnescape(raw) // 非 ASCII 的字符要做额外处理
	if err != nil {
		return raw // 当出错时就返回原始的
	}
	return res
}

func RemoveDuplicate[T comparable](values []T) []T {
	results := make([]T, 0, len(values))
	seen := make(map[T]struct{}, len(values))
	for _, v := range values {
		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}
			results = append(results, v)
		}
	}
	return results
}
