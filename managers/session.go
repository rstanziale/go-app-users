package managers

var sessionMap = make(map[string]int)

// Add user info in session
func SetSession(sessionId string, userId int) {
	sessionMap[sessionId] = userId
}

// Return user info in session
func GetSession(sessionId string) (int, bool) {
	value, ok := sessionMap[sessionId]
	return value, ok
}

// Delete user info in session
func DeleteSession(sessionId string) {
	delete(sessionMap, sessionId)
}