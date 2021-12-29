package common

import "encoding/json"

// SwapTo 通过json tag进行结构体赋予值
func SwapTo(request, category interface{}) (err error) {
	dataByte, err := json.Marshal(request)
	if err != nil {
		return
	}
	err = json.Unmarshal(dataByte, category)
	return err
}
