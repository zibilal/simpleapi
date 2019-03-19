package logger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"testing"
)

const (
	success = "\u2713"
	failed  = "\u2717"
)

func TestLogger(t *testing.T) {
	t.Log("Testing Info")
	{
		var buff1 = bytes.NewBufferString("")
		SetOutput(buff1)
		Info("info1", "info2", "info3")
		data1 := struct {
			Level   string `json:"level"`
			Message string `json:"msg"`
			On      string `json:"on"`
			Time    string `json:"time"`
			Input0  string `json:"input-0"`
			Input1  string `json:"input-1"`
			Input2  string `json:"input-2"`
		}{}
		err := json.Unmarshal(buff1.Bytes(), &data1)
		if err != nil {
			t.Fatalf("%s expected error nil, got %s", failed, err.Error())
		} else {
			t.Logf("%s expected error nil", success)
		}

		if CheckEmpty(data1) {
			t.Fatalf("%s expected not empty, got empty data1", failed)
		} else {
			t.Logf("%s expected not empty", success)
		}

		var info1 = "info1"
		if data1.Input0 == info1 {
			t.Logf("%s expected Input1 == %s", success, info1)
		} else {
			t.Fatalf("%s expected Input1 == %s, got %s", failed, info1, data1.Input0)
		}

		var info2 = "info2"
		if data1.Input1 == info2 {
			t.Logf("%s expected Input2 == %s", success, info2)
		} else {
			t.Fatalf("%s expected Input2 == %s, got %s", failed, info2, data1.Input1)
		}

		var info3 = "info3"
		if data1.Input2 == info3 {
			t.Logf("%s expected Input3 == %s", success, info3)
		} else {
			t.Fatalf("%s expected Input3 == %s, got %s", failed, info3, data1.Input2)
		}

		var message = "info1info2info3"
		if data1.Message == message {
			t.Logf("%s expected Message == %s", success, message)
		} else {
			t.Fatalf("%s expected Message == %s, got %s", failed, message, data1.Message)
		}

		if CheckEmpty(data1.On) {
			t.Fatalf("%s expected on not empty, got empty", failed)
		} else {
			t.Logf("%s expected on not empty", success)
		}

		var level = "info"
		if data1.Level == level {
			t.Logf("%s expected Level == %s", success, level)
		} else {
			t.Fatalf("%s expected Level == %s, got %s", failed, level, data1.Level)
		}

		if CheckEmpty(data1.Time) {
			t.Fatalf("%s expected time not empty, got empty", failed)
		} else {
			t.Logf("%s expected time not empty", success)
		}
	}

	t.Log("Testing Warn")
	{
		var buff1 = bytes.NewBufferString("")
		SetOutput(buff1)
		Warn("warn1", "warn2", "warn3")
		data1 := struct {
			Level   string `json:"level"`
			Message string `json:"msg"`
			On      string `json:"on"`
			Time    string `json:"time"`
			Input0  string `json:"input-0"`
			Input1  string `json:"input-1"`
			Input2  string `json:"input-2"`
		}{}
		err := json.Unmarshal(buff1.Bytes(), &data1)
		if err != nil {
			t.Fatalf("%s expected error nil, got %s", failed, err.Error())
		} else {
			t.Logf("%s expected error nil", success)
		}

		if CheckEmpty(data1) {
			t.Fatalf("%s expected not empty, got empty data1", failed)
		} else {
			t.Logf("%s expected not empty", success)
		}

		var info1 = "warn1"
		if data1.Input0 == info1 {
			t.Logf("%s expected Input1 == %s", success, info1)
		} else {
			t.Fatalf("%s expected Input1 == %s, got %s", failed, info1, data1.Input0)
		}

		var info2 = "warn2"
		if data1.Input1 == info2 {
			t.Logf("%s expected Input2 == %s", success, info2)
		} else {
			t.Fatalf("%s expected Input2 == %s, got %s", failed, info2, data1.Input1)
		}

		var info3 = "warn3"
		if data1.Input2 == info3 {
			t.Logf("%s expected Input3 == %s", success, info3)
		} else {
			t.Fatalf("%s expected Input3 == %s, got %s", failed, info3, data1.Input2)
		}

		var message = "warn1warn2warn3"
		if data1.Message == message {
			t.Logf("%s expected Message == %s", success, message)
		} else {
			t.Fatalf("%s expected Message == %s, got %s", failed, message, data1.Message)
		}

		if CheckEmpty(data1.On) {
			t.Fatalf("%s expected on not empty, got empty", failed)
		} else {
			t.Logf("%s expected on not empty", success)
		}

		var level = "warning"
		if data1.Level == level {
			t.Logf("%s expected Level == %s", success, level)
		} else {
			t.Fatalf("%s expected Level == %s, got %s", failed, level, data1.Level)
		}

		if CheckEmpty(data1.Time) {
			t.Fatalf("%s expected time not empty, got empty", failed)
		} else {
			t.Logf("%s expected time not empty", success)
		}
	}

	t.Log("Testing Warn")
	{
		var buff1 = bytes.NewBufferString("")
		SetOutput(buff1)
		SetLevel("debug")
		Debug("debug1", "debug2", "debug3")
		data1 := struct {
			Level   string `json:"level"`
			Message string `json:"msg"`
			On      string `json:"on"`
			Time    string `json:"time"`
			Input0  string `json:"input-0"`
			Input1  string `json:"input-1"`
			Input2  string `json:"input-2"`
		}{}
		err := json.Unmarshal(buff1.Bytes(), &data1)
		if err != nil {
			t.Fatalf("%s expected error nil, got %s", failed, err.Error())
		} else {
			t.Logf("%s expected error nil", success)
		}

		if CheckEmpty(data1) {
			t.Fatalf("%s expected not empty, got empty data1", failed)
		} else {
			t.Logf("%s expected not empty", success)
		}

		var info1 = "debug1"
		if data1.Input0 == info1 {
			t.Logf("%s expected Input1 == %s", success, info1)
		} else {
			t.Fatalf("%s expected Input1 == %s, got %s", failed, info1, data1.Input0)
		}

		var info2 = "debug2"
		if data1.Input1 == info2 {
			t.Logf("%s expected Input2 == %s", success, info2)
		} else {
			t.Fatalf("%s expected Input2 == %s, got %s", failed, info2, data1.Input1)
		}

		var info3 = "debug3"
		if data1.Input2 == info3 {
			t.Logf("%s expected Input3 == %s", success, info3)
		} else {
			t.Fatalf("%s expected Input3 == %s, got %s", failed, info3, data1.Input2)
		}

		var message = "debug1debug2debug3"
		if data1.Message == message {
			t.Logf("%s expected Message == %s", success, message)
		} else {
			t.Fatalf("%s expected Message == %s, got %s", failed, message, data1.Message)
		}

		if CheckEmpty(data1.On) {
			t.Fatalf("%s expected on not empty, got empty", failed)
		} else {
			t.Logf("%s expected on not empty", success)
		}

		var level = "debug"
		if data1.Level == level {
			t.Logf("%s expected Level == %s", success, level)
		} else {
			t.Fatalf("%s expected Level == %s, got %s", failed, level, data1.Level)
		}

		if CheckEmpty(data1.Time) {
			t.Fatalf("%s expected time not empty, got empty", failed)
		} else {
			t.Logf("%s expected time not empty", success)
		}
	}
}

func TestLoggerNew(t *testing.T) {
	t.Log("Testing Second Logger")
	{
		var buff = bytes.NewBufferString("")
		SetOutput(buff)
		Info("reference:20180808003334343443", "name:example1", "class:testing")
		Warn("reference:20180808003334343443", "name:example1", "class:testing")
		Error("reference:20180808003334343443", "name:example1", "class:testing")

		splits := strings.Split(buff.String(), "\n")

		for _, split := range splits {
			if split == "" {
				break
			}
			var data struct {
				Reference string `json:"reference"`
				Name      string `json:"name"`
				Class     string `json:"class"`
				Level     string `json:"level"`
				Message   string `json:"msg"`
			}

			err := json.Unmarshal([]byte(split), &data)
			t.Log("Level", data.Level)
			t.Log("Message", data.Message)
			if err != nil {
				t.Fatalf("%s expected error nil, got %s", failed, err.Error())
			} else {
				t.Logf("%s expected error nil", success)
			}

			ref := "20180808003334343443"

			if data.Reference == ref {
				t.Logf("%s expected Reference == %s", success, ref)
			} else {
				t.Fatalf("%s expected Reference == %s, got %s", failed, ref, data.Reference)
			}

			nm := "example1"
			if data.Name == nm {
				t.Logf("%s expected Name == %s", success, nm)
			} else {
				t.Fatalf("%s expected Name == %s, got %s", failed, nm, data.Name)
			}

			cl := "testing"
			if data.Class == "testing" {
				t.Logf("%s expected Class == %s", success, cl)
			} else {
				t.Fatalf("%s expected Class == %s, got %s", failed, cl, data.Class)
			}
		}

	}

	t.Log("Testing Third typed logger")
	{
		var buff = bytes.NewBufferString("")
		SetOutput(buff)
		Info("reference: ", "201880888088003334344", " order_id:", 15)
		Warn("reference: ", "201880888088003334344", " order_id:", 15)
		Error("reference: ", "201880888088003334344", " order_id:", 15)

		splits := strings.Split(buff.String(), "\n")

		for _, split := range splits {
			if split == "" {
				break
			}
			var data struct {
				Reference string `json:"reference"`
				OrderId   string `json:"order_id"`
				Level     string `json:"level"`
				Message   string `json:"msg"`
			}
			err := json.Unmarshal([]byte(split), &data)
			if err != nil {
				t.Fatalf("%s expected error nil, got %s", failed, err.Error())
			} else {
				t.Logf("%s expected error nil", success)
			}
			t.Log("Level", data.Level)
			t.Log("Message", data.Message)

			ref := "201880888088003334344"
			if data.Reference == ref {
				t.Logf("%s expected reference == %s", success, ref)
			} else {
				t.Fatalf("%s expected reference == %s, got %s", failed, ref, data.Reference)
			}

			orderId := "15"
			if data.OrderId == orderId {
				t.Logf("%s expected order id == %s", success, orderId)
			} else {
				t.Fatalf("%s expected order id == %s, got %s", failed, orderId, data.OrderId)
			}
		}

	}
}

func TestFmt(t *testing.T) {
	t.Log("testing Sprintf formatting")
	{
		input := struct {
			Name    string
			Address string
			Rate    float64
		}{
			"Name Example", "Example Address", 2.87,
		}

		str1 := fmt.Sprintf("%v", input)
		str2 := fmt.Sprintf("%#v", input)

		t.Log("One", str1, "Two", str2)
	}
}

func TestJoiningInput(t *testing.T) {
	t.Log("Testing joiningInput")
	{
		result := joiningInput("one", "two", "three")
		t.Log("Result", result)
	}

	t.Log("Testing joiningInput")
	{
		result := joiningInput("one:", "two", "three")
		t.Log("Result", result)
	}
}

func CheckEmpty(t interface{}) bool {
	return reflect.DeepEqual(t, reflect.Zero(reflect.TypeOf(t)).Interface())
}
