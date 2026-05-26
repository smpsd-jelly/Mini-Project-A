package http

type CreateAccountRequest struct {
	AccountNumber string  `json:"account_number" validate:"required"`
	OwnerName     string  `json:"owner_name" validate:"required"`
	CitizenID     string  `json:"citizen_id" validate:"required"`
	PhoneNumber   string  `json:"phone_number" validate:"required"`
	AccountType   string  `json:"account_type" validate:"required,oneof=SAVING CURRENT"`
	Balance       float64 `json:"balance" validate:"required,gte=0"`
	Status        string  `json:"status" validate:"required,oneof=ACTIVE CLOSED"`
	CreatedAt     string  `json:"created_at" validate:"required,datetime=2006-01-02T15:04:05Z07:00"`
	UpdatedAt     *string `json:"updated_at,omitempty" validate:"omitempty,datetime=2006-01-02T15:04:05Z07:00"`
}
