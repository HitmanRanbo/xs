package xs

import "testing"

func TestInvalidUnmarshalError_Error(t *testing.T) {
	tests := []struct {
		name string
		e    *InvalidUnmarshalError
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Error(); got != tt.want {
				t.Errorf("InvalidUnmarshalError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOpenFileError_Error(t *testing.T) {
	tests := []struct {
		name string
		e    *OpenFileError
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Error(); got != tt.want {
				t.Errorf("OpenFileError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLackColError_Error(t *testing.T) {
	tests := []struct {
		name string
		e    *LackColError
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Error(); got != tt.want {
				t.Errorf("LackColError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLackHeaderError_Error(t *testing.T) {
	tests := []struct {
		name string
		e    *LackHeaderError
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Error(); got != tt.want {
				t.Errorf("LackHeaderError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTypeMismatchedError_Error(t *testing.T) {
	tests := []struct {
		name string
		e    *TypeMismatchedError
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Error(); got != tt.want {
				t.Errorf("TypeMismatchedError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}
