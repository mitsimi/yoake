package env

import "io/ioutil"

func WriteLog(message string) {
	// Write log message to file
	err := ioutil.WriteFile(LogFile(), []byte(message+"\n"), 0644)
	if err != nil {
		panic(err)
	}
}
