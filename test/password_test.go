package test

import (
	"fanatic/lib/crypt"
	"fmt"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestCrypt(t *testing.T) {
	hashPsw:=crypt.EncodePassword("hello123")
	fmt.Println(string(hashPsw))
	assert.Equal(t,true,crypt.VerfifyPsw([]byte("hello123"),hashPsw))
}
