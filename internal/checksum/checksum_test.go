package checksum

import (
	_ "embed"
	"testing"
)

func TestFileDigest(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "a",
			args: args{
				path: "testdata/a",
			},
			want: "0cc175b9c0f1b6a831c399e269772661",
		},
		{
			name: "a_newline",
			args: args{
				path: "testdata/a_newline",
			},
			want: "60b725f10c9c85c70d97880dfe8191b3",
		},
		{
			name: "random1e5",
			args: args{
				path: "testdata/random1e5",
			},
			want: "10f3b215c30252b86f1e09d5ae6af3a4",
		},
		{
			name: "random1e6",
			args: args{
				path: "testdata/random1e6",
			},
			want: "6bdddc7d121e9e12cc416068d2bc1559",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FileDigest(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("FileDigest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FileDigest() got = %v, want %v", got, tt.want)
			}
		})
	}
}
