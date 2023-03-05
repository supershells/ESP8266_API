package controllers

import (
	cfg "esp8266_api/configs"
	"esp8266_api/services"
	valid "esp8266_api/util/validator"
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/dgrijalva/jwt-go"
	//jwtware "github.com/gofiber/jwt/v2"
)

type userHandler struct {
	userSrv services.UserService
}

func NewUserHandler(userSrv services.UserService) UserHandler {
	return userHandler{userSrv: userSrv}
}

func (h userHandler) Login(c *fiber.Ctx) error {
	config := cfg.GetConfig()
	request := services.UserLogin{}
	if err := valid.ParseBodyAndValidate(c, &request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err)
	}

	user, err := h.userSrv.GetUser(request)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(err)
	}

	cliams := jwt.StandardClaims{
		Issuer:    user.Username,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, cliams)
	token, err := jwtToken.SignedString([]byte(config.JWTSECRET))
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return c.JSON(fiber.Map{
		"data":     user,
		"jwtToken": token,
	})

}

func (h userHandler) Register(c *fiber.Ctx) error {
	request := services.NewUserRequest{}

	if err := valid.ParseBodyAndValidate(c, &request); err != nil {
		return c.Status(http.StatusBadRequest).JSON((err))
	}
	chkuser, _ := h.userSrv.GetUserOne(request.Username)
	fmt.Println(chkuser)
	if chkuser != nil {
		return c.Status(http.StatusConflict).JSON(fiber.Map{
			"message": " Username already exists",
		})
	}

	user, err := h.userSrv.NewUser(request)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err)
	}
	return c.Status(http.StatusCreated).JSON(user)
}

func (h userHandler) LoadUserList(c *fiber.Ctx) error {
	users, err := h.userSrv.GetUserAll()
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(err)
	}

	return c.Status(http.StatusAccepted).JSON(users)
}
