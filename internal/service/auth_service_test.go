package service

import (
	"go-ecommerce-api/pkg/utils"
	"testing"
)

func TestHashPassword(t *testing.T) {

	password := "123456"

	hash, err := utils.HashPassword(password)

	if err != nil {

		t.Error("Hash Password Failed")
	}

	if hash == "" {

		t.Error("Empty")
	}
}
