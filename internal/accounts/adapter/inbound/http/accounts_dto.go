package http

import "mini-project-a/internal/accounts/core/entity"

type CreateAccountRequest struct {
	OwnerName   string  `json:"owner_name" validate:"required"`
	CitizenID   string  `json:"citizen_id" validate:"required"`
	PhoneNumber string  `json:"phone_number" validate:"required"`
	AccountType string  `json:"account_type" validate:"required,oneof=SAVING CURRENT"`
	Balance     float64 `json:"balance" validate:"required,gte=0"`
	Status      string  `json:"status" validate:"required,oneof=ACTIVE"`
}

func (r *CreateAccountRequest) ToAccountEntity() *entity.Accounts {
	return &entity.Accounts{
		OwnerName:   r.OwnerName,
		CitizenID:   r.CitizenID,
		PhoneNumber: r.PhoneNumber,
		AccountType: r.AccountType,
		Balance:     r.Balance,
		Status:      "ACTIVE",
	}
}
