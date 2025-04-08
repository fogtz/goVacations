package handler

import "github.com/gin-gonic/gin"

func CreateVacancyHandler(context *gin.Context) {
	request := CreateVacancyRequest{}
	context.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("Error validating request: %v", err.Error())
		return
	}
}

func ShowVacancyHandler(context *gin.Context) {
}

func DeleteVacancyHandler(context *gin.Context) {
}

func UpdateVacancyHandler(context *gin.Context) {
}

func ListVacancyHandler(context *gin.Context) {
}
