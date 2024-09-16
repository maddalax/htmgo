package h

type HxEvent = string

var (
	HxBeforeRequest   HxEvent = "hx-on::before-request"
	HxAfterRequest    HxEvent = "hx-on::after-request"
	HxOnMutationError HxEvent = "hx-on::mutation-error"
	HxOnLoad          HxEvent = "hx-on::load"
	HxOnLoadError     HxEvent = "hx-on::load-error"
	HxRequestTimeout  HxEvent = "hx-on::request-timeout"
	HxTrigger         HxEvent = "hx-on::trigger"
	HxRequestStart    HxEvent = "hx-on::xhr:loadstart"
	HxRequestProgress HxEvent = "hx-on::xhr:progress"
)
