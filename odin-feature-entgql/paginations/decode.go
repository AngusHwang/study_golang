package paginations

import (
	"encoding/base64"
	"log"
)

func Decode(cursor string) (int, error) {
	decodedID, err := base64.StdEncoding.DecodeString(cursor)
	if err != nil {
		log.Println(err)
	}

	conversionTypeInt := int(decodedID[0])

	return conversionTypeInt, nil
}
