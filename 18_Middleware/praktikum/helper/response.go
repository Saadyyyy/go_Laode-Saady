package helper

//respon Err
func ResponErrId(massage string) map[string]interface{} {
	return map[string]interface{}{
		"massage": "Invalid Id",
	}
}

func ResponErr(massage string) map[string]interface{} {
	return map[string]interface{}{
		"massage": massage,
	}
}

func ResponSuccess(massage string, data any) map[string]interface{} {
	return map[string]interface{}{
		"massage": massage,
		"Data":    data,
	}
}
