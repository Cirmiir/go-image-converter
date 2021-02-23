package main

import (
	"io/ioutil"
	"testing"
)

func TestDefaultConvertResolve(t *testing.T) {
	list := [2]string{"binary", "base64"}

	for _, name := range list {
		_, err := Resolve(name)
		if err != nil {
			t.Errorf("Default converter %v is not supported.", name)
		}
	}
}

func TestUnexistedConverterResolve(t *testing.T) {
	list := [2]string{"test"}

	for _, name := range list {
		_, err := Resolve(name)
		if err == nil {
			t.Errorf("Unexisted converter %v should raise error.", name)
		}
	}
}

func TestSimpleImageConvert(t *testing.T) {
	data, err := ioutil.ReadFile("Test.jpg")
	if err != nil {
		t.Error("Unexisted converter should raise error.")
	}

	converter, _ := Resolve("binary")
	encodedResult, _ := converter.Convert(data, true)
	decodedResult, _ := converter.Convert(encodedResult, false)

	if len(data) == len(decodedResult) {
		for i := range data {
			if data[i] != decodedResult[i] {
				t.Errorf("The result of decode is different than original data at index %v.", i)
			}
		}
	} else {
		t.Error("The result of decode has the different size than original data.")
	}
}

func ErrorForWrongSymbolInEncodedString(convertName string, t *testing.T) {
	converter, _ := Resolve(convertName)
	_, err := converter.Convert([]byte("!!"), false)

	if err == nil {
		t.Error("Convert should be failed if there are symbols that can not be converted.")
	}

}

func TestErrorForWrongSymbolInEncodedStringBinary(t *testing.T) {
	ErrorForWrongSymbolInEncodedString("binary", t)
}

func TestErrorForWrongSymbolInEncodedStringBase64(t *testing.T) {
	ErrorForWrongSymbolInEncodedString("base64", t)
}
