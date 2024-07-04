package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joaovit0r0/goOpportunities/schemas"
)

func CreateOpeningHadnler(ctx *gin.Context) {
	// First, create the "object"
	request := CreateOpeningRequest{}

	// After, fill the "object"
	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.ErrorF("validation error: %v", err.Error())
		SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	// it is necessary because the datatype of the request (CreateOpeningRequest) is not
	// the same as that db model (Opening)
	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.ErrorF("error creating opening: %v", err.Error())
		SendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(ctx, "create-opening", opening)
}
