package fiberhtmx

import (
	"github.com/gofiber/fiber/v2"

	"github.com/stackus/htmx"
)

// Response modifies the fiber.Ctx to add HTMX headers and status codes.
//
// The following options are available:
//   - Status(int) | StatusStopPolling: Sets the HTTP status code of the HTMX response.
//   - Location(path, ...properties): Enables client-side redirection without a full page reload.
//   - PushUrl(string): Pushes a new URL into the history stack.
//   - Redirect(string): Performs a client-side redirect with a full page reload.
//   - Refresh(bool): If set to "true", triggers a full refresh of the client-side page.
//   - ReplaceUrl(string): Replaces the current URL in the location bar.
//   - Reswap(string) | {Swap constants}: Specifies how the response will be swapped.
//   - Retarget(string): A CSS selector to update the target of the content update to a different page element.
//   - Reselect(string): A CSS selector to select a part of the response to be swapped in, overriding existing hx-select on the triggering element.
//   - Trigger(...events): Triggers client-side events.
//   - TriggerAfterSettle(...events): Triggers client-side events after the settle step.
//   - TriggerAfterSwap(...events): Triggers client-side events after the swap step.
func Response(ctx *fiber.Ctx, options ...htmx.ResponseOption) (*htmx.HtmxResponse, error) {
	r, err := htmx.BuildResponse(options...)
	if err != nil {
		return nil, err
	}

	for k, v := range r.Headers() {
		ctx.Set(k, v)
	}

	if r.StatusCode() != 0 {
		ctx.Status(r.StatusCode())
	}

	return r, nil
}
