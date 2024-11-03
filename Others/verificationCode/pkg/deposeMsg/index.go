package deposeMsg

import (
	"encoding/json"
	"errors"
	"os"
	"verification/model"
)

func Get() (map[string]model.PhoneDetail, error) {
	file, err := os.OpenFile("msg.txt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return nil, errors.New(model.OpenFileError)
	}
	defer file.Close()
	// 检查文件的大小
	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}
	if fileInfo.Size() == 0 {
		// 如果文件是空的，就返回一个空的map和nil错误
		return map[string]model.PhoneDetail{}, nil
	}

	var data map[string]model.PhoneDetail
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func Save(p *map[string]model.PhoneDetail) error {
	var f *os.File
	file, err := os.OpenFile("msg.txt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return errors.New(model.OpenFileError)
	}
	defer f.Close()
	decoder := json.NewEncoder(file)
	err = decoder.Encode(p)
	if err != nil {
		return err
	}
	return nil
}
