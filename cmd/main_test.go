package cmd

import (
	"reflect"
	"strings"
	"testing"
)

func TestSplitIntoSets(t *testing.T) {
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
			gotChunks, err := splitIntoSets(tt.args.s, tt.args.chunkSize)
			if (err != nil) != tt.wantErr {
				t.Errorf("splitIntoSets() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotChunks, tt.wantChunks) {
				t.Errorf("splitIntoSets() = %v, want %v", gotChunks, tt.wantChunks)
			}
		})
	}
}

func TestGetSlice(t *testing.T) {
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

func TestGetCVC(t *testing.T) {
	const numRuns = 100
	for i := 0; i < numRuns; i++ {
		result, err := getCVC(config{})
		if err != nil {
			t.Error(err)
		}
		if len(result) != CHUNKSIZE {
			t.Errorf("length of the result is not %d", CHUNKSIZE)
		}
	}
}

func TestGetCVCCVCsString(t *testing.T) {
	const numRuns = 10
	for i := range numRuns {
		result, err := getCVCCVCsString(config{setsNum: i})
		if err != nil {
			t.Error(err)
		}
		if len(result) != i*CHUNKSIZE*CHUNKSPERSET {
			t.Errorf("length of the result is not %d", i*CHUNKSIZE*CHUNKSPERSET)
		}
	}
}

func TestConfigValidate(t *testing.T) {
	tests := []struct {
		name    string
		cfg     config
		wantErr bool
	}{
		{"valid", config{setsNum: 4, upperNum: 2, digitsNum: 2}, false},
		{"sets too small", config{setsNum: 0}, true},
		{"upper too large", config{setsNum: 1, upperNum: 3}, true},
		{"digits too large", config{setsNum: 1, digitsNum: 3}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.cfg.validate()
			if (err != nil) != tt.wantErr {
				t.Fatalf("validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGeneratePassword(t *testing.T) {
	cfg := config{
		setsNum:       3,
		upperNum:      2,
		digitsNum:     2,
		separator:     "-",
		lessNonPolish: true,
	}

	password, err := GeneratePassword(cfg)
	if err != nil {
		t.Fatalf("GeneratePassword() error = %v", err)
	}

	parts := strings.Split(password, cfg.separator)
	if len(parts) != cfg.setsNum {
		t.Fatalf("GeneratePassword() produced %d parts, want %d", len(parts), cfg.setsNum)
	}

	for _, part := range parts {
		if len(part) != CHUNKSIZE*CHUNKSPERSET {
			t.Fatalf("set length = %d, want %d", len(part), CHUNKSIZE*CHUNKSPERSET)
		}
	}
}
