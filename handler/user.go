package handler

import (
	"github.com/cgradwohl/goth-stack/model"
	"github.com/cgradwohl/goth-stack/view/user"
	"github.com/labstack/echo/v4"
)

type UserHandler struct{}

func (h UserHandler) HandleUserShow(c echo.Context) error {
	u := model.User{
		Email: "c@cs.com",
	}
	return render(c, user.Show(u))

}
