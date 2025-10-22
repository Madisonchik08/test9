package main

// Пишите тесты в этом файле
import "testing"

func TestGenerateRandomElement(t *testing.T) {
	tests := []struct {
		name    string
		size    int
		wantLen int
	}{
		{"positive size", 10, 10},
		{"negative size", -5, 0},
		{"zero size", 0, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateRandomElements(tt.size)
			if len(got) != tt.wantLen {
				t.Errorf("generateRandomElements() = %v, want %v", len(got), tt.wantLen)
			}
			if tt.wantLen > 0 {
				for _, v := range got {
					if v <= 0 {
						t.Errorf("generateRandomElements() = %v, want > 0", v)
					}
				}
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"empty slice", []int{}, 0},
		{"single element", []int{5}, 5},
		{"multiple elements", []int{1, 5, 2, 8, 3}, 8},
		{"negative numbers", []int{-1, -5, -2, -8, -3}, -1},
		{"mixed numbers", []int{-1, 5, -2, 8, 3}, 8},
		{"all same", []int{7, 7, 7, 7}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximum(tt.data)
			if got != tt.want {
				t.Errorf("maximum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxChunks(t *testing.T) {
	largeSlice := generateRandomElements(100000)
	expectedMaxLargeSlice := maximum(largeSlice)

	tests := []struct {
		name string
		data []int
		want int
	}{
		{"empty slice", []int{}, 0},
		{"single element", []int{5}, 5},
		{"slice smaller than chunks", []int{1, 8, 3}, 8},
		{"multiple elements", []int{1, 5, 2, 8, 3, 9, 4, 6, 10, 0, 7}, 10},
		{"large slice", largeSlice, expectedMaxLargeSlice},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxChunks(tt.data)
			if got != tt.want {
				t.Errorf("maxChunks() = %v, want %v", got, tt.want)
			}
		})
	}
}
