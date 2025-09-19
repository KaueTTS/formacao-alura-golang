package api_common

import (
	api_common_dto "testes1/src/api/common/dto"
)

func NewErrorResponse(message string, err string, code int) api_common_dto.ErrorDto {
	return api_common_dto.ErrorDto{
		Message: message,
		Error:   err,
		Code:    code,
	}
}
