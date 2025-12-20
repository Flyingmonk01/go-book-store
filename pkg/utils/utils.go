package utils

import "encoding/json"

func CheckNilError(err error) {
	if err != nil {
		panic(err)
	}
}

func ParseBody(body []byte, v interface{}) error {
	return json.Unmarshal(body, v)
}
func ToJson(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}
func MustToJson(v interface{}) []byte {
	data, err := json.Marshal(v)
	CheckNilError(err)
	return data
}
