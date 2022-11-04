package test

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestEnvLoad(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		t.Error(err.Error())
	}

	t.Log(os.Getenv("AA_AA"))

}
