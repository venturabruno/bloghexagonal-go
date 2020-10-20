package domain

import (
	"database/sql/driver"
	"fmt"
	"reflect"
	"time"

	"github.com/go-sql-driver/mysql"
)

type NullTime mysql.NullTime

func (nullTime *NullTime) Scan(value interface{}) error {
	var t mysql.NullTime
	if err := t.Scan(value); err != nil {
		return err
	}

	if reflect.TypeOf(value) == nil {
		*nullTime = NullTime{t.Time, false}
		return nil
	}

	*nullTime = NullTime{t.Time, true}
	return nil
}

func (nullTime *NullTime) MarshalJSON() ([]byte, error) {
	if !nullTime.Valid {
		return []byte("null"), nil
	}
	val := fmt.Sprintf("\"%s\"", nullTime.Time.Format(time.RFC3339))
	return []byte(val), nil
}

func (nullTime NullTime) Value() (driver.Value, error) {
	if !nullTime.Valid {
		return nil, nil
	}
	return nullTime.Time, nil
}
