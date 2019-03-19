package logger

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"runtime"
	"strings"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
}

func getCallerInfo() (file string, line int, ok bool) {
	_, file, line, ok = runtime.Caller(3)
	return
}

// SetLevel sets the standard entry level
func SetLevel(levelStr string) error {
	level, err := log.ParseLevel(levelStr)
	if err != nil {
		return err
	}

	log.SetLevel(level)
	return nil
}

// SetOutput sets the standard logger output.
func SetOutput(output io.Writer) {
	log.SetOutput(output)
}

// An entryWithInputs is the final or intermediate logging entry. It contains all
// the fields passed with WithField{,s}. It's finally logged when Debug, Info,
// Warn, Error, Fatal or Panic is called on it. These objects can be reused and
// passed around as much as you wish to avoid field duplication.
func entryWithInputs(inputs ...interface{}) *log.Entry {

	file, line, _ := getCallerInfo()
	fields := log.Fields{}
	fields["on"] = fmt.Sprintf("%s:%d", file, line)
	var (
		skey   string
		splits []string
	)

	tmpResult := joiningInput(inputs...)
	for i, input := range tmpResult {
		splits = strings.Split(input, ":")
		if len(splits) > 1 {
			fields[strings.TrimSpace(splits[0])] = strings.TrimSpace(splits[1])
		} else {
			skey = fmt.Sprintf("%s-%d", "input", i)
			fields[skey] = input
		}
	}

	return log.WithFields(fields)
}

func Debug(inputs ...interface{}) {
	entryWithInputs(inputs...).Debug(inputs...)
}

func Info(inputs ...interface{}) {
	entryWithInputs(inputs...).Info(inputs...)
}

func Warn(inputs ...interface{}) {
	entryWithInputs(inputs...).Warn(inputs...)
}

func Error(inputs ...interface{}) {
	entryWithInputs(inputs...).Error(inputs...)
}

func Fatal(inputs ...interface{}) {
	entryWithInputs(inputs...).Fatal(inputs...)
}

func joiningInput(inputs ...interface{}) []string {
	var (
		tmp, splits     []string
		input, tmpInput string
	)

	for i := 0; i < len(inputs); i++ {
		input = fmt.Sprintf("%v", inputs[i])
		if strings.Contains(input, ":") {
			splits = strings.Split(input, ":")
			if len(splits) == 2 && strings.TrimSpace(splits[1]) == "" && len(inputs) > 2{
				i += 1
				tmpInput = fmt.Sprintf("%s%v", input, inputs[i])
			} else {
				tmpInput = input
			}
			tmp = append(tmp, tmpInput)
		} else {
			tmp = append(tmp, input)
		}
	}

	return tmp
}
