package erro

import "net/http"

const (
	OKCode              = 0
	NoRouteMatchedCode  = 10000
	UnKnownCode         = 10050
	NoMethodMatchedCode = 20000
	UserNotFoundCode    = 30000
	UserExistCode       = 30001
	ParamsErrCode       = 40000
	ForbiddenCode       = 60000
	UnauthorizedCode    = 70000
)

var (
	OK              = NewHttpError(OKCode, http.StatusOK, "OK!")
	NoRouteMatched  = NewHttpError(NoRouteMatchedCode, http.StatusNotFound, "路由不存在")
	NoMethodMatched = NewHttpError(NoMethodMatchedCode, http.StatusForbidden, "请求方法不允许")
	UnKnown         = NewHttpError(UnKnownCode, http.StatusBadRequest, "服务器未知错误")
	UserNotFound    = NewHttpError(UserNotFoundCode, http.StatusNotFound, "没有找到用户")
	UserExist       = NewHttpError(UserExistCode, http.StatusBadRequest, "用户已经存在")
	ParamsErr       = NewHttpError(ParamsErrCode, http.StatusBadRequest, "参数错误")
	Forbidden       = NewHttpError(ForbiddenCode, http.StatusForbidden, "禁止访问")
	Unauthorized    = NewHttpError(UnauthorizedCode, http.StatusUnauthorized, "认证失败")
)
