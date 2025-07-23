package auth

import (
	"errors"
	"fmt"
	"github.com/arishahmad661/stealth_x_pi/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func SignUpHandler(c *gin.Context) {
	var reqUser SigningUpUser

	if err := c.ShouldBindJSON(&reqUser); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			missingFields := []string{}
			for _, field := range ve {
				missingFields = append(missingFields, field.Field())
			}
			utils.FailResponse(c, 400, "Missing or invalid fields: "+fmt.Sprint(missingFields), err)
			return
		}

		utils.FailResponse(c, 400, "Failed to bind JSON.", err)
		return
	}

	fmt.Println(reqUser)

	user, err := RegisterUser(reqUser)
	if err != nil {
		utils.FailResponse(c, 400, "Failed to register.", err)
	} else {
		utils.SuccessResponse(c, 200, "User added successfully.", user)
	}
}
