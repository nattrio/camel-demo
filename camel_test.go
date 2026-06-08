package main

import (
	"reflect"
	"testing"
)

func TestToCamelCase(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "Diagram example",
			input: "Hello world 123-1",
			want:  "helloWorld1231",
		},
		{
			name:  "Snake Case to Camel Case",
			input: "user_profile_data",
			want:  "userProfileData",
		},
		{
			name:  "Kebab Case with Numbers",
			input: "version-2-update",
			want:  "version2Update",
		},
		{
			name:  "Empty String",
			input: "",
			want:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToCamelCase(tt.input); got != tt.want {
				t.Errorf("ToCamelCase() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestRemoveSpecialChar(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"Diagram example", "Hello world 123-1", "Hello world 123 1"},
		{"Multiple dashes", "a---b---c", "a b c"},
		{"No special chars", "hello world", "hello world"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeSpecialChar(tt.input); got != tt.want {
				t.Errorf("removeSpecialChar() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestCapitalizeFromIndex(t *testing.T) {
	tests := []struct {
		name  string
		Words []string
		index int
		want  []string
	}{
		{"Diagram example", []string{"hello", "world", "123", "1"}, 1, []string{"hello", "World", "123", "1"}},
		{"Single word", []string{"hello"}, 1, []string{"hello"}},
		{"Empty slice", []string{}, 1, []string{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := capitalizeWordsFromIndex(tt.Words, tt.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("capitalizeWordsFromIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
