package hmac

import (
	goHmac "crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
)

func SignInterface(obj interface{}, secretKey []byte) (string, error) {
	jsoned, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}
	return Sign(jsoned, secretKey), nil
}

func Sign(msg []byte, secretKey []byte) string {
	mac := goHmac.New(sha256.New, secretKey)
	mac.Write(msg)

	return hex.EncodeToString(mac.Sum(nil))
}
