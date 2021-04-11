package dictionary

import "errors"

type Dictionary map[string]string

var ErrUnKnowWord = errors.New("该 key 不在字典中！")

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrUnKnowWord
	}

	return definition, nil
}

func (d Dictionary) Add(key, value string) {
	// d.Add(key, value)
	d[key] = value
}

// func Search(dic map[string]string, key string) string {
// 	return dic[key]
// }
