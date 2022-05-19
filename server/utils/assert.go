package utils

const AlertMessage = "异常::"

func NotNil(o interface{}, message string) {
	if o == nil {
		panic(AlertMessage + message)
	}
}

func IsTrue(expression bool, message string) {
	if !expression {
		panic(AlertMessage + message)
	}
}

func Nil(o interface{}, message string) {
	if o != nil {
		panic(AlertMessage + message)
	}
}
