package htmx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPushUrl(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		options ResponseOption
		want    map[string]string
	}{
		"Set new path": {
			options: PushUrl("/foo"),
			want: map[string]string{
				HxPushUrl: "/foo",
			},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			o := &HtmxResponse{headers: make(map[string]string)}
			tc.options.apply(o)

			assert.Equal(t, tc.want, o.headers)
		})
	}
}

func TestRedirect(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		options ResponseOption
		want    map[string]string
	}{
		"Set redirect path": {
			options: Redirect("/foo"),
			want: map[string]string{
				HxRedirect: "/foo",
			},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			o := &HtmxResponse{headers: make(map[string]string)}
			tc.options.apply(o)

			assert.Equal(t, tc.want, o.headers)
		})
	}
}

func TestRefresh(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		options ResponseOption
		want    map[string]string
	}{
		"Set refresh": {
			options: Refresh(),
			want: map[string]string{
				HxRefresh: "true",
			},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			o := &HtmxResponse{headers: make(map[string]string)}
			tc.options.apply(o)

			assert.Equal(t, tc.want, o.headers)
		})
	}
}

func TestReplaceUrl(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		options ResponseOption
		want    map[string]string
	}{
		"Set replacement path": {
			options: ReplaceUrl("/foo"),
			want: map[string]string{
				HxReplaceUrl: "/foo",
			},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			o := &HtmxResponse{headers: make(map[string]string)}
			tc.options.apply(o)

			assert.Equal(t, tc.want, o.headers)
		})
	}
}

func TestReselect(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		options ResponseOption
		want    map[string]string
	}{
		"Set new select target": {
			options: Reselect("#foo"),
			want: map[string]string{
				HxReselect: "#foo",
			},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			o := &HtmxResponse{headers: make(map[string]string)}
			tc.options.apply(o)

			assert.Equal(t, tc.want, o.headers)
		})
	}
}

func TestRetarget(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		options ResponseOption
		want    map[string]string
	}{
		"Set new target": {
			options: Retarget("#foo"),
			want: map[string]string{
				HxRetarget: "#foo",
			},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			o := &HtmxResponse{headers: make(map[string]string)}
			tc.options.apply(o)

			assert.Equal(t, tc.want, o.headers)
		})
	}
}

func TestStatus(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		options ResponseOption
		want    int
	}{
		"Set status": {
			options: Status(200),
			want:    200,
		},
		"Set stop polling status": {
			options: StatusStopPolling,
			want:    286,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			o := &HtmxResponse{headers: make(map[string]string)}
			tc.options.apply(o)

			assert.Equal(t, tc.want, o.status)
		})
	}
}
