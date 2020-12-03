package mysql_utils

import (
	"bookstore_users-api/utils/errors"
	"strings"

	"github.com/go-sql-driver/mysql"
)

const (
	errorNoRows = "no rows in result set"
)

func ParseError(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundError("no record matching given id")
		}
		return errors.NewInternalServerError("error parsing database response")
	}

	switch sqlErr.Number {
	// 1062 is sqlError for duplicated key
	case 1062:
		return errors.NewBadRequestError("duplicate key must be unique")
	}
	return errors.NewInternalServerError("error when processing request")
}
