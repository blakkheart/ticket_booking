package model_interface

import (
	"ticket-booking/models/domain/account"
)

type IAccountRepository interface {
	IBaseRepository[account.Account, account.AccountIn]
}
