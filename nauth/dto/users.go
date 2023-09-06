package dto

import "github.com/RichXan/sdk-go/common"

type UserListDto struct {
	Data []User
	common.BaseResponse
}
