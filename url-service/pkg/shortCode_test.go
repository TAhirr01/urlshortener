package pkg

import (
	"strings"
	"testing"
)

// This is a bad example of testing!!
func TestGenerateShortCode(t *testing.T) {
	shortCode := GenerateShortCode(6)
	shortCode2 := GenerateShortCode(6)

	expectedLen := 6
	if shortCode == shortCode2 {
		t.Errorf("ShortCode generator generates same code twice code_1:%s code_2:%s", shortCode, shortCode2)
	}
	if len(shortCode) != expectedLen {
		t.Errorf("Expected len to be %d but is %d", expectedLen, len(shortCode))
	}
}

//Table Driven Test

func TestTableDrivenTestShortCode(t *testing.T) {
	tests := []struct {
		name   string
		length int
	}{
		{"lenght 6", 6},
		{"length 10", 10},
		{"length 1", 1},
		{"length 0", 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			code := GenerateShortCode(tt.length)

			if len(code)!=tt.length{
				t.Fatalf("expected %d, got %d",tt.length,len(code))
			}

			for _,c:=range code{
				if !strings.ContainsRune(charset,c){
					t.Fatalf("invalid character:%c",c)
				}
			}
		})
	}
}
