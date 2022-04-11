package domain

var NewMap map[string]string = make(map[string]string)

func Create() {
	NewMap["firstKey"] = "First value"
	NewMap["secondKey"] = "Second value"
	NewMap["thirdKey"] = "Third value"
	NewMap["fourthKey"] = "Fourth value"
	NewMap["fifthKey"] = "Fifth value"

}

func GetValue(key string) string {
	return NewMap[key]
}

func PostValue(key string, value string) {
	NewMap[key] = value
}
