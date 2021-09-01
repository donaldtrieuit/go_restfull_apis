package utils

import "os"

/**
 * @author : Donald Trieu
 * @created : 9/1/21, Wednesday
**/

func GetEnvString(key string, def string) string {
	result, ok := os.LookupEnv(key)
	if ok {
		return result
	}

	return def
}