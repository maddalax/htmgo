### Converting Raw HTML to Go

In some cases, you may want to convert raw HTML to Go code. 
A tool to do this automatically is in progress, but in the meantime, I've had success using an LLM such as ChatGPT
to convert HTML to Go code.

Steps:
1. Open an LLM such as ChatGPT
2. Use this prompt:
```text
Using htmgo go html builder, please convert and html below that I send to htmgo code.

The builder looks like this:

package pages

import (
	"github.com/maddalax/htmgo/framework/h"
	"github.com/maddalax/htmgo/framework/hx"
	"github.com/maddalax/htmgo/framework/js"
)

func IndexPage(ctx *h.RequestContext) *h.Page {
	return h.NewPage(
		RootPage(
			h.Div(
				h.JoinExtensions(
					h.TriggerChildren(),
					h.HxExtension("ws"),
				),
				h.Attribute("ws-connect", "/chat"),
				h.Class("flex flex-col gap-4 items-center pt-24 min-h-screen bg-neutral-100"),
				Form(ctx),
				h.Div(
					h.Div(
						h.Id("messages"),
						h.Class("flex flex-col gap-2 w-full"),
					),
				),
			),
		),
	)
}

func MessageInput() *h.Element {
	return h.Input("text",
		h.Id("message-input"),
		h.Required(),
		h.Class("p-4 rounded-md border border-slate-200"),
		h.Name("message"),
		h.Placeholder("Message"),
		h.HxBeforeWsSend(
			js.SetValue(""),
		),
		h.OnEvent(hx.KeyDownEvent, js.SubmitFormOnEnter()),
	)
}

func Form(ctx *h.RequestContext) *h.Element {
	return h.Div(
		h.Class("flex flex-col items-center justify-center p-4 gap-6"),
		h.H2F("Form submission with ws example", h.Class("text-2xl font-bold")),
        h.P(h.Text("example test")),
		h.Form(
			h.Attribute("ws-send", ""),
			h.Class("flex flex-col gap-2"),
			h.LabelFor("name", "Your Message"),
			MessageInput(),
			SubmitButton(),
		),
	)
}

Imagine it has every attribute and every html tag at your disposal. Make things as reusable as possible by separating elements into reusable components.
```

3. Paste the HTML you want to convert to Go code

4. Here is the result from pasting in example HTML from tailwind components:
```html
<div class="relative isolate overflow-hidden bg-gray-900 py-24 sm:py-32">
  <img src="https://images.unsplash.com/photo-1521737604893-d14cc237f11d?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&crop=focalpoint&fp-y=.8&w=2830&h=1500&q=80&blend=111827&sat=-100&exp=15&blend-mode=multiply" alt="" class="absolute inset-0 -z-10 h-full w-full object-cover object-right md:object-center">
  <div class="hidden sm:absolute sm:-top-10 sm:right-1/2 sm:-z-10 sm:mr-10 sm:block sm:transform-gpu sm:blur-3xl" aria-hidden="true">
    <div class="aspect-[1097/845] w-[68.5625rem] bg-gradient-to-tr from-[#ff4694] to-[#776fff] opacity-20" style="clip-path: polygon(74.1% 44.1%, 100% 61.6%, 97.5% 26.9%, 85.5% 0.1%, 80.7% 2%, 72.5% 32.5%, 60.2% 62.4%, 52.4% 68.1%, 47.5% 58.3%, 45.2% 34.5%, 27.5% 76.7%, 0.1% 64.9%, 17.9% 100%, 27.6% 76.8%, 76.1% 97.7%, 74.1% 44.1%)"></div>
  </div>
  <div class="absolute -top-52 left-1/2 -z-10 -translate-x-1/2 transform-gpu blur-3xl sm:top-[-28rem] sm:ml-16 sm:translate-x-0 sm:transform-gpu" aria-hidden="true">
    <div class="aspect-[1097/845] w-[68.5625rem] bg-gradient-to-tr from-[#ff4694] to-[#776fff] opacity-20" style="clip-path: polygon(74.1% 44.1%, 100% 61.6%, 97.5% 26.9%, 85.5% 0.1%, 80.7% 2%, 72.5% 32.5%, 60.2% 62.4%, 52.4% 68.1%, 47.5% 58.3%, 45.2% 34.5%, 27.5% 76.7%, 0.1% 64.9%, 17.9% 100%, 27.6% 76.8%, 76.1% 97.7%, 74.1% 44.1%)"></div>
  </div>
  <div class="mx-auto max-w-7xl px-6 lg:px-8">
    <div class="mx-auto max-w-2xl lg:mx-0">
      <h2 class="text-4xl font-bold tracking-tight text-white sm:text-6xl">Work with us</h2>
      <p class="mt-6 text-lg leading-8 text-gray-300">Anim aute id magna aliqua ad ad non deserunt sunt. Qui irure qui lorem cupidatat commodo. Elit sunt amet fugiat veniam occaecat fugiat aliqua.</p>
    </div>
    <div class="mx-auto mt-10 max-w-2xl lg:mx-0 lg:max-w-none">
      <div class="grid grid-cols-1 gap-x-8 gap-y-6 text-base font-semibold leading-7 text-white sm:grid-cols-2 md:flex lg:gap-x-10">
        <a href="#">Open roles <span aria-hidden="true">&rarr;</span></a>
        <a href="#">Internship program <span aria-hidden="true">&rarr;</span></a>
        <a href="#">Our values <span aria-hidden="true">&rarr;</span></a>
        <a href="#">Meet our leadership <span aria-hidden="true">&rarr;</span></a>
      </div>
      <dl class="mt-16 grid grid-cols-1 gap-8 sm:mt-20 sm:grid-cols-2 lg:grid-cols-4">
        <div class="flex flex-col-reverse">
          <dt class="text-base leading-7 text-gray-300">Offices worldwide</dt>
          <dd class="text-2xl font-bold leading-9 tracking-tight text-white">12</dd>
        </div>
        <div class="flex flex-col-reverse">
          <dt class="text-base leading-7 text-gray-300">Full-time colleagues</dt>
          <dd class="text-2xl font-bold leading-9 tracking-tight text-white">300+</dd>
        </div>
        <div class="flex flex-col-reverse">
          <dt class="text-base leading-7 text-gray-300">Hours per week</dt>
          <dd class="text-2xl font-bold leading-9 tracking-tight text-white">40</dd>
        </div>
        <div class="flex flex-col-reverse">
          <dt class="text-base leading-7 text-gray-300">Paid time off</dt>
          <dd class="text-2xl font-bold leading-9 tracking-tight text-white">Unlimited</dd>
        </div>
      </dl>
    </div>
  </div>
</div>
```

```go
package pages

import (
	"github.com/maddalax/htmgo/framework/h"
	"github.com/maddalax/htmgo/framework/hx"
)

func WorkWithUsSection(ctx *h.RequestContext) *h.Element {
	return h.Div(
		h.Class("relative isolate overflow-hidden bg-gray-900 py-24 sm:py-32"),
		BackgroundImage(),
		HiddenBlurredElement(),
		TopBlurredElement(),
		ContentContainer(ctx),
	)
}

func BackgroundImage() *h.Element {
	return h.Img(
		h.Src("https://images.unsplash.com/photo-1521737604893-d14cc237f11d?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&crop=focalpoint&fp-y=.8&w=2830&h=1500&q=80&blend=111827&sat=-100&exp=15&blend-mode=multiply"),
		h.Alt(""),
		h.Class("absolute inset-0 -z-10 h-full w-full object-cover object-right md:object-center"),
	)
}

func HiddenBlurredElement() *h.Element {
	return h.Div(
		h.Class("hidden sm:absolute sm:-top-10 sm:right-1/2 sm:-z-10 sm:mr-10 sm:block sm:transform-gpu sm:blur-3xl"),
		h.AriaHidden(true),
		BlurredGradient(),
	)
}

func TopBlurredElement() *h.Element {
	return h.Div(
		h.Class("absolute -top-52 left-1/2 -z-10 -translate-x-1/2 transform-gpu blur-3xl sm:top-[-28rem] sm:ml-16 sm:translate-x-0 sm:transform-gpu"),
		h.AriaHidden(true),
		BlurredGradient(),
	)
}

func BlurredGradient() *h.Element {
	return h.Div(
		h.Class("aspect-[1097/845] w-[68.5625rem] bg-gradient-to-tr from-[#ff4694] to-[#776fff] opacity-20"),
		h.Style("clip-path: polygon(74.1% 44.1%, 100% 61.6%, 97.5% 26.9%, 85.5% 0.1%, 80.7% 2%, 72.5% 32.5%, 60.2% 62.4%, 52.4% 68.1%, 47.5% 58.3%, 45.2% 34.5%, 27.5% 76.7%, 0.1% 64.9%, 17.9% 100%, 27.6% 76.8%, 76.1% 97.7%, 74.1% 44.1%)"),
	)
}

func ContentContainer(ctx *h.RequestContext) *h.Element {
	return h.Div(
		h.Class("mx-auto max-w-7xl px-6 lg:px-8"),
		ContentHeader(),
		LinksAndStatsSection(),
	)
}

func ContentHeader() *h.Element {
	return h.Div(
		h.Class("mx-auto max-w-2xl lg:mx-0"),
		h.H2F("Work with us", h.Class("text-4xl font-bold tracking-tight text-white sm:text-6xl")),
		h.P(
			h.Class("mt-6 text-lg leading-8 text-gray-300"),
			h.Text("Anim aute id magna aliqua ad ad non deserunt sunt. Qui irure qui lorem cupidatat commodo. Elit sunt amet fugiat veniam occaecat fugiat aliqua."),
		),
	)
}

func LinksAndStatsSection() *h.Element {
	return h.Div(
		h.Class("mx-auto mt-10 max-w-2xl lg:mx-0 lg:max-w-none"),
		LinksGrid(),
		StatsGrid(),
	)
}

func LinksGrid() *h.Element {
	return h.Div(
		h.Class("grid grid-cols-1 gap-x-8 gap-y-6 text-base font-semibold leading-7 text-white sm:grid-cols-2 md:flex lg:gap-x-10"),
		Link("Open roles"),
		Link("Internship program"),
		Link("Our values"),
		Link("Meet our leadership"),
	)
}

func Link(text string) *h.Element {
	return h.A(
		h.Href("#"),
		h.Text(text),
		h.Span(h.Text(" →"), h.AriaHidden(true)),
	)
}

func StatsGrid() *h.Element {
	return h.Dl(
		h.Class("mt-16 grid grid-cols-1 gap-8 sm:mt-20 sm:grid-cols-2 lg:grid-cols-4"),
		Stat("Offices worldwide", "12"),
		Stat("Full-time colleagues", "300+"),
		Stat("Hours per week", "40"),
		Stat("Paid time off", "Unlimited"),
	)
}

func Stat(label, value string) *h.Element {
	return h.Div(
		h.Class("flex flex-col-reverse"),
		h.Dt(
			h.Class("text-base leading-7 text-gray-300"),
			h.Text(label),
		),
		h.Dd(
			h.Class("text-2xl font-bold leading-9 tracking-tight text-white"),
			h.Text(value),
		),
	)
}
```
