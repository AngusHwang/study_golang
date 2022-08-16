package paginations

import "encoding/base64"

func Encode(id int) string {
	encodedID := base64.StdEncoding.EncodeToString([]byte(string(rune(id))))

	return encodedID
}
