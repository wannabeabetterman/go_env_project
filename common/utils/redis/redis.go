package redis

import "strings"

const (
	DefaultDelimiter   = ":"
	RedisLockKeyPrefix = "_LOCK"
	RedisLockTime      = 10
)

// 组装key
func CreateKey(useToLock bool, delimiter string) func(names ...string) string {

	if delimiter == "" {
		delimiter = DefaultDelimiter
	}

	if useToLock {
		return func(names ...string) string {
			return RedisLockKeyPrefix + delimiter + strings.Join(names, delimiter)
		}
	}

	return func(names ...string) string {
		return strings.Join(names, delimiter)
	}
}
