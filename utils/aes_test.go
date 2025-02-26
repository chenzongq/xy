package utils

import (
	"encoding/hex"
	"testing"
)

func TestAes(t *testing.T) {
	key := []byte("tcxcntcxcntcxcnx")
	encrypted := AesEncryptCBC([]byte("finance-a"), key)
	encodeToString := hex.EncodeToString(encrypted)
	t.Logf("密文(hex)：%s", encodeToString)
	decrypted := AesDecryptCBC(encrypted, key)
	t.Logf("解密结果：%s", string(decrypted))

	decodeString, _ := hex.DecodeString("0274974123a1ac801188c7f37c648af7")

	cbc := AesDecryptCBC(decodeString, key)
	t.Logf("解密 %s", string(cbc))
}
