package common

func LogInfo(logs *Logger, msg string) {
	logs.Warn(msg)
}

func LogWarning(logs *Logger, err error) {
	if err != nil {
		logs.Warn(err.Error())
	}
}

func LogError(logs *Logger, err error) {
	if err != nil {
		logs.Fatal(err.Error())
	}
}
