package pages

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"mhtml/h"
	"os"
	"time"
)

func IndexPage(c *fiber.Ctx) *h.Page {
	return h.NewPage(h.Html(
		h.Class("bg-background flex flex-col items-center"),
		h.Head(
			h.Link("/public/main.css", "stylesheet"),
			h.Script("/public/mhtml.js"),
			h.Script("/public/scripts/shiki.js"),
		),
		h.Body(
			h.Class("flex flex-col gap-3"),
			// Navbar
			h.Div(
				h.Class("flex justify-between items-center w-full p-6"),
				h.Div(
					h.Class("text-white text-xl font-bold"),
					h.Text("MHTML"),
				),
				h.Div(
					h.Class("flex gap-4"),
					h.A(h.Href("/pricing"), h.Class("text-white"), h.Text("Pricing")),
					h.A(h.Href("/docs"), h.Class("text-white"), h.Text("Docs")),
					h.A(h.Href("/app"), h.Class("text-white"), h.Text("App")),
				),
			),

			// Hero Section
			h.Div(
				h.Class("flex flex-col items-center justify-center gap-6 p-12 bg-background text-center"),
				h.H1(
					h.Class("text-white text-4xl sm:text-5xl font-bold max-w-3xl"),
					h.Text("Go and HTMX: The Simple Stack"),
				),
				h.P(
					h.Class("text-white text-lg sm:text-xl max-w-2xl"),
					h.Text("Combine the simplicity of Go with the power of HTMX for dynamic, JavaScript-light web development."),
				),
				h.A(h.Href("/get-started"),
					h.Class("bg-white text-background px-6 py-3 rounded-md font-semibold mt-4"),
					h.Text("Join the waitlist"),
				),
			),

			// Explore Section
			h.Div(
				h.Class("w-full max-w-4xl"),
				CodeExample(),
			),

			// Footer Section
			h.Div(
				h.Class("flex justify-center items-center py-6"),
				h.Text(fmt.Sprintf("© %d MHTML", time.Now().Year())),
			),
		),
	))
}

func CodeExample() h.Renderable {
	code, err := os.ReadFile("pages/assets/_example.go")
	scriptSrc, err := os.ReadFile("pages/assets/shiki.js")

	if err != nil {
		return h.Pf("Error loading code example")
	}

	fmt.Printf("%s\n", code)

	script := fmt.Sprintf(string(scriptSrc), string(code))

	return h.Div(
		h.Class("text-white rounded-lg"),
		h.Pre(h.Id("foo")),
		h.RawF(`
			<script type="module">
				%s
			</script>
		`, script),
	)
}
