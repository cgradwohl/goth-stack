package handler

import (
	"github.com/cgradwohl/goth-stack/app/model"
	"github.com/cgradwohl/goth-stack/app/view/user"
	"github.com/labstack/echo/v4"
)

type UserHandler struct{}

// convert this to type HandlerFunc func(Context) error
// type HandlerFunc func(Context) error

func (h UserHandler) HandleUserShow(c echo.Context) error {
	u := model.User{
		Email: "c@cs.com",
	}

	return render(c, user.Show(u))

}
