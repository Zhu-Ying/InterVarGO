package utils

import (
	"encoding/json"
	"log"
	"strconv"
)

func FlipACGT(acgt string) string {
	acgtMap := map[string]string{"A": "T", "T": "A", "G": "C", "C": "G", "N": "N", "X": "X"}
	if flip, ok := acgtMap[acgt]; ok {
		return flip
	} else {
		return acgt
	}
}

func StrToInt(from string) (to int) {
	to, err := strconv.Atoi(from)
	if err != nil {
		log.Panic(err)
	}
	return to
}

func StrToFloat64(from string) (to float64) {
	to, err := strconv.ParseFloat(from, 10)
	if err != nil {
		log.Panic(err)
	}
	return to
}

func FromJSON(from string, to interface{}) {
	err := json.Unmarshal([]byte(from), &to)
	if err != nil {
		log.Panic(err)
	}
}

func ToJSON(from interface{}) (to string) {
	dat, err := json.Marshal(from)
	if err != nil {
		log.Panic(err)
	}
	return string(dat)
}

func TransvertMap(data map[string]string) map[string]string {
	newData := make(map[string]string)
	for key, val := range data {
		newData[val] = key
	}
	return newData
}
