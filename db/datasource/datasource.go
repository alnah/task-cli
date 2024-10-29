package db

import (
	"fmt"
	"net"
	"reflect"
	"strconv"
)

const requiredFields = "User, Password, Host, Port, DBName"

var (
	ErrUserLength     = fmt.Errorf("user must be 64 characters max")
	ErrPasswordLength = fmt.Errorf("user must be 64 characters max")
	ErrInvalidHost    = fmt.Errorf("invalid host")
	ErrInvalidPort    = fmt.Errorf("invalid port")
	ErrDBNameLength   = fmt.Errorf("DBName must be 64 characters max")
	ErrMissingFields  = fmt.Errorf("missing required fields: " + requiredFields)
)

type GenerateDSNParams struct {
	User     string
	Password string
	Host     string
	Port     string
	DBName   string
}

// GenerateDatasource constructs a Data Source Name for a PostgreSQL connection.
func GenerateDatasource(params GenerateDSNParams) (string, error) {
	requiredFields := reflect.ValueOf(params)
	for i := 0; i < requiredFields.NumField(); i++ {
		if requiredFields.Field(i).String() == "" {
			return "", ErrMissingFields
		}
	}

	if len(params.User) > 64 {
		return "", ErrUserLength
	}
	if len(params.Password) > 64 {
		return "", ErrPasswordLength
	}
	if !isValidHost(params.Host) {
		return "", ErrInvalidHost
	}
	if !isValidPort(params.Port) {
		return "", ErrInvalidPort
	}
	if len(params.DBName) > 64 {
		return "", ErrDBNameLength
	}

	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		params.User,
		params.Password,
		params.Host,
		params.Port,
		params.DBName,
	), nil
}

func isValidHost(host string) bool {
	if net.ParseIP(host) != nil {
		return true
	}
	_, err := net.LookupHost(host)
	return err == nil
}

func isValidPort(port string) bool {
	p, err := strconv.Atoi(port)
	if err != nil {
		return false
	}

	return p > 0 && p <= 65535
}
