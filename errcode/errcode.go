package errcode

type Code int

type Message string

//go:generate stringer -type=Code -linecomment -output errcode_string.go
const (
	SystemError Code = -1  // SystemError 系统错误
	OK          Code = 200 // OK

	NoContent Code = 204 // NoContent 内容为空

	BadRequest Code = 400 // BadRequest 请求错误

	Authorization Code = 401 // Authorization 访问被拒绝

	NotFound Code = 404 // NotFound 资源不存在

	TokenIsNotAvailable Code = 1000 // TokenIsNotAvailable 服务令牌不可用

	DataNotFound Code = 4004 // DataNotFound 数据不存在

	ParamRequired Code = 4001 //ParamRequired 缺少必填参数

	LogicalException Code = 4002 // LogicalException 逻辑异常

	AccountFrozen Code = 4003 // AccountFrozen 账户被冻结

	DbQueryError Code = 1101 //DbQueryError 查询错误

	DbInsertError Code = 1102 //DbInsertError 创建错误

	DbUpdateError Code = 1103 //DbUpdateError 更新错误

	DbDeleteError Code = 1104 //DbDeleteError 删除错误

	PhoneNumHasBeenRegistered Code = 1301 //PhoneNumHasBeenRegistered 手机号已经被注册

	PasswordPatternError Code = 1302 //PasswordPatternError 密码格式不正确

	IdentifierHasBeenAuthorized = 1303 //IdentifierHasBeenAuthorized 第三方唯一标识已被授权

	PasswordError = 1304 //PasswordError 密码输入不正确

	ConfirmPassError = 1305 //ConfirmPassError 两次密码输入不一致

	PhoneNumPatternError = 1306 //PhoneNumPatternError 手机号格式不正确
	// 待续...

)
