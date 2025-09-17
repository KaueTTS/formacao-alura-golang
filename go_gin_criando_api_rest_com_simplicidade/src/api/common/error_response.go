package api_common

import (
	api_common_dto "go_gin_criando_api_rest_com_simplicidade/src/api/common/dto"
)

func NewErrorResponse(message string, err string, code int) api_common_dto.ErrorDto {
	return api_common_dto.ErrorDto{
		Message: message,
		Error:   err,
		Code:    code,
	}
}
