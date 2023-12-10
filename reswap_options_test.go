package htmx

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestReswap_FocusScroll(t *testing.T) {
	t.Parallel()

	type args struct {
		focus bool
	}
	tests := map[string]struct {
		s    Reswap
		args args
		want Reswap
	}{
		"Set focus-scroll": {
			s:    SwapInnerHtml,
			args: args{focus: true},
			want: Reswap("innerHTML focus-scroll:true"),
		},
		"Unset focus-scroll": {
			s:    SwapInnerHtml,
			args: args{focus: false},
			want: Reswap("innerHTML focus-scroll:false"),
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.s.FocusScroll(tt.args.focus), "FocusScroll(%v)", tt.args.focus)
		})
	}
}

func TestReswap_IgnoreTitle(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		s    Reswap
		want Reswap
	}{
		"Set ignoreTitle": {
			s:    SwapInnerHtml,
			want: Reswap("innerHTML ignoreTitle:true"),
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.s.IgnoreTitle(), "IgnoreTitle()")
		})
	}
}

func TestReswap_Scroll(t *testing.T) {
	t.Parallel()

	type args struct {
		target string
	}
	tests := map[string]struct {
		s    Reswap
		args args
		want Reswap
	}{
		"Set scroll": {
			s:    SwapInnerHtml,
			args: args{target: "top"},
			want: Reswap("innerHTML scroll:top"),
		},
		"Set scroll with target": {
			s:    SwapInnerHtml,
			args: args{target: "#another-div:top"},
			want: Reswap("innerHTML scroll:#another-div:top"),
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.s.Scroll(tt.args.target), "Scroll(%v)", tt.args.target)
		})
	}
}

func TestReswap_Settle(t *testing.T) {
	t.Parallel()

	type args struct {
		dur time.Duration
	}
	tests := map[string]struct {
		s    Reswap
		args args
		want Reswap
	}{
		"Set settle with 1s": {
			s:    SwapInnerHtml,
			args: args{dur: 1 * time.Second},
			want: Reswap("innerHTML settle:1s"),
		},
		"Set settle with 1m30s": {
			s:    SwapInnerHtml,
			args: args{dur: 1*time.Minute + 30*time.Second},
			want: Reswap("innerHTML settle:1m30s"),
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.s.Settle(tt.args.dur), "Settle(%v)", tt.args.dur)
		})
	}
}

func TestReswap_Show(t *testing.T) {
	t.Parallel()

	type args struct {
		target string
	}
	tests := map[string]struct {
		s    Reswap
		args args
		want Reswap
	}{
		"Set show": {
			s:    SwapInnerHtml,
			args: args{target: "top"},
			want: Reswap("innerHTML show:top"),
		},
		"Set show with target": {
			s:    SwapInnerHtml,
			args: args{target: "#another-div:top"},
			want: Reswap("innerHTML show:#another-div:top"),
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.s.Show(tt.args.target), "Show(%v)", tt.args.target)
		})
	}
}

func TestReswap_Swap(t *testing.T) {
	t.Parallel()

	type args struct {
		dur time.Duration
	}
	tests := map[string]struct {
		s    Reswap
		args args
		want Reswap
	}{
		"Set swap with 1s": {
			s:    SwapInnerHtml,
			args: args{dur: 1 * time.Second},
			want: Reswap("innerHTML swap:1s"),
		},
		"Set swap with 1m30s": {
			s:    SwapInnerHtml,
			args: args{dur: 1*time.Minute + 30*time.Second},
			want: Reswap("innerHTML swap:1m30s"),
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.s.Swap(tt.args.dur), "Swap(%v)", tt.args.dur)
		})
	}
}

func TestReswap_Transition(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		s    Reswap
		want Reswap
	}{
		"Set transition": {
			s:    SwapInnerHtml,
			want: Reswap("innerHTML transition:true"),
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.s.Transition(), "Transition()")
		})
	}
}

func TestSwapInnerHtml(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		s    Reswap
		want Reswap
	}{
		"Set innerHTML": {
			s:    SwapInnerHtml,
			want: Reswap("innerHTML"),
		},
		"Set innerHTML with focus-scroll": {
			s:    SwapInnerHtml.FocusScroll(true),
			want: Reswap("innerHTML focus-scroll:true"),
		},
		"Set innerHTML with multiple modifiers": {
			s:    SwapInnerHtml.IgnoreTitle().FocusScroll(true).Transition().Settle(1 * time.Second).Swap(1 * time.Second).Show("top").Scroll("top"),
			want: Reswap("innerHTML ignoreTitle:true focus-scroll:true transition:true settle:1s swap:1s show:top scroll:top"),
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.s)
		})
	}
}

func TestSwapOuterHtml(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		s    Reswap
		want Reswap
	}{
		"Set outerHTML": {
			s:    SwapOuterHtml,
			want: Reswap("outerHTML"),
		},
		"Set outerHTML with focus-scroll": {
			s:    SwapOuterHtml.FocusScroll(true),
			want: Reswap("outerHTML focus-scroll:true"),
		},
		"Set outerHTML with multiple modifiers": {
			s:    SwapOuterHtml.IgnoreTitle().FocusScroll(true).Transition().Settle(1 * time.Second).Swap(1 * time.Second).Show("top").Scroll("top"),
			want: Reswap("outerHTML ignoreTitle:true focus-scroll:true transition:true settle:1s swap:1s show:top scroll:top"),
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.s)
		})
	}
}

func TestSwapBeforeBegin(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		s    Reswap
		want Reswap
	}{
		"Set beforebegin": {
			s:    SwapBeforeBegin,
			want: Reswap("beforebegin"),
		},
		"Set beforebegin with focus-scroll": {
			s:    SwapBeforeBegin.FocusScroll(true),
			want: Reswap("beforebegin focus-scroll:true"),
		},
		"Set beforebegin with multiple modifiers": {
			s:    SwapBeforeBegin.IgnoreTitle().FocusScroll(true).Transition().Settle(1 * time.Second).Swap(1 * time.Second).Show("top").Scroll("top"),
			want: Reswap("beforebegin ignoreTitle:true focus-scroll:true transition:true settle:1s swap:1s show:top scroll:top"),
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.s)
		})
	}
}

func TestSwapBeforeEnd(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		s    Reswap
		want Reswap
	}{
		"Set beforeend": {
			s:    SwapBeforeEnd,
			want: Reswap("beforeend"),
		},
		"Set beforeend with focus-scroll": {
			s:    SwapBeforeEnd.FocusScroll(true),
			want: Reswap("beforeend focus-scroll:true"),
		},
		"Set beforeend with multiple modifiers": {
			s:    SwapBeforeEnd.IgnoreTitle().FocusScroll(true).Transition().Settle(1 * time.Second).Swap(1 * time.Second).Show("top").Scroll("top"),
			want: Reswap("beforeend ignoreTitle:true focus-scroll:true transition:true settle:1s swap:1s show:top scroll:top"),
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.s)
		})
	}
}

func TestSwapAfterBegin(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		s    Reswap
		want Reswap
	}{
		"Set afterbegin": {
			s:    SwapAfterBegin,
			want: Reswap("afterbegin"),
		},
		"Set afterbegin with focus-scroll": {
			s:    SwapAfterBegin.FocusScroll(true),
			want: Reswap("afterbegin focus-scroll:true"),
		},
		"Set afterbegin with multiple modifiers": {
			s:    SwapAfterBegin.IgnoreTitle().FocusScroll(true).Transition().Settle(1 * time.Second).Swap(1 * time.Second).Show("top").Scroll("top"),
			want: Reswap("afterbegin ignoreTitle:true focus-scroll:true transition:true settle:1s swap:1s show:top scroll:top"),
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.s)
		})
	}
}

func TestSwapAfterEnd(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		s    Reswap
		want Reswap
	}{
		"Set afterend": {
			s:    SwapAfterEnd,
			want: Reswap("afterend"),
		},
		"Set afterend with focus-scroll": {
			s:    SwapAfterEnd.FocusScroll(true),
			want: Reswap("afterend focus-scroll:true"),
		},
		"Set afterend with multiple modifiers": {
			s:    SwapAfterEnd.IgnoreTitle().FocusScroll(true).Transition().Settle(1 * time.Second).Swap(1 * time.Second).Show("top").Scroll("top"),
			want: Reswap("afterend ignoreTitle:true focus-scroll:true transition:true settle:1s swap:1s show:top scroll:top"),
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.s)
		})
	}
}

func TestSwapDelete(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		s    Reswap
		want Reswap
	}{
		"Set delete": {
			s:    SwapDelete,
			want: Reswap("delete"),
		},
		"Set delete with focus-scroll": {
			s:    SwapDelete.FocusScroll(true),
			want: Reswap("delete focus-scroll:true"),
		},
		"Set delete with multiple modifiers": {
			s:    SwapDelete.IgnoreTitle().FocusScroll(true).Transition().Settle(1 * time.Second).Swap(1 * time.Second).Show("top").Scroll("top"),
			want: Reswap("delete ignoreTitle:true focus-scroll:true transition:true settle:1s swap:1s show:top scroll:top"),
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.s)
		})
	}
}

func TestSwapNone(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		s    Reswap
		want Reswap
	}{
		"Set none": {
			s:    SwapNone,
			want: Reswap("none"),
		},
		"Set none with focus-scroll": {
			s:    SwapNone.FocusScroll(true),
			want: Reswap("none focus-scroll:true"),
		},
		"Set none with multiple modifiers": {
			s:    SwapNone.IgnoreTitle().FocusScroll(true).Transition().Settle(1 * time.Second).Swap(1 * time.Second).Show("top").Scroll("top"),
			want: Reswap("none ignoreTitle:true focus-scroll:true transition:true settle:1s swap:1s show:top scroll:top"),
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.s)
		})
	}
}
