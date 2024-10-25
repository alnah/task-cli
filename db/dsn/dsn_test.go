package db

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateDSN(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name    string
		params  GenerateDSNParams
		wantDsn string
		wantErr error
	}{
		{
			name: "Valid Input",
			params: GenerateDSNParams{
				User:     "john",
				Password: "r@nd0mp@ssw0rd",
				Host:     "localhost",
				Port:     "5432",
				DBName:   "test_db",
			},
			wantDsn: "postgres://john:r@nd0mp@ssw0rd@localhost:5432/test_db?sslmode=disable",
			wantErr: nil,
		},
		{
			name: "Invalid User",
			params: GenerateDSNParams{
				User:     strings.Repeat("a", 100),
				Password: "r@nd0mp@ssw0rd",
				Host:     "localhost",
				Port:     "5432",
				DBName:   "test_db",
			},
			wantDsn: "",
			wantErr: ErrUserLength,
		},
		{
			name: "Invalid Password",
			params: GenerateDSNParams{
				User:     "john",
				Password: strings.Repeat("a", 100),
				Host:     "localhost",
				Port:     "5432",
				DBName:   "test_db",
			},
			wantDsn: "",
			wantErr: ErrPasswordLength,
		},
		{
			name: "Invalid Host",
			params: GenerateDSNParams{
				User:     "john",
				Password: "r@nd0mp@ssw0rd",
				Host:     "34.111111.23",
				Port:     "5432",
				DBName:   "test_db",
			},
			wantDsn: "",
			wantErr: ErrInvalidHost,
		},
		{
			name: "Invalid Port",
			params: GenerateDSNParams{
				User:     "john",
				Password: "r@nd0mp@ssw0rd",
				Host:     "localhost",
				Port:     "invalid_port",
				DBName:   "test_db",
			},
			wantDsn: "",
			wantErr: ErrInvalidPort,
		},
		{
			name: "Empty Password",
			params: GenerateDSNParams{
				User:     "john",
				Password: "",
				Host:     "localhost",
				Port:     "5432",
				DBName:   "test_db",
			},
			wantDsn: "",
			wantErr: ErrMissingFields,
		},
		{
			name: "Invalid DBName",
			params: GenerateDSNParams{
				User:     "john",
				Password: "r@nd0mp@ssw0rd",
				Host:     "localhost",
				Port:     "5432",
				DBName:   strings.Repeat("a", 100),
			},
			wantDsn: "",
			wantErr: ErrDBNameLength,
		},
		{
			name: "Empty Host",
			params: GenerateDSNParams{
				User:     "john",
				Password: "r@nd0mp@ssw0rd",
				Host:     "",
				Port:     "5432",
				DBName:   "test_db",
			},
			wantDsn: "",
			wantErr: ErrMissingFields,
		},
		{
			name: "Empty Port",
			params: GenerateDSNParams{
				User:     "john",
				Password: "r@nd0mp@ssw0rd",
				Host:     "localhost",
				Port:     "",
				DBName:   "test_db",
			},
			wantDsn: "",
			wantErr: ErrMissingFields,
		},
		{
			name: "Empty DBName",
			params: GenerateDSNParams{
				User:     "john",
				Password: "r@nd0mp@ssw0rd",
				Host:     "localhost",
				Port:     "5432",
				DBName:   "",
			},
			wantDsn: "",
			wantErr: ErrMissingFields,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gotDsn, err := GenerateDSN(tc.params)
			if err != nil {
				if tc.wantErr != nil {
					require.EqualError(t, err, tc.wantErr.Error())
				} else {
					t.Fatalf("unexpected error: %v", err)
				}
				require.Equal(t, gotDsn, tc.wantDsn)
			} else {
				require.Equal(t, gotDsn, tc.wantDsn)
			}
		})
	}
}
