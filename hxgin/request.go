package hxgin

import (
	"github.com/gin-gonic/gin"

	"github.com/stackus/hxgo"
)

// IsBoosted checks the HX-Boosted header
//
// Returns true if the request is a boosted request
func IsBoosted(ctx *gin.Context) bool {
	return ctx.GetHeader(hx.HxBoosted) == "true"
}

// GetCurrentUrl extracts the HX-Current-URL header from an HTTP request.
//
// It returns the current URL of the browser if the header exists.
// If the header is not present, it returns an empty string.
func GetCurrentUrl(ctx *gin.Context) string {
	return ctx.GetHeader(hx.HxCurrentUrl)
}

// IsHistoryRestoreRequest determines if an HTTP request is a history restore request.
//
// It checks the presence of the HX-History-Restore-Request header in the request.
// Returns true if the header is present, otherwise returns false.
func IsHistoryRestoreRequest(ctx *gin.Context) bool {
	return ctx.GetHeader(hx.HxHistoryRestoreRequest) != ""
}

// GetPrompt extracts the HX-Prompt header from an HTTP request.
//
// It returns the user response to an Hx-Prompt if the header exists.
// If the header is not present, it returns an empty string.
func GetPrompt(ctx *gin.Context) string {
	return ctx.GetHeader(hx.HxPrompt)
}

// IsRequest determines if an HTTP request is an HTMX request.
//
// It checks the presence of the HX-Request header in the request.
// Returns true if the header is present, otherwise returns false.
func IsRequest(ctx *gin.Context) bool {
	return ctx.GetHeader(hx.HxRequest) != ""
}

// IsHtmx determines if an HTTP request is an HTMX request.
//
// Does the same thing as IsRequest, only with a more user-friendly name.
func IsHtmx(ctx *gin.Context) bool {
	return IsRequest(ctx)
}

// GetTarget extracts the HX-Target header from an HTTP request.
//
// It returns the ID of the target element if the header exists.
// If the header is not present, it returns an empty string.
func GetTarget(ctx *gin.Context) string {
	return ctx.GetHeader(hx.HxTarget)
}

// GetTriggerName extracts the HX-Trigger-Name header from an HTTP request.
//
// It returns the name of the triggered element if the header exists.
// If the header is not present, it returns an empty string.
func GetTriggerName(ctx *gin.Context) string {
	return ctx.GetHeader(hx.HxTriggerName)
}

// GetTrigger extracts the HX-Trigger header from an HTTP request.
//
// It returns the ID of the trigger element if the header exists.
// If the header is not present, it returns an empty string.
func GetTrigger(ctx *gin.Context) string {
	return ctx.GetHeader(hx.HxTrigger)
}
