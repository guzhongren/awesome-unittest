package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf(`got %s, want %s`, got, want)
		}
	}
	type input struct {
		language string
		somebody string
	}
	tests := []struct {
		input
		expect string
	}{
		{input{``, ``},
			`Hello, world`,
		},
		{input{``, `li`},
			`Hello, li`,
		},
		{input{`chinese`, `李`},
			`你好, 李`,
		},
		{input{`english`, `li`},
			`Hello, li`,
		},
		{input{`spanish`, `li`},
			`Hola, li`,
		},
	}

	for _, tt := range tests {
		t.Run(`all`, func(t *testing.T) {
			actual := Hello(tt.input.language, tt.input.somebody)
			expect := tt.expect
			assertCorrectMessage(t, actual, expect)
		})
	}
}
