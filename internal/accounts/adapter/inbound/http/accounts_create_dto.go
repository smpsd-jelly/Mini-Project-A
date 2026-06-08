package http

import "mini-project-a/internal/accounts/core/entity"

type CreateAccountRequest struct {
	OwnerName   string  `json:"owner_name" binding:"required"`
	CitizenID   string  `json:"citizen_id" binding:"required"`
	PhoneNumber string  `json:"phone_number" binding:"required"`
	AccountType string  `json:"account_type" binding:"required,oneof=SAVING CURRENT"`
	Balance     float64 `json:"balance" binding:"gte=0"`
}

func (r *CreateAccountRequest) ToAccountEntity() *entity.Accounts {
	return &entity.Accounts{
		OwnerName:   r.OwnerName,
		CitizenID:   r.CitizenID,
		PhoneNumber: r.PhoneNumber,
		AccountType: r.AccountType,
		Balance:     r.Balance,
	}
}
