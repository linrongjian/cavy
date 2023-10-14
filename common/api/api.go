package api

import (
	uuid "github.com/satori/go.uuid"
)

// NewUUID 生成UUID
func NewUUID() uuid.UUID {
	//id, err := uuid.NewV4()
	//if err != nil {
	//	v1, _ := uuid.NewV1()
	//	return v1
	//}
	//
	//return id
	return uuid.NewV4()
}
