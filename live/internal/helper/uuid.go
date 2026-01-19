package helper

import (
	"strings"

	"github.com/google/uuid"
)

// GenerateGUID
// 產生沒有-字元的UUID
func GenerateGUID() string {
	id := uuid.New()

	return strings.ReplaceAll(id.String(), "-", "")
}
