package hx

import "fmt"

type Modifier interface {
	Modifier() string
}

type RawModifier struct {
	modifier string
}

func StringModifier(modifier string) RawModifier {
	return RawModifier{modifier}
}

func (r RawModifier) Modifier() string {
	return r.modifier
}

type OnceModifier struct{}

func (o OnceModifier) Modifier() string {
	return "once"
}

type ThrottleModifier struct {
	durationSeconds int
}

func (t ThrottleModifier) Modifier() string {
	return fmt.Sprintf("throttle:%ds", t.durationSeconds)
}

func Throttle(durationSeconds int) ThrottleModifier {
	return ThrottleModifier{durationSeconds}
}

type DelayModifier struct {
	durationSeconds int
}

func (t DelayModifier) Modifier() string {
	return fmt.Sprintf("delay:%ds", t.durationSeconds)
}

func Delay(durationSeconds int) DelayModifier {
	return DelayModifier{durationSeconds}
}
