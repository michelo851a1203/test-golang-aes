package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAesRestoreAndConvert(t *testing.T) {
	assert := assert.New(t)
	input := "testingHello"
	mainKey := "secret12348khxeg"

	encryptedByte, err := EncryptAesToByte(input, mainKey)
	assert.Nil(err)
	restoreString, err := DecryptAesFromByte(encryptedByte, mainKey)
	assert.Nil(err)
	assert.Equal(input, restoreString)
}
