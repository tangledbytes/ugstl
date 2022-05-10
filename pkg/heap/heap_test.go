package heap

import (
	"testing"

	"github.com/utkarsh-pro/ugstl/pkg/types"
	"github.com/utkarsh-pro/ugstl/pkg/utils"
)

func TestNew(t *testing.T) {
	type args struct {
		typ  Type
		less types.Less[int]
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "New must panic when type is neither MIN nor MAX",
			args: args{
				typ:  Type(2),
				less: func(v1, v2 int) bool { return v1 < v2 },
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			utils.ShouldPanic(t, func() { New(tt.args.typ, tt.args.less) })
		})
	}
}
