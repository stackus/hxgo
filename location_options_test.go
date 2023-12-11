package hx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLocation(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		location ResponseOption
		want     map[string]string
	}{
		"Set path": {
			location: Location("/foo"),
			want: map[string]string{
				HxLocation: "/foo",
			},
		},
		"Set value struct": {
			location: Location("/foo",
				Values(struct {
					K string `json:"k"`
				}{
					K: "v",
				}),
			),
			want: map[string]string{
				HxLocation: `{"path":"/foo","values":{"k":"v"}}`,
			},
		},
		"Set triggering event": {
			location: Location("/foo",
				EventName("myEvent"),
			),
			want: map[string]string{
				HxLocation: `{"path":"/foo","event":"myEvent"}`,
			},
		},
		"Set swap": {
			location: Location("/foo",
				Swap("innerHTML"),
			),
			want: map[string]string{
				HxLocation: `{"path":"/foo","swap":"innerHTML"}`,
			},
		},
		"Set swap with constant": {
			location: Location("/foo",
				Swap(SwapOuterHtml.FocusScroll(true)),
			),
			want: map[string]string{
				HxLocation: `{"path":"/foo","swap":"outerHTML focus-scroll:true"}`,
			},
		},
		"Set them all": {
			location: Location("/foo",
				Source("bar"),
				EventName("click"),
				Handler("baz"),
				Target("#fizz"),
				Swap("#buzz"),
				Values(map[string]string{"k": "v"}),
				Headers(map[string]string{"H": "h"}),
				Select("#foobar"),
			),
			want: map[string]string{
				HxLocation: `{"path":"/foo","source":"bar","event":"click","handler":"baz","target":"#fizz","swap":"#buzz","values":{"k":"v"},"headers":{"H":"h"},"select":"#foobar"}`,
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			o := &HtmxResponse{headers: make(map[string]string)}
			tc.location.apply(o)

			assert.Equal(t, tc.want, o.headers)
		})
	}
}
