package xerr

type ErrCodeType uint32

const (
	// 成功返回
	OK ErrCodeType = 200
	/**(前3位代表业务, 后三位代表具体功能)**/
	/**全局错误码*/
	// 服务器开小差
	ServerCommonErr ErrCodeType = 100000 + iota
	// 请求参数错误
	ReuqestParamErr
	// 数据库繁忙,请稍后再试
	DbErr
	// 更新数据影响行数为0
	DbUpdateAffectedZeroErr
	// 数据不存在
	DataNoExistErr

	// 用户服务
	UserNotExistErr ErrCodeType = 200000 + iota
	UserExistedErr

	// 画板服务
	PersistenceErr ErrCodeType = 300000 + iota
	PageExistedErr
	PageNotExistErr
)
