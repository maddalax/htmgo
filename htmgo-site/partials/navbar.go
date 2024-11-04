package partials

import (
	"fmt"
	"github.com/maddalax/htmgo/framework/h"
	"htmgo-site/internal/httpjson"
	"time"
)

type NavItem struct {
	Name string
	Url  string
}

var navItems = []NavItem{
	{Name: "Docs", Url: "/docs"},
	{Name: "Examples", Url: "/examples"},
	{Name: "Convert HTML", Url: "/html-to-go"},
}

type NavBarProps struct {
	Expanded       bool
	ShowPreRelease bool
}

func ToggleNavbar(ctx *h.RequestContext) *h.Partial {
	return h.SwapManyPartial(
		ctx,
		MobileNav(ctx, h.GetQueryParam(ctx, "expanded") == "true"),
	)
}

var CachedStar = h.CachedT(time.Minute*15, func(t *h.RequestContext) *h.Element {
	return Star(t)
})

func Logo() *h.Element {
	return h.Svg(
		h.Attribute("viewBox", "0 0 370 80.8775381264543"),
		h.Class("h-full w-full"),
		h.Tag(
			"G",
			h.Attribute("transform", "matrix(1.276616840702525,0,0,1.276616840702525,-4.447757222875277,-26.431067200135733)"),
			h.Attribute("fill", "#111111"),
			h.Path(
				h.Attribute("xmlns", "http://www.w3.org/2000/svg"),
				h.Attribute("fill", "#111111"),
				h.Attribute("d", "M48.34863,25.46777c-0.23438,0.48438-0.47461,0.99414-0.72461,1.55859  c-3.42871,7.72266-11.42285,25.09375-13.74707,30.13672c-1.5293-1.76367-3.08398-3.55469-4.55859-5.24609L48.34863,25.46777z   M5.83398,68.5127l10.04492-21.2168c0.31445-0.83008,1.5293-3.20117,2.93848-3.20117c0.02832,0,0.05957,0.01172,0.08887,0.01367  c-0.67969,0.76563-1.42188,2.04102-2.00195,4.17578c-1.03125,3.79492-4.48535,16.94922-5.34473,20.22852H5.83398z M13.62695,68.5127  c1.03516-3.94531,4.22461-16.08984,5.20703-19.70313c0.63086-2.32227,1.39551-3.26367,1.83691-3.63477  c0.00684,0.00781,0.01367,0.01172,0.02051,0.01953C23.94629,48.68066,36.72852,63.5127,41.0293,68.5127H13.62695z M43.66699,68.5127  c-1.00293-1.16602-4.45117-5.17773-8.35352-9.68945c1.65039-3.58008,10.4834-22.75195,14.13867-30.98633  c2.32031-5.22852,3.6582-6.17773,4.04297-6.34961c0.82715,0.41797,1.73926,1.29102,2.66992,2.61719  c3.94141,5.61719,23.45703,37.07422,28.00098,44.4082H43.66699z"),
			),
		),
		h.Tag(
			"G",
			h.Attribute("transform", "matrix(4.097970099125154,0,0,4.097970099125154,114.03704346489575,-20.491491909735604)"),
			h.Attribute("fill", "#111111"),
			h.Path(
				h.Attribute("d", "M7.5293 9.766 c2.2461 0 3.5938 1.25 3.5938 3.7598 l0 6.4746 l-2.8223 0 l0 -6.0156 c0 -1.4746 -0.82031 -1.9629 -1.8262 -1.9629 c-1.0449 0 -2.1875 0.51758 -2.207 2.4414 l0 5.5371 l-2.8125 0 l0 -15 l2.8125 0 l0 6.1621 c0.71289 -0.86914 1.8359 -1.3965 3.2617 -1.3965 z M17.568346875 20 c-1.9531 0 -3.0664 -1.1328 -3.0664 -3.1348 l0 -4.7461 l-1.9727 0 l0 -2.1582 l0.63477 0 c1.0645 0 1.6504 -0.41016 1.6504 -1.9141 l0 -1.3281 l2.5391 0 l0 3.2422 l2.0703 0 l0 2.1582 l-2.0703 0 l0 4.4434 c0 0.89844 0.43945 1.2988 1.1621 1.2988 l0.9082 0 l0 2.1387 l-1.8555 0 z M33.496484375 9.766 c2.1484 0 3.5352 1.0938 3.5352 3.1543 l0 7.0801 l-2.8125 0 l0 -6.2793 c0 -1.1816 -0.74219 -1.6992 -1.582 -1.6992 c-1.0059 0 -1.8945 0.57617 -1.8945 2.3145 l0 5.6641 l-2.8418 0 l0 -6.25 c0 -1.2012 -0.72266 -1.7285 -1.6113 -1.7285 c-0.97656 0 -1.8848 0.57617 -1.8848 2.4609 l0 5.5176 l-2.8027 0 l0 -10.039 l2.8027 0 l0 1.1816 c0.66406 -0.88867 1.6797 -1.377 2.9102 -1.377 c1.4551 0 2.5488 0.52734 3.0762 1.5039 c0.70313 -1.0059 1.7773 -1.5039 3.1055 -1.5039 z M46.679646875 9.961 l2.6758 0 l0 9.2871 c0 3.9063 -2.1191 5.4883 -5.3223 5.4883 c-2.8809 0 -4.4434 -1.2109 -5.1758 -3.1152 l2.334 -0.99609 c0.56641 1.2988 1.3867 1.9238 2.7344 1.9238 c1.7773 0 2.6074 -1.1133 2.6074 -3.0957 l0 -1.1719 c-0.58594 0.80078 -1.7383 1.3672 -3.0469 1.3672 c-2.4902 0 -4.5801 -1.9629 -4.5801 -4.9609 c0 -3.0078 2.0996 -4.9219 4.5996 -4.9219 c1.4063 0 2.5586 0.625 3.1055 1.5234 z M44.208946875 17.373 c1.4648 0 2.5977 -1.1914 2.5977 -2.6855 c0 -1.5039 -1.1133 -2.6953 -2.5977 -2.6953 c-1.4746 0 -2.5879 1.1426 -2.5879 2.6953 c0 1.5332 1.1328 2.6855 2.5879 2.6855 z M56.9531125 20.19531 c-3.1934 0 -5.498 -1.9434 -5.498 -5.2246 c0 -3.2617 2.2852 -5.2051 5.498 -5.2051 c3.2324 0 5.5078 1.9434 5.5078 5.2051 c0 3.2813 -2.2852 5.2246 -5.5078 5.2246 z M56.9238125 17.959 c1.6309 0 2.7441 -1.1914 2.7441 -2.9883 s-1.1133 -2.9883 -2.7441 -2.9883 c-1.5723 0 -2.6758 1.1914 -2.6758 2.9883 s1.1035 2.9883 2.6758 2.9883 z"),
			),
		),
	)
}

func Star(ctx *h.RequestContext) *h.Element {

	type Repo struct {
		StarCount int `json:"stargazers_count"`
	}

	fmt.Printf("making github star request\n")
	count := 0
	response, err := httpjson.Get[Repo]("https://api.github.com/repos/maddalax/htmgo")

	if err == nil && response != nil {
		count = response.StarCount
	}

	return h.A(
		h.Href("https://github.com/maddalax/htmgo"),
		h.Target("_blank"),
		h.Class("inline-flex items-center rounded overflow-hidden shadow-sm"),
		h.Div(
			h.Class("flex items-center px-2 py-1 bg-gray-800 text-white text-sm font-semibold hover:bg-gray-700 transition"),
			h.Svg(
				h.Class("w-4 h-4 -mt-0.5 mr-0.5 stroke-current text-white"),
				h.Attribute("xmlns", "http://www.w3.org/2000/svg"),
				h.Attribute("viewBox", "0 0 24 24"),
				h.Attribute("fill", "none"),
				h.Attribute("stroke", "currentColor"),
				h.Attribute("stroke-width", "2"),
				h.Path(
					h.D("M12 17.27l5.18 3.05-1.64-5.68 4.46-3.87-5.88-.5L12 3.5l-2.12 6.77-5.88.5 4.46 3.87-1.64 5.68L12 17.27z"),
				),
			),
			h.Text("Star"),
		),
		h.If(
			count > 0,
			h.Div(
				h.Class("flex items-center px-3 py-1 bg-black text-white text-sm font-semibold"),
				h.Pf("%d", count),
			),
		),
	)
}

func DiscordSvg() *h.Element {
	return h.Svg(
		h.Attribute("viewBox", "0 -28.5 256 256"),
		h.Attribute("xmlns", "http://www.w3.org/2000/svg"),
		h.Attribute("preserveaspectratio", "xMidYMid"),
		h.Attribute("fill", "#000000"),
		h.Tag(
			"g",
			h.Id("SVGRepo_bgCarrier"),
			h.Attribute("stroke-width", "0"),
		),
		h.Tag(
			"g",
			h.Id("SVGRepo_tracerCarrier"),
			h.Attribute("stroke-linecap", "round"),
			h.Attribute("stroke-linejoin", "round"),
		),
		h.Tag(
			"g",
			h.Id("SVGRepo_iconCarrier"),
			h.Tag(
				"g",
				h.Path(
					h.Attribute("d", "M216.856339,16.5966031 C200.285002,8.84328665 182.566144,3.2084988 164.041564,0 C161.766523,4.11318106 159.108624,9.64549908 157.276099,14.0464379 C137.583995,11.0849896 118.072967,11.0849896 98.7430163,14.0464379 C96.9108417,9.64549908 94.1925838,4.11318106 91.8971895,0 C73.3526068,3.2084988 55.6133949,8.86399117 39.0420583,16.6376612 C5.61752293,67.146514 -3.4433191,116.400813 1.08711069,164.955721 C23.2560196,181.510915 44.7403634,191.567697 65.8621325,198.148576 C71.0772151,190.971126 75.7283628,183.341335 79.7352139,175.300261 C72.104019,172.400575 64.7949724,168.822202 57.8887866,164.667963 C59.7209612,163.310589 61.5131304,161.891452 63.2445898,160.431257 C105.36741,180.133187 151.134928,180.133187 192.754523,160.431257 C194.506336,161.891452 196.298154,163.310589 198.110326,164.667963 C191.183787,168.842556 183.854737,172.420929 176.223542,175.320965 C180.230393,183.341335 184.861538,190.991831 190.096624,198.16893 C211.238746,191.588051 232.743023,181.531619 254.911949,164.955721 C260.227747,108.668201 245.831087,59.8662432 216.856339,16.5966031 Z M85.4738752,135.09489 C72.8290281,135.09489 62.4592217,123.290155 62.4592217,108.914901 C62.4592217,94.5396472 72.607595,82.7145587 85.4738752,82.7145587 C98.3405064,82.7145587 108.709962,94.5189427 108.488529,108.914901 C108.508531,123.290155 98.3405064,135.09489 85.4738752,135.09489 Z M170.525237,135.09489 C157.88039,135.09489 147.510584,123.290155 147.510584,108.914901 C147.510584,94.5396472 157.658606,82.7145587 170.525237,82.7145587 C183.391518,82.7145587 193.761324,94.5189427 193.539891,108.914901 C193.539891,123.290155 183.391518,135.09489 170.525237,135.09489 Z"),
					h.Attribute("fill", "#000000"),
					h.Attribute("fill-rule", "nonzero"),
				),
			),
		),
	)
}

func NavBar(ctx *h.RequestContext, props NavBarProps) *h.Element {
	banner := h.If(
		true,
		h.A(
			h.Href("/discord"),
			h.Attribute("target", "_blank"),
			h.Class("bg-blue-200 text-slate-700 text-center p-2 flex items-center justify-center gap-2"),
			h.Div(
				h.Class("w-[30px] h-[30px]"),
				DiscordSvg(),
			),
			h.P(
				h.Text("Join the "),
				h.Span(
					h.Text("discord server"),
					h.Class("underline"),
				),
				h.Text(" for support and the latest news."),
			),
		),
	)
	desktopNav := h.Nav(
		h.Class("hidden sm:block bg-neutral-100 border border-b-slate-300 p-4 md:p-3 max-h-[100vh - 9rem] overflow-y-auto"),
		h.Div(
			h.Class("max-w-[95%] md:max-w-3xl px-4 mx-auto"),
			h.Div(
				h.Class("flex justify-between items-center"),
				h.A(
					h.Href("/"),
					h.Class("mt-1 max-w-[125px]"),
					Logo(),
				),
				h.Div(
					h.Id("search-container"),
				),
				h.Div(
					h.Class("flex gap-4 items-center"),
					h.List(navItems, func(item NavItem, index int) *h.Element {
						return h.Div(
							h.Class("flex items-center"),
							h.A(
								h.Class(""),
								h.Href(item.Url),
								h.Text(item.Name),
							),
						)
					}),
					CachedStar(ctx),
				),
			),
		),
	)

	return h.Div(
		h.Id("navbar"),
		banner,
		MobileNav(ctx, props.Expanded),
		desktopNav,
	)
}

func MobileNav(ctx *h.RequestContext, expanded bool) *h.Element {
	return h.Nav(
		h.Id("mobile-nav"),
		h.Class("block sm:hidden bg-neutral-100 border border-b-slate-300 p-4 md:p-3"),
		h.Div(
			h.Class("max-w-[95%] md:max-w-prose mx-auto"),
			h.Div(
				h.Class("flex justify-between items-center"),
				h.Div(
					h.Class("flex items-center"),
					h.A(
						h.Href("/"),
						h.Class("mt-1 max-w-[125px]"),
						Logo(),
					),
				),
				h.Div(
					h.Class("flex items-center gap-3"),
					h.Div(
						h.Class("mt-1"),
						CachedStar(ctx),
					),
					h.Button(
						h.Boost(),
						h.GetPartialWithQs(
							ToggleNavbar,
							h.NewQs(
								"expanded",
								h.Ternary(expanded, "false", "true"),
								"test",
								"true",
							),
							"click",
						),
						h.AttributePairs(
							"class",
							"text-2xl",
							"aria-expanded",
							h.Ternary(expanded, "true", "false"),
						),
						h.Class("text-2xl"),
						h.UnsafeRaw("&#9776;"),
					),
				),
			),
		),
		h.If(
			expanded,
			h.Div(
				h.Class("mt-2 ml-2 flex flex-col gap-2"),
				h.List(navItems, func(item NavItem, index int) *h.Element {
					return h.Div(
						h.Class("flex items-center"),
						h.A(
							h.Boost(),
							h.Class(""),
							h.Href(item.Url),
							h.Text(item.Name),
						),
					)
				}),
			),
		),
	)
}
