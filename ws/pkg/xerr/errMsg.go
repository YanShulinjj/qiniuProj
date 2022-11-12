package xerr

var message map[ErrCodeType]string

func init() {
	message = make(map[ErrCodeType]string)
	message[OK] = "SUCCESS"
	message[ServerCommonErr] = "服务器开小差啦,稍后再来试一试"
	message[ReuqestParamErr] = "参数错误"
	message[DbErr] = "数据库繁忙,请稍后再试"
	message[DbUpdateAffectedZeroErr] = "更新数据影响行数为0"
	message[DataNoExistErr] = "数据不存在"
	// 用户服务
	message[UserNotExistErr] = "该用户还没注册"
	message[UserExistedErr] = "该用户已经注册过啦，请直接登陆"
	// 画板服务
	message[PersistenceErr] = "持久化画板错误"
	message[PageExistedErr] = "该页面重复保存"
	message[PageNotExistErr] = "该页面不存在"
}

func MapErrMsg(errcode ErrCodeType) string {
	if msg, ok := message[errcode]; ok {
		return msg
	} else {
		return "服务器开小差啦,稍后再来试一试"
	}
}

func IsCodeErr(errcode ErrCodeType) bool {
	if _, ok := message[errcode]; ok {
		return true
	} else {
		return false
	}
}
