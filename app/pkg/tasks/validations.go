package tasks

import (
	"encoding/json"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
)

type incomingPostData struct {
	ID        pgtype.UUID `json:"id"`
	Text      string      `json:"text"`
	Done      bool        `json:"done"`
	UpdatedAt pgtype.Date `json:"updatedAt"`
}

// This function process and validate the incoming request from the user
func ProcessAndValidateIncomingTasks(data []byte) ([]incomingPostData, error) {
	parsedDate := make([]incomingPostData, 0)
	err := json.Unmarshal(data, &parsedDate)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}
	for i := 0; i < len(parsedDate); i++ {
		if !parsedDate[i].ID.Valid {
			return nil, fmt.Errorf("invalid UUID at index %d", i)
		}
	}
	return parsedDate, nil
}
