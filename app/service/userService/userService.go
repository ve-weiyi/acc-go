package userService

import (
	"acc/app/model/bean"
	"acc/app/model/dao"
	"acc/app/model/dto"
	"acc/app/model/entity"
	"acc/config"
	"acc/lib/errCode"
	"acc/lib/jjwt"
	"acc/lib/logger"
	"acc/lib/orm"
)

func UserGetInfoById(uid int) (*entity.UserInfo, *errCode.ApiError) {
	var q = dao.Use(orm.DB()).UserInfo
	user, _ := q.Where(q.ID.Eq(uid)).First()
	return user, nil
}

func UserLogin(auth dto.AuthReq) (*bean.LoginResp, error) {
	logger.Debug("username:" + auth.Username)
	logger.Debug("password:" + auth.Password)

	userAuth, err := userAuthCheckLogin(auth.Username, auth.Password)
	if err != nil {
		return nil, err
	}

	user, apierr := UserGetInfoById(userAuth.ID)

	token := jjwt.CreateToken(*userAuth)

	login := &bean.LoginResp{
		Uid:      userAuth.ID,
		Username: userAuth.Username,
		Details:  user,
		Token:    config.AppConfig.JwtTokenHeader + " " + token,
	}

	return login, apierr
}

// user_auth
func userAuthCheckLogin(username string, password string) (*entity.UserAuth, *errCode.ApiError) {
	q := dao.Use(orm.DB()).UserAuth

	auth, err := q.Where(q.Username.Eq(username)).First()
	if err != nil {
		return nil, errCode.NewErrorMsg("用户名不存在")
	}

	if auth.Password != password {
		return nil, errCode.NewErrorMsg("密码不正确")
	}

	return auth, nil
}
