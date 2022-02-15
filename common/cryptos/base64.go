package cryptos

import "encoding/base64"

func EncodeBase64(src []byte) string {
	return base64.StdEncoding.EncodeToString(src)
}

func DecodeBase64(src string) (string, error) {
	dest, err := base64.StdEncoding.DecodeString(src)
	if err != nil {
		return "", err
	}
	return string(dest), nil
}
