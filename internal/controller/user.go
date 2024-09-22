package controller

import (
	"myapi/configs"
	"myapi/internal/bootstrap/web"
	"myapi/internal/request"
	_ "myapi/internal/response"
	"myapi/internal/service"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	actionPrefix  string
	userService *service.UserService
}

func NewUserController(app *fiber.App, mid *web.Middlewares, userService *service.UserService) {
	userController := &UserController{
		actionPrefix: `/users`,
		userService: userService,
	}

	app.Route(userController.actionPrefix, func(router fiber.Router) {
		router.Get(userDetailsAction, mid.OPT_Auth, userController.UserDetails)
	})
}

const userDetailsAction = `/details`

// @Summary 	Get user data by session id (via cookie / api key)
// @Tags			User
// @Success		200 {object} response.UserDetailsOut "User Details Out"
// @Produce		json
// @Router		/users/details [get]
func (_uc *UserController) UserDetails(c *fiber.Ctx) error {
	sess, _, err := getSession(_uc.userService.Db, c)
	if err != nil {
		c.ClearCookie(configs.AUTH_COOKIE)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			`errors`: err.Error(),
		})
	}

	in := request.UserDetailsIn{
		UserID: sess.UserID,
	}

	out, err := _uc.userService.UserDetails(in)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			`errors`: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(out)
}

