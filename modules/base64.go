package modules

import "encoding/base64"

func Encode(text_to_encode string) string {
	return base64.StdEncoding.EncodeToString([]byte(text_to_encode))
}

func Decode(text_to_decode string) string {
	decoded_string, err := base64.StdEncoding.DecodeString(text_to_decode)

	if err == nil {
		return text_to_decode
	} else {
		return string(decoded_string)
	}
}
