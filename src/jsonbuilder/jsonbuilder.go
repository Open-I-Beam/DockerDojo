package jsonbuilder

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
)

func MarshalIntArray(inta []int) []byte {
	buffer := new(bytes.Buffer)
	e := gob.NewEncoder(buffer)
	err := e.Encode(inta)
	if err != nil {
		fmt.Println("Failed to marshal ", err)
	}
	return buffer.Bytes()
}

func UnmarshalIntArray(ba []byte) []int {
	buffer := bytes.NewBuffer(ba)
	var inta = new([]int)
	d := gob.NewDecoder(buffer)
	err := d.Decode(&inta)
	if err != nil {
		fmt.Println("Failed to unmarshal ", ba, err)
	}
	return *inta
}

func ReadData(r io.Reader, intf interface{}) error {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &intf)
	if err != nil {
		return err
	}
	return nil
}

func WriteData(w io.Writer, intf interface{}) error {
	b, err := json.Marshal(intf)
	fmt.Println(len(b), string(b))
	if err != nil {
		fmt.Println("error:", err)
	}

	i, err := w.Write(b)
	if i != len(b) || err != nil {
		fmt.Println("Nothing to write", i, err)
	}
	return err
}
