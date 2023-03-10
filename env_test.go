package env

import (
	"os"
	"testing"
)

type MyStruct struct {
	Field1 string `env:"FIELD_1"`
	Field2 int    `env:"FIELD_2"`
}

type MainStruct struct {
	MyStruct `obj:"ref"`
}

func TestPopulateFromEnv(t *testing.T) {
	// Set up test environment variables
	os.Setenv("FIELD_1", "test value 1")
	os.Setenv("FIELD_2", "5")

	s := MyStruct{}
	err := PopulateWithEnv(&s)

	if err != nil {
		t.Error("Unexpected error, ", err)
	}

	// Assert that the fields of the struct were correctly populated
	if s.Field1 != "test value 1" {
		t.Error("Expected Field1 to be 'test value 1', got ", s.Field1)
	}
	if s.Field2 != 5 {
		t.Error("Expected Field2 to be 5, got ", s.Field2)
	}

	// Unset the environment variables after test
	os.Unsetenv("FIELD_1")
	os.Unsetenv("FIELD_2")
}

func TestPopulateStructWithinStructs(t *testing.T) {
	os.Setenv("FIELD_1", "test value 1")
	os.Setenv("FIELD_2", "5")

	s := MainStruct{}
	err := PopulateWithEnv(&s)

	if err != nil {
		t.Error("Unexpected error, ", err.Error())
	}

	if s.Field1 != "test value 1" {
		t.Error("Expected Field1 to be 'test value 1', got ", s.Field1)
	}

	if s.Field2 != 5 {
		t.Error("Expected Field2 to be 5, got ", s.Field2)
	}

	os.Unsetenv("FIELD_1")
	os.Unsetenv("FIELD_2")
}

func TestErrorHandling(t *testing.T) {
	os.Setenv("FIELD_1", "1234")
	s := MainStruct{}
	err := PopulateWithEnv(&s)

	if err == nil {
		t.Error("Expected error, ", err)
	}
	os.Unsetenv("FIELD_1")
}
