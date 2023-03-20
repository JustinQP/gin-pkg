package app

const (
	ErrAuthFailed   = 401
	ErrForbidden    = 403
	ErrServerFailed = 500
)

const (
	ErrForm         = iota + 2001
	ErrNotFound     // 数据不存在
	ErrDataExist    // 数据已存在
	ErrParamUnknown // 参数类型错误
)

const (
	ErrUserNameExist = iota + 4001
	ErrUserPwd
	ErrUserPwdUnMatch
	ErrUserPwdNotChange
	ErrUserEmailExist
	ErrUser404
	ErrFileFile
	ErrFileType
	ErrFileSize
	ErrFollow
	ErrFollowed
	ErrFollowedNo
	ErrEmailSend
	ErrEmailInvalid
	ErrCollectionNameExit
	ErrAddressInvalid
	ErrAddressNotBound
	ErrAccountInsufficient
	ErrUserAddressExist
	ErrSignatureInvalid // signature verification failure
	ErrNftCreate
	ErrParamInvalid

	ErrUserMSGClosed
	ErrNftBalances

	ErrOrderMyself
	ErrOrderBuild
	ErrOrderFinishOrCanceled

	ErrRpcConnect
	ErrTransactionFailed

	ErrWsMsg500
	ErrWsMsgSend
	ErrWsMsgFmt
	ErrWsToken
	ErrWsAuth
	ErrWsMsgType
)

const (
	ErrDB = iota + 5001
	ErrDBRead
	ErrDBWrite
	ErrDBUpdate
	ErrDBDelete
)

var ErrInfo = map[int]string{
	ErrAuthFailed:   "Unauthorized",          // 401
	ErrForbidden:    "Forbidden",             // 403
	ErrServerFailed: "Server internal error", // 500

}
