package htmx

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrigger(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		trigger ResponseOption
		want    map[string]any
	}{
		"Set named event": {
			trigger: Trigger(Event("myEvent")),
			want: map[string]any{
				"myEvent": nil,
			},
		},
		"Set event with value": {
			trigger: Trigger(Event("myEvent", "myValue")),
			want: map[string]any{
				"myEvent": "myValue",
			},
		},
		"Set event with multiple values": {
			trigger: Trigger(Event("myEvent", "myValue1", "myValue2")),
			want: map[string]any{
				"myEvent": []any{
					"myValue1",
					"myValue2",
				},
			},
		},
		"Set event with value struct": {
			trigger: Trigger(Event("myEvent",
				struct {
					K string `json:"k"`
				}{
					K: "v",
				},
			)),
			want: map[string]any{
				"myEvent": map[string]any{
					"k": "v",
				},
			},
		},
		"Set event with map": {
			trigger: Trigger(Event("myEvent",
				map[string]string{
					"k": "v",
				},
			)),
			want: map[string]any{
				"myEvent": map[string]any{
					"k": "v",
				},
			},
		},
		"Set multiple events": {
			trigger: Trigger(
				Event("myEvent"),
				Event("myOtherEvent", "myValue"),
				Event("myStructEvent", struct {
					K string `json:"k"`
				}{
					K: "v",
				}),
				Event("myMapEvent", map[string]string{
					"k": "v",
				}),
			),
			want: map[string]any{
				"myEvent":      nil,
				"myOtherEvent": "myValue",
				"myStructEvent": map[string]any{
					"k": "v",
				},
				"myMapEvent": map[string]any{
					"k": "v",
				},
			},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			o := &HtmxResponse{headers: make(map[string]string)}
			tc.trigger.apply(o)

			gotHeader := o.headers[HxTrigger]
			assert.NotEmpty(t, gotHeader)
			got := make(map[string]any)
			assert.NoError(t, json.Unmarshal([]byte(gotHeader), &got))
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestTriggerAfterSettle(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		trigger ResponseOption
		want    map[string]any
	}{
		"Set named event": {
			trigger: TriggerAfterSettle(Event("myEvent")),
			want: map[string]any{
				"myEvent": nil,
			},
		},
		"Set event with value": {
			trigger: TriggerAfterSettle(Event("myEvent", "myValue")),
			want: map[string]any{
				"myEvent": "myValue",
			},
		},
		"Set event with multiple values": {
			trigger: TriggerAfterSettle(Event("myEvent", "myValue1", "myValue2")),
			want: map[string]any{
				"myEvent": []any{
					"myValue1",
					"myValue2",
				},
			},
		},
		"Set event with value struct": {
			trigger: TriggerAfterSettle(Event("myEvent",
				struct {
					K string `json:"k"`
				}{
					K: "v",
				},
			)),
			want: map[string]any{
				"myEvent": map[string]any{
					"k": "v",
				},
			},
		},
		"Set event with map": {
			trigger: TriggerAfterSettle(Event("myEvent",
				map[string]string{
					"k": "v",
				},
			)),
			want: map[string]any{
				"myEvent": map[string]any{
					"k": "v",
				},
			},
		},
		"Set multiple events": {
			trigger: TriggerAfterSettle(
				Event("myEvent"),
				Event("myOtherEvent", "myValue"),
				Event("myStructEvent", struct {
					K string `json:"k"`
				}{
					K: "v",
				}),
				Event("myMapEvent", map[string]string{
					"k": "v",
				}),
			),
			want: map[string]any{
				"myEvent":      nil,
				"myOtherEvent": "myValue",
				"myStructEvent": map[string]any{
					"k": "v",
				},
				"myMapEvent": map[string]any{
					"k": "v",
				},
			},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			o := &HtmxResponse{headers: make(map[string]string)}
			tc.trigger.apply(o)

			gotHeader := o.headers[HxTriggerAfterSettle]
			assert.NotEmpty(t, gotHeader)
			got := make(map[string]any)
			assert.NoError(t, json.Unmarshal([]byte(gotHeader), &got))
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestTriggerAfterSwap(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		trigger ResponseOption
		want    map[string]any
	}{
		"Set named event": {
			trigger: TriggerAfterSwap(Event("myEvent")),
			want: map[string]any{
				"myEvent": nil,
			},
		},
		"Set event with value": {
			trigger: TriggerAfterSwap(Event("myEvent", "myValue")),
			want: map[string]any{
				"myEvent": "myValue",
			},
		},
		"Set event with multiple values": {
			trigger: TriggerAfterSwap(Event("myEvent", "myValue1", "myValue2")),
			want: map[string]any{
				"myEvent": []any{
					"myValue1",
					"myValue2",
				},
			},
		},
		"Set event with value struct": {
			trigger: TriggerAfterSwap(Event("myEvent",
				struct {
					K string `json:"k"`
				}{
					K: "v",
				},
			)),
			want: map[string]any{
				"myEvent": map[string]any{
					"k": "v",
				},
			},
		},
		"Set event with map": {
			trigger: TriggerAfterSwap(Event("myEvent",
				map[string]string{
					"k": "v",
				},
			)),
			want: map[string]any{
				"myEvent": map[string]any{
					"k": "v",
				},
			},
		},
		"Set multiple events": {
			trigger: TriggerAfterSwap(
				Event("myEvent"),
				Event("myOtherEvent", "myValue"),
				Event("myStructEvent", struct {
					K string `json:"k"`
				}{
					K: "v",
				}),
				Event("myMapEvent", map[string]string{
					"k": "v",
				}),
			),
			want: map[string]any{
				"myEvent":      nil,
				"myOtherEvent": "myValue",
				"myStructEvent": map[string]any{
					"k": "v",
				},
				"myMapEvent": map[string]any{
					"k": "v",
				},
			},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			o := &HtmxResponse{headers: make(map[string]string)}
			tc.trigger.apply(o)

			gotHeader := o.headers[HxTriggerAfterSwap]
			assert.NotEmpty(t, gotHeader)
			got := make(map[string]any)
			assert.NoError(t, json.Unmarshal([]byte(gotHeader), &got))
			assert.Equal(t, tc.want, got)
		})
	}
}
