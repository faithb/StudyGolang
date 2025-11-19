package model

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
)

type ItemStatus int

const (
	ItemStatusDoing ItemStatus = iota
	ItemStatusDone
	ItemStatusDeleted
)

var allItemStatus = [3]string{
	"Doing",
	"Done",
	"Deleted",
}

func (item *ItemStatus) String() string {
	return allItemStatus[*item]
}

func parseItemStatus(str string) (ItemStatus, error) {
	for i := range allItemStatus {
		if allItemStatus[i] == str {
			return ItemStatus(i), nil
		}
	}

	return ItemStatusDoing, errors.New("invalid ItemStatus: " + str)
}

func (item *ItemStatus) MarshalJSON() ([]byte, error) {
	if item == nil {
		return []byte(`null`), nil
	}
	return []byte(fmt.Sprintf(`"%s"`, item.String())), nil
}

func (item *ItemStatus) UnmarshalJSON(data []byte) error {
	str := strings.ReplaceAll(string(data), "\"", "")
	parsed, err := parseItemStatus(str)
	if err != nil {
		return err
	}

	*item = parsed
	return nil
}

func (item *ItemStatus) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprintf("Failed to unmarshal JSONB value:", value))
	}

	parsed, err := parseItemStatus(string(bytes))
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to unmarshal JSONB value:", err))
	}

	*item = parsed

	return nil
}

func (item *ItemStatus) Value() (driver.Value, error) {
	return item.String(), nil
}
