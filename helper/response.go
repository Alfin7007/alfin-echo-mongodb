package helper

import "net/http"

func BadRequest() (int, map[string]interface{}) {
	return http.StatusBadRequest, map[string]interface{}{
		"code":    http.StatusBadRequest,
		"message": "bad request",
	}
}

func BadRequestWithMSG(msg string) (int, map[string]interface{}) {
	return http.StatusBadRequest, map[string]interface{}{
		"code":    http.StatusBadRequest,
		"message": msg,
	}
}

func SuccessInsert() (int, map[string]interface{}) {
	return http.StatusCreated, map[string]interface{}{
		"code":    http.StatusCreated,
		"message": "success insert",
	}
}

func SuccessGetData(data interface{}) (int, map[string]interface{}) {
	return http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "success get data",
		"data":    data,
	}
}

func SuccessLogin(id, token string) (int, map[string]interface{}) {
	return http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "success login",
		"data": map[string]string{
			"id":    id,
			"token": token,
		},
	}
}

func Forbidden() (int, map[string]interface{}) {
	return http.StatusForbidden, map[string]interface{}{
		"code":    http.StatusForbidden,
		"message": "not authorized",
	}
}
