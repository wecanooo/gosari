package context

import (
	"github.com/wecanooo/gosari/core/errno"
	"github.com/wecanooo/gosari/core/pkg/jwt"
)

type TokenResp struct {
	AccessToken  *jwt.AppJWTInfo `json:"access_token"`
	RefreshToken *jwt.AppJWTInfo `json:"refresh_token,omitempty"`
}

func (c *AppContext) TokenSign(userID uint) (*TokenResp, error) {
	a, r, err := jwt.CreateToken(userID)
	if err != nil {
		return nil, errno.TokenErr.WithErr(err)
	}

	return &TokenResp{AccessToken: a, RefreshToken: r}, nil
}

func (c *AppContext) TokenRefresh(t string) (*TokenResp, error) {
	t, err := jwt.GetToken(c.Context)
	if err != nil {
		return nil, errno.TokenErr.WithErr(err)
	}

	td, err := jwt.RefreshToken(t)
	if err != nil {
		return nil, errno.TokenErr.WithErr(err)
	}

	return &TokenResp{AccessToken: td}, nil
}
