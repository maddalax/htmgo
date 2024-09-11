package main

import "mhtml/h"

func init() {
	h.Div(
		h.Id("sandbox"),
		h.Button(
			h.Id("btn"),
			h.Class("bg-blue-500 text-white p-2"),
		),
	)

	h.Html(
		h.Head(),
		h.Body(
			h.Nav(
				h.Class("flex gap-4 items-center p-4 text-slate-600 "),
				h.A(h.Href("/"), h.Class("cursor-pointer hover:text-blue-400 ")),
				h.A(h.Class("cursor-pointer hover:text-blue-400 "),
					h.Href("/news")), h.A(
					h.Href("/patients"),
					h.Class("cursor-pointer hover:text-blue-400 "))),
			h.Div(
				h.Id("active-modal")),
			h.Div(h.Class("flex flex-col gap-2 bg-white h-full "),
				h.Div(h.Class("flex flex-col p-4 w-full "),
					h.Div(h.Div(h.Class("flex justify-between items-center "),
						h.P(
							h.Class("text-lg font-bold ")),
						h.Button(h.HxTarget("#active-modal"),
							h.Type("button"),
							h.Id("add-patient"),
							h.Class("flex gap-1 items-center border p-4 rounded cursor-hover bg-blue-700 text-white rounded p-2 h-12 "),
							h.HxGet("mhtml/partials/patient.AddPatientSheet"))),
						h.Div(h.HxGet("mhtml/partials/patient.List"),
							h.HxTrigger("load, patient-added from:body"),
							h.Class(""), h.Div(h.Class("mt-8"),
								h.Id("patient-list"),
								h.Div(h.Class("flex flex-col gap-2 rounded p-4 bg-red-100 "),
									h.P(),
									h.P()),
							),
						),
					),
				),
			),
			h.Div(
				h.HxGet("/livereload"),
				h.HxTrigger("every 200ms"),
				h.Class(""),
			),
		))

}
