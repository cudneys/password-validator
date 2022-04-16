package api

import (
	models "github.com/cudneys/password-validator/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func getIntQP(c *gin.Context, key string, defaultValue string) (int, error) {
	length := c.DefaultQuery(key, defaultValue)
	lNum, err := strconv.Atoi(length)
	if err != nil {
		return 0, err
	}
	return lNum, nil
}

// @BasePath /v1
// passwordValidator godoc
// @Summary Password Validator
// @Schemes
// @Description Validates Passwords
// @Param        password   query      string  false  "Password to validate"
// @Produce json
// @Success 200 {object} models.Response
// @Router /validate [get]
func PasswordValidator(c *gin.Context) {
	var reqParams models.RequestParams

	// Bind the request params
	err := c.ShouldBind(&reqParams)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Error: string(err.Error()), IsValid: false})
		return
	}
	assessment, err := reqParams.Validate()
	if err != nil {
		c.JSON(499, models.Response{Error: string(err.Error()), IsValid: false})
		return
	}

	c.JSON(200, models.Response{IsValid: true, Assessment: assessment})

}
