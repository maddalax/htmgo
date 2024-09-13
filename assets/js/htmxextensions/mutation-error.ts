import htmx from "htmx.org";

htmx.defineExtension("mutation-error", {
  // @ts-ignore
  onEvent: (name, evt) => {
    if (!(evt instanceof CustomEvent)) {
      return false;
    }
    if (name === "htmx:afterRequest") {
      if (!evt.detail || !evt.detail.xhr) {
        return;
      }
      const status = evt.detail.xhr.status;
      if (status >= 400) {
        htmx.findAll("[hx-on\\:\\:mutation-error]").forEach((element) => {
          htmx.trigger(element, "htmx:mutation-error", { status });
        });
      }
    }
  },
});
