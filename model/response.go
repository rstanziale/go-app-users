package model

func ResponseOK() map[string]bool {
	return map[string]bool{"ok": true}
}

func ResponseKO() map[string]bool {
	return map[string]bool{"ok": false}
}

func ResponseMessage(data any) map[string]any {
	return map[string]any{"response": data}
}

func ResponseError(err string) map[string]string {
	return map[string]string{"error": err}
}
