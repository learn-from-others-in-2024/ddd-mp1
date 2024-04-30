package transfers

import "time"

// Transfer contains attributes of a transfer entity
type Transfer struct {
	id        string
	date      time.Time
	accountId int64
}

// NewTransfer is the only way to create a transfer entity
func NewTransfer(id string, date time.Time, accountId int64) *Transfer {
	return &Transfer{
		id:        id,
		date:      date,
		accountId: accountId,
	}
}

// IsDateAllowed checks whether transfer is executed in the current day
func (t *Transfer) IsDateAllowed() bool {
	now := time.Now()

	return !t.date.After(now)
}
