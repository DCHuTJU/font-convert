package convert

import (
	"encoding/json"
	"github.com/sbinet/go-python"
	"log"
)

var (
	simple2tradition *python.PyObject
	tradition2simple *python.PyObject
)

//  go build main.go && PYTHONPATH=./convert ./main

func init() {
	err := python.Initialize()
	if err != nil {
		log.Fatal(err)
	}
	fontConvert := python.PyImport_ImportModule("font-convert")
	if fontConvert == nil {
		panic("Error importing module")
	}

	simple2tradition = fontConvert.GetAttrString("simple2tradition")
	if simple2tradition == nil {
		panic("Error importing function")
	}

	tradition2simple = fontConvert.GetAttrString("tradition2simple")
	if tradition2simple == nil {
		panic("Error importing function")
	}
}

func PythonFinalize() {
	python.Finalize()
}
// src：source string
// condition: determine convert from Traditional to Simplified or Simplified to Traditional
func FontConvert(src, condition string) string {
	bArgs := python.PyTuple_New(1)
	err := python.PyTuple_SetItem(bArgs, 0, python.PyString_FromString(src))
	if err != nil {
		panic(err)
	}
	var res *python.PyObject
	if condition == "traditional" {
		res = simple2tradition.Call(bArgs, python.PyDict_New())
	} else if condition == "simplified" {
		res = tradition2simple.Call(bArgs, python.PyDict_New())
	} else {
		return src
	}

	return python.PyString_AS_STRING(res)
}


func Convert(v interface{}, font string) string {
	bytes,err := json.Marshal(v)
	if err != nil {
		log.Println(err)
		return ""
	}
	return FontConvert(string(bytes), font)
}