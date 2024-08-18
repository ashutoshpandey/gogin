package utils

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ParseQueryParams parses an array of query string keys and returns their values, including errors if any.
func ParseQueryParams(c *gin.Context, keys []string) (map[string]int, error) {
	queryParams := make(map[string]int)

	for _, key := range keys {
		valueStr := c.Query(key)
		if valueStr == "" {
			return nil, errors.New("missing query parameter: " + key)
		}

		value, err := strconv.Atoi(valueStr)
		if err != nil {
			return nil, errors.New("invalid value for query parameter: " + key)
		}

		queryParams[key] = value
	}

	return queryParams, nil
}
