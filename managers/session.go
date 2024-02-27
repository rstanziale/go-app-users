package managers

var sessionMap = make(map[string]int)

func SetSession(sessionId string, userId int) {
	sessionMap[sessionId] = userId
}

func GetSession(sessionId string) (int, bool) {
	value, ok := sessionMap[sessionId]
	return value, ok
}

func DeleteSession(sessionId string) {
	delete(sessionMap, sessionId)
}