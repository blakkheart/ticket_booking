package model_interface

import (
	"ticket-booking/models/domain/account"
	"ticket-booking/repository"
)

type IAccountRepository interface {
	IBaseRepository[repository.Account, account.AccountIn]
}
