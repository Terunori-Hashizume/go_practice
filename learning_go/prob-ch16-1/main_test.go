package main

import (
	"strings"
	"testing"
)

func TestValidateStringLength(t *testing.T) {
	type User struct {
		Name  string `minStrlen:"3"`
		Email string `minStrlen:"5"`
		Note  string
	}

	tests := []struct {
		name        string
		input       any
		wantErr     bool
		containsAll []string
	}{
		{
			name:    "valid struct",
			input:   User{Name: "John", Email: "abcde"},
			wantErr: false,
		},
		{
			name:        "single invalid field",
			input:       User{Name: "Jo", Email: "abcde"},
			wantErr:     true,
			containsAll: []string{"Name", "min=3"},
		},
		{
			name:        "multiple invalid fields aggregated",
			input:       User{Name: "Jo", Email: "a@b"},
			wantErr:     true,
			containsAll: []string{"Name", "Email"},
		},
		{
			name:    "pointer to struct",
			input:   &User{Name: "John", Email: "abcde"},
			wantErr: false,
		},
		{
			name:        "non struct input",
			input:       "not-struct",
			wantErr:     true,
			containsAll: []string{"input must be a struct"},
		},
		{
			name:        "nil pointer input",
			input:       (*User)(nil),
			wantErr:     true,
			containsAll: []string{"input must be a struct"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateStringLength(tt.input)
			if tt.wantErr && err == nil {
				t.Fatalf("expected error but got nil")
			}
			if !tt.wantErr && err != nil {
				t.Fatalf("expected nil error but got %v", err)
			}
			if err != nil {
				msg := err.Error()
				for _, want := range tt.containsAll {
					if !strings.Contains(msg, want) {
						t.Fatalf("expected error message to contain %q, got %q", want, msg)
					}
				}
			}
		})
	}
}

func TestValidateStringLength_InvalidTagValue(t *testing.T) {
	type BadTag struct {
		Name string `minStrlen:"abc"`
	}

	err := ValidateStringLength(BadTag{Name: "John"})
	if err == nil {
		t.Fatalf("expected error for invalid tag value")
	}
	msg := err.Error()
	if !strings.Contains(msg, "Name") || !strings.Contains(msg, "invalid") {
		t.Fatalf("unexpected error message: %q", msg)
	}
}
