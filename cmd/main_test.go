package cmd

import (
	"reflect"
	"testing"
)

func Test_splitIntoChunks(t *testing.T) {
	type args struct {
		s         string
		chunkSize int
	}
	tests := []struct {
		name       string
		args       args
		wantChunks []string
		wantErr    bool
	}{
		{"6-lenght chunks", args{"123456789012", 6}, []string{"123456", "789012"}, false},
		{"6-lenght chunks, incomplete", args{"1234567890", 6}, []string{"123456", "7890"}, false},
		{"3-lenght chunks", args{"123456789012", 3}, []string{"123", "456", "789", "012"}, false},
		{"1-lenght chunks", args{"1234", 1}, []string{"1", "2", "3", "4"}, false},
		{"1-lenght chunks", args{"1234", 99}, []string{"1234"}, false},
		{"empty", args{"", 1}, nil, false},
		{"0-lenght chunks", args{"123456789012", 0}, nil, true},
		{"negative-lenght chunks", args{"123456789012", -1}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotChunks, err := splitIntoChunks(tt.args.s, tt.args.chunkSize)
			if (err != nil) != tt.wantErr {
				t.Errorf("splitIntoChunks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotChunks, tt.wantChunks) {
				t.Errorf("splitIntoChunks() = %v, want %v", gotChunks, tt.wantChunks)
			}
		})
	}
}

func Test_getSlice(t *testing.T) {
	type args struct {
		function   operation
		iterations int
	}
	tests := []struct {
		name    string
		args    args
		wantSeq []int
	}{
		{"6", args{func(x int) int { return 3 * x }, 6}, []int{0, 3, 6, 9, 12, 15}},
		{"3", args{func(x int) int { return 3 * x }, 3}, []int{0, 3, 6}},
		{"1", args{func(x int) int { return 3 * x }, 1}, []int{0}},
		{"0", args{func(x int) int { return 3 * x }, 0}, nil},
		{"-2", args{func(x int) int { return 3 * x }, -2}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSeq := getSlice(tt.args.function, tt.args.iterations); !reflect.DeepEqual(gotSeq, tt.wantSeq) {
				t.Errorf("getSlice() = %v, want %v", gotSeq, tt.wantSeq)
			}
		})
	}
}

func Test_getCVC(t *testing.T) {
	const numRuns = 100
	for i := 0; i < numRuns; i++ {
		result, err := getCVC()
		if err != nil {
			t.Error(err)
		}
		if len(result) != 3 {
			t.Errorf("Lenght of the result is not 3")
		}
	}
}

func Test_getCVCCVCsString(t *testing.T) {
	const numRuns = 10
	for i := range numRuns {
		result, err := getCVCCVCsString(i)
		if err != nil {
			t.Error(err)
		}
		if len(result) != i*3*2 {
			t.Errorf("Lenght of the result is not %d", i*3*2)
		}
	}
}
