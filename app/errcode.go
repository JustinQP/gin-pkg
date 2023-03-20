package app

const (
	ErrAuthFailed   = 401
	ErrForbidden    = 403
	ErrServerFailed = 500
)

const (
	ErrForm = iota + 2001
	ErrNotFound
	ErrDataExist
	ErrParamUnknown
	ErrOpBusy
	ErrValueOutOfRange
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

	ErrForm:              "form validation error",
	ErrDataExist:         "data already exists",
	ErrNotFound:          "not Found",
	ErrParamUnknown:      "unknown parameter",
	ErrOpBusy:            "frequent operation",
	ErrValueOutOfRange:   "the data length or size exceeds the limit",
	ErrRpcConnect:        "RPC connection failed",
	ErrTransactionFailed: "transaction failed",

	ErrUserNameExist:       "username already exists",
	ErrUserEmailExist:      "email already exists",
	ErrUserPwd:             "wrong user name or password",
	ErrUserPwdUnMatch:      "the two passwords do not match",
	ErrUserPwdNotChange:    "the old and new passwords are the same",
	ErrUser404:             "user does not exist",
	ErrFollow:              "follow failure",
	ErrFollowed:            "followed",
	ErrFollowedNo:          "not following",
	ErrEmailSend:           "Email sending failed",
	ErrEmailInvalid:        "Email verification code check failed",
	ErrCollectionNameExit:  "Collection name already exists",
	ErrAddressInvalid:      "Address verification failed",
	ErrAccountInsufficient: "Insufficient balance",

	ErrUserAddressExist: "Wallet address is bound",
	ErrAddressNotBound:  "Wallet not bound",
	ErrSignatureInvalid: "Signature verification failed",

	ErrNftBalances:   "Insufficient account balance",
	ErrUserMSGClosed: "User has closed MSG",
	ErrNftCreate:     "NFT create failed",

	ErrOrderMyself:           "Can't buy own order",
	ErrOrderBuild:            "Order build failed",
	ErrOrderFinishOrCanceled: "Order completed or cancelled",

	ErrFileFile: "File storage failed",
	ErrFileType: "file type not allowed",
	ErrFileSize: "file size exceeds limit",

	ErrDB:       "Data operation failed",
	ErrDBRead:   "Data read failed",
	ErrDBWrite:  "Data write failed",
	ErrDBUpdate: "Data update failed",
	ErrDBDelete: "Data deletion failed",

	ErrWsMsgSend: "Failed to send message",
	ErrWsMsgFmt:  "Message format error",
	ErrWsMsg500:  "Service exception",
	ErrWsToken:   "Invalid or not found",
	ErrWsAuth:    "Authentication failed",
	ErrWsMsgType: "Incorrect message type",

	ErrParamInvalid: "Invalid parameter",
}
