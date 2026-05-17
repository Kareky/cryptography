package coordinates_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/Kareky/cryptography/internal/coordinates"
)

func TestNew(t *testing.T) {
	c := coordinates.New(1, 2)
	if c.Row != 1 || c.Col != 2 {
		t.Errorf("New(1,2) = %+v, want Row=1, Col=2", c)
	}
}

func TestCoordinate_ToString(t *testing.T) {
	tests := []struct {
		name string
		c    coordinates.Coordinate
		want string
	}{
		{"one-digit", coordinates.New(1, 2), "12"},
		{"zeroes", coordinates.New(0, 0), "00"},
		{"larger digits", coordinates.New(5, 7), "57"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.c.ToString()
			if got != tt.want {
				t.Errorf("ToString() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestCoordinate_ToInt(t *testing.T) {
	c := coordinates.New(3, 4)
	got := c.ToInt()
	want := []int{3, 4}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("ToInt() = %v, want %v", got, want)
	}
}

func TestToString(t *testing.T) {
	tests := []struct {
		name string
		c    []coordinates.Coordinate
		want string
	}{
		{"empty", []coordinates.Coordinate{}, ""},
		{"single", []coordinates.Coordinate{coordinates.New(1, 2)}, "12"},
		{"multiple", []coordinates.Coordinate{coordinates.New(1, 2), coordinates.New(3, 4)}, "12 34"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := coordinates.ToString(tt.c)
			if got != tt.want {
				t.Errorf("ToString() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestToPairStrings(t *testing.T) {
	c := []coordinates.Coordinate{coordinates.New(1, 2), coordinates.New(3, 4)}
	got := coordinates.ToPairStrings(c)
	want := []string{"12", "34"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("ToPairStrings() = %v, want %v", got, want)
	}
}

func TestToInt(t *testing.T) {
	c := []coordinates.Coordinate{coordinates.New(1, 2), coordinates.New(3, 4)}
	got := coordinates.ToInt(c)
	want := []int{1, 2, 3, 4}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("ToInt() = %v, want %v", got, want)
	}
}

func TestToPairInts(t *testing.T) {
	c := []coordinates.Coordinate{coordinates.New(1, 2), coordinates.New(3, 4)}
	got := coordinates.ToPairInts(c)
	want := [][]int{{1, 2}, {3, 4}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("ToPairs() = %v, want %v", got, want)
	}
}

func TestFromString(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    []coordinates.Coordinate
		wantErr error
	}{
		{"valid single", "12", []coordinates.Coordinate{coordinates.New(1, 2)}, nil},
		{"valid multiple", "12 34", []coordinates.Coordinate{coordinates.New(1, 2), coordinates.New(3, 4)}, nil},
		{"empty string", "", []coordinates.Coordinate{}, nil},
		{"extra space", "  12   34  ", []coordinates.Coordinate{coordinates.New(1, 2), coordinates.New(3, 4)}, nil},
		{"invalid length (single digit)", "1", nil, coordinates.ErrNotAPair},
		{"invalid length (three digits)", "123", nil, coordinates.ErrNotAPair},
		{"non-digit character", "1a", nil, coordinates.ErrNotAnInt},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := coordinates.FromString(tt.input)
			if tt.wantErr != nil {
				if err == nil || !errors.Is(err, tt.wantErr) {
					t.Errorf("FromString() error = %v, want %v", err, tt.wantErr)
				}
				return
			}
			if err != nil {
				t.Errorf("FromString() unexpected error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromStringPairs(t *testing.T) {
	tests := []struct {
		name    string
		input   []string
		want    []coordinates.Coordinate
		wantErr error
	}{
		{"valid", []string{"12", "34"}, []coordinates.Coordinate{coordinates.New(1, 2), coordinates.New(3, 4)}, nil},
		{"invalid length", []string{"1"}, nil, coordinates.ErrNotAPair},
		{"non-digit", []string{"ab"}, nil, coordinates.ErrNotAnInt},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := coordinates.FromStringPairs(tt.input)
			if tt.wantErr != nil {
				if err == nil || !errors.Is(err, tt.wantErr) {
					t.Errorf("FromStringPairs() error = %v, want %v", err, tt.wantErr)
				}
				return
			}
			if err != nil {
				t.Errorf("FromStringPairs() unexpected error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromStringPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromInt(t *testing.T) {
	tests := []struct {
		name    string
		input   []int
		want    []coordinates.Coordinate
		wantErr error
	}{
		{"valid", []int{1, 2, 3, 4}, []coordinates.Coordinate{coordinates.New(1, 2), coordinates.New(3, 4)}, nil},
		{"odd length", []int{1}, nil, coordinates.ErrNotAPair},
		{"empty slice", []int{}, []coordinates.Coordinate{}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := coordinates.FromInt(tt.input)
			if tt.wantErr != nil {
				if err == nil || !errors.Is(err, tt.wantErr) {
					t.Errorf("FromInt() error = %v, want %v", err, tt.wantErr)
				}
				return
			}
			if err != nil {
				t.Errorf("FromInt() unexpected error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromIntPairs(t *testing.T) {
	tests := []struct {
		name    string
		input   [][]int
		want    []coordinates.Coordinate
		wantErr error
	}{
		{"valid", [][]int{{1, 2}, {3, 4}}, []coordinates.Coordinate{coordinates.New(1, 2), coordinates.New(3, 4)}, nil},
		{"invalid inner length", [][]int{{1}}, nil, coordinates.ErrNotAPair},
		{"empty", [][]int{}, []coordinates.Coordinate{}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := coordinates.FromIntPairs(tt.input)
			if tt.wantErr != nil {
				if err == nil || !errors.Is(err, tt.wantErr) {
					t.Errorf("FromIntPairs() error = %v, want %v", err, tt.wantErr)
				}
				return
			}
			if err != nil {
				t.Errorf("FromIntPairs() unexpected error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromIntPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRoundTrip(t *testing.T) {
	original := []coordinates.Coordinate{
		coordinates.New(0, 0),
		coordinates.New(4, 5),
		coordinates.New(9, 2),
	}

	// ToString → FromString
	str := coordinates.ToString(original)
	parsed, err := coordinates.FromString(str)
	if err != nil {
		t.Fatalf("FromString(ToString()) error: %v", err)
	}
	if !reflect.DeepEqual(original, parsed) {
		t.Errorf("ToString → FromString: got %v, want %v", parsed, original)
	}

	// ToInt → FromInt
	ints := coordinates.ToInt(original)
	parsedInt, err := coordinates.FromInt(ints)
	if err != nil {
		t.Fatalf("FromInt(ToInt()) error: %v", err)
	}
	if !reflect.DeepEqual(original, parsedInt) {
		t.Errorf("ToInt → FromInt: got %v, want %v", parsedInt, original)
	}

	// ToPairs → FromIntPairs
	pairs := coordinates.ToPairInts(original)
	parsedPairs, err := coordinates.FromIntPairs(pairs)
	if err != nil {
		t.Fatalf("FromIntPairs(ToPairs()) error: %v", err)
	}
	if !reflect.DeepEqual(original, parsedPairs) {
		t.Errorf("ToPairs → FromIntPairs: got %v, want %v", parsedPairs, original)
	}

	// ToPairStrings → FromStringPairs
	pairStrs := coordinates.ToPairStrings(original)
	parsedStrs, err := coordinates.FromStringPairs(pairStrs)
	if err != nil {
		t.Fatalf("FromStringPairs(ToPairStrings()) error: %v", err)
	}
	if !reflect.DeepEqual(original, parsedStrs) {
		t.Errorf("ToPairStrings → FromStringPairs: got %v, want %v", parsedStrs, original)
	}
}