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
        document.querySelectorAll("*").forEach((element) => {
          if (element.hasAttribute("hx-on::on-mutation-error")) {
            htmx.trigger(element, "htmx:on-mutation-error", { status });
          }
        });
      }
    }
  },
});
