package main

import (
	"testing"
)

func TestGenerateTemplate(t *testing.T) {
	source := `
K={{ env "ALWAYS_THERE" }}
K={{ env "NONEXISTING" }}
K={{ .NONEXISTING }}
K={{ default .NonExisting "default value" }}
K={{ default (env "ALWAYS_THERE") }}
K={{ default (env "NONEXISTING") "default value" }}
	`
	correctOutput := `
K=always_there
K=
K=<no value>
K=default value
K=always_there
K=default value
	`
	result, err := generateTemplate(source, "test")
	if err != nil {
		t.Fatal(err)
	}

	if result != correctOutput {
		t.Fatalf("Result:\n%s\n==== is not equal to correct template output:\n%s\n", result, correctOutput)
	}
}
