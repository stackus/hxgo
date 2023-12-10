package htmx

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCurrentUrl(t *testing.T) {
	t.Parallel()

	type args struct {
		r *http.Request
	}
	tests := map[string]struct {
		args args
		want string
	}{
		"Get current Url": {
			args: args{
				r: &http.Request{
					Header: func() http.Header {
						h := http.Header{}
						h.Add(HxCurrentUrl, "/foo")
						return h
					}(),
				},
			},
			want: "/foo",
		},
		"Blank": {
			args: args{
				r: &http.Request{},
			},
			want: "",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equalf(t, tt.want, GetCurrentUrl(tt.args.r), "Headers: %v", tt.args.r.Header)
		})
	}
}

func TestGetPrompt(t *testing.T) {
	t.Parallel()

	type args struct {
		r *http.Request
	}
	tests := map[string]struct {
		args args
		want string
	}{
		"Get prompt": {
			args: args{
				r: &http.Request{
					Header: func() http.Header {
						h := http.Header{}
						h.Add(HxPrompt, "foo")
						return h
					}(),
				},
			},
			want: "foo",
		},
		"Blank": {
			args: args{
				r: &http.Request{},
			},
			want: "",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equalf(t, tt.want, GetPrompt(tt.args.r), "Headers: %v", tt.args.r.Header)
		})
	}
}

func TestGetTarget(t *testing.T) {
	t.Parallel()

	type args struct {
		r *http.Request
	}
	tests := map[string]struct {
		args args
		want string
	}{
		"Get target": {
			args: args{
				r: &http.Request{
					Header: func() http.Header {
						h := http.Header{}
						h.Add(HxTarget, "foo")
						return h
					}(),
				},
			},
			want: "foo",
		},
		"Blank": {
			args: args{
				r: &http.Request{},
			},
			want: "",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equalf(t, tt.want, GetTarget(tt.args.r), "Headers: %v", tt.args.r.Header)
		})
	}
}

func TestGetTrigger(t *testing.T) {
	t.Parallel()

	type args struct {
		r *http.Request
	}
	tests := map[string]struct {
		args args
		want string
	}{
		"Get trigger": {
			args: args{
				r: &http.Request{
					Header: func() http.Header {
						h := http.Header{}
						h.Add(HxTrigger, "#foo")
						return h
					}(),
				},
			},
			want: "#foo",
		},
		"Blank": {
			args: args{
				r: &http.Request{},
			},
			want: "",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equalf(t, tt.want, GetTrigger(tt.args.r), "Headers: %v", tt.args.r.Header)
		})
	}
}

func TestGetTriggerName(t *testing.T) {
	t.Parallel()

	type args struct {
		r *http.Request
	}
	tests := map[string]struct {
		args args
		want string
	}{
		"Get trigger name": {
			args: args{
				r: &http.Request{
					Header: func() http.Header {
						h := http.Header{}
						h.Add(HxTriggerName, "foo")
						return h
					}(),
				},
			},
			want: "foo",
		},
		"Blank": {
			args: args{
				r: &http.Request{},
			},
			want: "",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equalf(t, tt.want, GetTriggerName(tt.args.r), "Headers: %v", tt.args.r.Header)
		})
	}
}

func TestIsBoosted(t *testing.T) {
	t.Parallel()

	type args struct {
		r *http.Request
	}
	tests := map[string]struct {
		args args
		want bool
	}{
		"Get boosted": {
			args: args{
				r: &http.Request{
					Header: func() http.Header {
						h := http.Header{}
						h.Add(HxBoosted, "true")
						return h
					}(),
				},
			},
			want: true,
		},
		"Blank": {
			args: args{
				r: &http.Request{},
			},
			want: false,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equalf(t, tt.want, IsBoosted(tt.args.r), "Headers: %v", tt.args.r.Header)
		})
	}
}

func TestIsHtmx(t *testing.T) {
	t.Parallel()

	type args struct {
		r *http.Request
	}
	tests := map[string]struct {
		args args
		want bool
	}{
		"Get request": {
			args: args{
				r: &http.Request{
					Header: func() http.Header {
						h := http.Header{}
						h.Add(HxRequest, "true")
						return h
					}(),
				},
			},
			want: true,
		},
		"Blank": {
			args: args{
				r: &http.Request{},
			},
			want: false,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equalf(t, tt.want, IsHtmx(tt.args.r), "Headers: %v", tt.args.r.Header)
		})
	}
}

func TestIsHistoryRestoreRequest(t *testing.T) {
	t.Parallel()

	type args struct {
		r *http.Request
	}
	tests := map[string]struct {
		args args
		want bool
	}{
		"Get history restore request": {
			args: args{
				r: &http.Request{
					Header: func() http.Header {
						h := http.Header{}
						h.Add(HxHistoryRestoreRequest, "true")
						return h
					}(),
				},
			},
			want: true,
		},
		"Blank": {
			args: args{
				r: &http.Request{},
			},
			want: false,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equalf(t, tt.want, IsHistoryRestoreRequest(tt.args.r), "Headers: %v", tt.args.r.Header)
		})
	}
}

func TestIsRequest(t *testing.T) {
	t.Parallel()

	type args struct {
		r *http.Request
	}
	tests := map[string]struct {
		args args
		want bool
	}{
		"Get request": {
			args: args{
				r: &http.Request{
					Header: func() http.Header {
						h := http.Header{}
						h.Add(HxRequest, "true")
						return h
					}(),
				},
			},
			want: true,
		},
		"Blank": {
			args: args{
				r: &http.Request{},
			},
			want: false,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equalf(t, tt.want, IsRequest(tt.args.r), "Headers: %v", tt.args.r.Header)
		})
	}
}
