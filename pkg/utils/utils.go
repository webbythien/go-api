package utils

import (
	"lido-core/v1/pkg/constants"
	"strings"
	"time"

	"github.com/google/uuid"
)

func TimeNow() string {
	return time.Now().UTC().Format(constants.TimeFormat)
}

func NewID() string {
	id := uuid.New().String()
	id = strings.ReplaceAll(id, "-", "")
	return id
}
