package model

// Return JSON just Ok response
func ResponseOK() map[string]bool {
	return map[string]bool{"ok": true}
}

// Return JSON just KO message
func ResponseKO() map[string]bool {
	return map[string]bool{"ok": false}
}

// Return JSON message response
func ResponseResult(data any) map[string]any {
	return map[string]any{"response": data, "ok": true}
}

// Return JSON message error
func ResponseError(err string) map[string]any {
	return map[string]any{"error": err, "ok": false}
}
