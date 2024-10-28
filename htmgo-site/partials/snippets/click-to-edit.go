package snippets

import (
	"fmt"
	"github.com/maddalax/htmgo/framework/h"
)

// RowClasses defined here for simplicity of the example
var RowClasses = "whitespace-nowrap px-4 py-4 font-medium text-gray-900 text-left"
var ButtonClasses = "inline-block rounded bg-slate-900 px-4 py-2 text-xs font-medium text-white hover:bg-slate-800"
var InputClasses = "-ml-2 max-w-[125px] border p-2 rounded focus:outline-none focus:ring focus:ring-slate-800"

type Record struct {
	Id       string
	Name     string
	Birthday string
	Role     string
	Salary   string
}

var records = []Record{
	{
		Id:       "1",
		Name:     "John Doe",
		Birthday: "24/05/1995",
		Role:     "htmgo developer",
		Salary:   "$250,000",
	},
	{
		Id:       "2",
		Name:     "Jake Smith",
		Birthday: "24/05/1995",
		Role:     "htmx developer",
		Salary:   "$100,000",
	},
}

func ClickToEdit(ctx *h.RequestContext) *h.Partial {
	return h.NewPartial(
		h.Div(
			h.Class("flex gap-2 items-center w-full"),
			Table(),
		),
	)
}

// StartEditing is a partial that is called when the user clicks on the edit button,
// it will swap in the form for editing for the given record
func StartEditing(ctx *h.RequestContext) *h.Partial {
	id := ctx.QueryParam("id")

	record := h.Find(records, func(record *Record) bool {
		return record.Id == id
	})

	if record == nil {
		return h.EmptyPartial()
	}

	return h.SwapManyPartial(
		ctx,
		TableRow(record, true),
	)
}

// SaveEditing is a partial that is called when the user clicks on the save button while editing,
// it will update the record with the new values and swap it back out
// note: in the example, we are just creating a new record in memory instead of updating the existing one,
// normally you would update a persistent record in a database
func SaveEditing(ctx *h.RequestContext) *h.Partial {
	id := ctx.QueryParam("id")

	// just for the example, create  a new record so it doesn't affect the global original
	record := Record{
		Id:       id,
		Name:     ctx.FormValue("name"),
		Birthday: ctx.FormValue("birthday"),
		Role:     ctx.FormValue("role"),
		Salary:   ctx.FormValue("salary"),
	}

	return h.SwapPartial(ctx, TableRow(&record, false))
}

func Table() *h.Element {
	return h.Div(
		h.Class("overflow-x-auto w-full"),
		h.Table(
			h.Class("divide-y divide-gray-200 bg-white table-fixed"),
			h.THead(
				h.Tr(
					h.Th(
						h.Class(RowClasses),
						h.Text("Name"),
					),
					h.Th(
						h.Class(RowClasses),
						h.Text("Date of Birth"),
					),
					h.Th(
						h.Class(RowClasses),
						h.Text("Role"),
					),
					h.Th(
						h.Class(RowClasses),
						h.Text("Salary"),
					),
					h.Th(
						h.Class("px-4 py-2"),
					),
				),
			),
			h.TBody(
				h.Class("divide-y divide-gray-200"),
				h.List(records, func(record Record, index int) *h.Element {
					return TableRow(&record, false)
				}),
			),
		),
	)
}

func TableRow(record *Record, editing bool) *h.Element {
	recordId := fmt.Sprintf("record-%s", record.Id)

	var Cell = func(name string, value string) *h.Element {
		return h.Td(
			h.Class(RowClasses, "h-[75px]"),
			h.IfElse(
				!editing,
				h.Pf(
					value,
					h.Class("w-[125px]"),
				),
				h.Input(
					"text",
					h.Class(InputClasses),
					h.Value(value),
					h.Name(name),
				),
			),
		)
	}

	return h.Tr(
		h.If(
			editing,
			// this is important to make sure the inputs are included in the form submission
			h.HxInclude("input"),
		),
		h.Id(recordId),
		Cell("name", record.Name),
		Cell("birthday", record.Birthday),
		Cell("role", record.Role),
		Cell("salary", record.Salary),
		// Edit button
		h.Td(
			h.Button(
				h.Class(ButtonClasses),
				h.PostPartialWithQs(
					h.Ternary(!editing, StartEditing, SaveEditing),
					h.NewQs("id", record.Id),
				),
				h.Text(
					h.Ternary(!editing, "Edit", "Save"),
				),
			),
		),
	)
}
