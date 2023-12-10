package htmx

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResponse(t *testing.T) {
	t.Parallel()

	type args struct {
		options []responseOption
	}
	tests := map[string]struct {
		args        args
		wantHeaders http.Header
		wantStatus  int
		wantErr     error
	}{
		"Set status": {
			args: args{
				options: []responseOption{
					Status(http.StatusAccepted),
				},
			},
			wantHeaders: http.Header{},
			wantStatus:  http.StatusAccepted,
		},
		"Set stop polling": {
			args: args{
				options: []responseOption{
					StatusStopPolling,
				},
			},
			wantHeaders: http.Header{},
			wantStatus:  int(StatusStopPolling),
		},
		"Set headers": {
			args: args{
				options: []responseOption{
					Location("/foo"),
					SwapOuterHtml.FocusScroll(true),
				},
			},
			wantHeaders: http.Header{
				HxLocation: []string{`/foo`},
				HxReswap:   []string{`outerHTML focus-scroll:true`},
			},
			wantStatus: http.StatusOK,
		},
		"Overwrite headers": {
			args: args{
				options: []responseOption{
					Location("/foo"),
					SwapOuterHtml.FocusScroll(true),
					Location("/bar"),
				},
			},
			wantHeaders: http.Header{
				HxLocation: []string{`/bar`},
				HxReswap:   []string{`outerHTML focus-scroll:true`},
			},
			wantStatus: http.StatusOK,
		},
		"Set status and headers": {
			args: args{
				options: []responseOption{
					Status(http.StatusAccepted),
					Location("/foo"),
					SwapOuterHtml.FocusScroll(false),
				},
			},
			wantHeaders: http.Header{
				HxLocation: []string{`/foo`},
				HxReswap:   []string{`outerHTML focus-scroll:false`},
			},
			wantStatus: http.StatusAccepted,
		},
		"Panics and recovers": {
			args: args{
				options: []responseOption{
					Location("/foo"),
					Trigger(func() map[string]any {
						panic(fmt.Errorf("bad event data"))
					}),
				},
			},
			wantErr: fmt.Errorf("bad event data"),
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			wr := httptest.NewRecorder()

			err := Response(wr, tt.args.options...)

			if tt.wantErr != nil {
				assert.EqualError(t, err, tt.wantErr.Error())
				return
			}
			gotHeaders := wr.Header()
			assert.Equal(t, tt.wantHeaders, gotHeaders)
			gotStatus := wr.Code
			assert.Equal(t, tt.wantStatus, gotStatus)
		})
	}
}
