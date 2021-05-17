package src

import (
	"os"
	"log"
	"errors"
)

type Logs struct {
	Log *log.Logger
}

type ErrorsInfo struct {
	log_error	bool
	errors		[]error
	length		int
	current_error	error
	logs		[]Logs
}

func Default(log_error bool) *ErrorsInfo {
	info := &ErrorsInfo{log_error: log_error}

	info.length += 1
	info.errors = make([]error, info.length)
	info.current_error = nil

	if info.log_error == true {
		_, err := os.Stat("errors")

		if err != nil {
			if os.IsNotExist(err) {
				_, err := os.Create("errors")

				if err != nil {
					log.Fatal(err)
				}
			}
		}

		file, _err := os.OpenFile("errors", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)

		if _err != nil {
			log.Fatal(err)
		}

		_log := Logs{}
		_log.Log = log.New(file, "ERROR:", log.Ldate)

		info.logs = append(info.logs, _log)
	}
	return info
}

func (err_info *ErrorsInfo) update(err error) {
	err_info.errors = append(err_info.errors, err)
	err_info.length += 1

	err_info.current_error = err_info.errors[err_info.length - 1]
}

func (err_info *ErrorsInfo) file_exists(filename string) (string, error) {
	_, err := os.Stat(filename)

	if err != nil {
		err_info.update(err)
		return "", errors.New("File " + filename + " does not exist")
	}

	return filename, nil
}

func (errs *ErrorsInfo) dereference_error(at_index int) error {
	if !(at_index > errs.length) {
		return errs.errors[at_index]
	}
	return nil
}

func (err_info *ErrorsInfo) is_dir(dir string) (string, error) {
	_, err := os.Stat(dir)

	if err != nil {
		err_info.update(err)
		return "", errors.New("Directory " + dir + " does not exist")
	}

	return dir, nil
}
