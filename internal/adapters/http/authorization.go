package httpapp

import (
	"net/http"

	"github.com/gin-gonic/gin"
)
const (
	accessHeader  = "access_token"
	refreshHeader = "refresh_header"
)

type UserProjection struct {
	Login string `json:"login"`
	Email string `json:"email"`
}


func(a *Adapter) Login(ctx *gin.Context) {
	login, password, ok := ctx.Request.BasicAuth()
	if !ok{
		ctx.Header("WWW-Authenticate", "Basic")
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "not auth data",
		})
		return
	}
	accessToken, refreshToken, err := a.auth.Login(ctx, login, password)
	if err != nil {
		a.BindError(ctx, err)
		return
	}
	ctx.SetCookie(accessHeader, accessToken, 0, "", "auth", true, true)
	ctx.SetCookie(refreshHeader, refreshToken, 0, "", "auth", true, true)

}

func(a *Adapter) Logout(ctx *gin.Context) {
	login, password, ok := ctx.Request.BasicAuth()
	if !ok{
		ctx.Header("WWW-Authenticate", "Basic")
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "not auth data",
		})
		return
	}
	ok, err := a.auth.Logout(ctx, login, password)
	if err != nil {
		a.BindError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
}

func(a *Adapter) Verify(ctx *gin.Context) {
	access, err := ctx.Cookie(accessHeader)
	if err != nil {
		return
	}
	refresh, err := ctx.Cookie(refreshHeader)
	if err != nil {
		return
	}

	user, err := a.auth.Verify(ctx, access, refresh)
	if err != nil {
		a.BindError(ctx, err)
		return
	}
	ctx.SetCookie(accessHeader, user.AccessToken, 0, "", "auth", true, true)
	ctx.SetCookie(refreshHeader, user.RefreshToken, 0, "", "auth", true, true)

	ctx.JSON(http.StatusOK, UserProjection{
		Login: user.User.Login,
		Email: user.User.Email,
	})
}


func(a *Adapter) Singup(ctx *gin.Context) {
	login, password, ok := ctx.Request.BasicAuth()
	if !ok{
		ctx.Header("WWW-Authenticate", "Basic")
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "not auth data",
		})
		return
	}

	u, err := a.auth.Singup(ctx, login, password, "")
	if err != nil {
		a.BindError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, UserProjection{
		Login: u.Login,
		Email: u.Email,
	})

}
