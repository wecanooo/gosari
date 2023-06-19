package context

import (
	"strconv"

	"github.com/wecanooo/gosari/core/errno"
)

func (c *AppContext) IntParam(key ...string) (int, error) {
	k := key[0]
	if k == "" {
		k = "id"
	}

	i, err := strconv.Atoi(c.Param(k))
	if err != nil {
		return 0, errno.ReqErr.WithErr(err)
	}

	return i, nil
}

func (c *AppContext) IntQuery(key ...string) (int, error) {
	k := key[0]
	if k == "" {
		k = "id"
	}

	i, err := strconv.Atoi(c.QueryParam(k))
	if err != nil {
		return 0, errno.ReqErr.WithErr(err)
	}

	return i, nil
}
