package dto

import (
	"time"
)

// TransactionMessage defines the standardized transaction message used across services
type TransactionMessage struct {
	ID              string  `json:"id"`
	ReferenceID     string  `json:"reference_id"`
	SessionID       string  `json:"session_id"`
	DebitFrom       string  `json:"debit_from"`
	SenderType      string  `json:"sender_type"`
	CreditTo        string  `json:"credit_to"`
	ReceiverType    string  `json:"receiver_type"`
	TransactionType string  `json:"transaction_type"` // PAYMENT, REFUND, etc.
	Amount          float64 `json:"amount"`
	Status          string  `json:"status"` // PENDING, SUCCESS, FAILED
	Metadata        string  `json:"metadata,omitempty"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
