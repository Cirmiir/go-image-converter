package main

import (
	"encoding/base64"
	"encoding/hex"
	"errors"
)

var (
	m = make(map[string]Converter)
)

func init() {
	m["base64"] = Converter{Encode: func(data []byte) ([]byte, error) {
		return []byte(base64.StdEncoding.EncodeToString(data)), nil
	},
		Decode: func(data []byte) ([]byte, error) {
			return base64.StdEncoding.DecodeString(string(data))
		}}
	m["binary"] = Converter{Encode: func(data []byte) ([]byte, error) {
		return []byte(hex.EncodeToString(data)), nil
	},
		Decode: func(data []byte) ([]byte, error) {
			return hex.DecodeString(string(data))
		}}
}

/*Register add converter*/
func Register(converterName string, converter Converter) {
	m[converterName] = converter
}

/*Resolve return converted associated with converterName*/
func Resolve(converterName string) (Converter, error) {
	if conv, ok := m[converterName]; ok {
		return conv, nil
	}
	return Converter{}, errors.New("there is no converter associated with this name")
}

/*Unregister remove converter*/
func Unregister(converterName string) {
	delete(m, converterName)
}

/*UnregisterAll remove all converter*/
func UnregisterAll() {
	m = make(map[string]Converter)
}
