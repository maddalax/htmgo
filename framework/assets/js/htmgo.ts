import htmx from "htmx.org";
import "./htmxextensions/pathdeps";
import "./htmxextensions/trigger-children";
import "./htmxextensions/debug";
import "./htmxextensions/response-targets";
import "./htmxextensions/mutation-error";
import "./htmxextensions/livereload"
import "./htmxextensions/htmgo";
import "./htmxextensions/sse"

// @ts-ignore
window.htmx = htmx;

function watchUrl(callback: (oldUrl: string, newUrl: string) => void) {
  let lastUrl = window.location.href;
  setInterval(() => {
    if (window.location.href !== lastUrl) {
      callback(lastUrl, window.location.href);
      lastUrl = window.location.href;
    }
  }, 101);
}

watchUrl((_, newUrl) => {
  onUrlChange(newUrl);
});

function onUrlChange(newUrl: string) {
  let url = new URL(newUrl);

  document.querySelectorAll("[hx-trigger]").forEach(function (element) {
    const triggers = element.getAttribute("hx-trigger");
    if (!triggers) {
      return;
    }
    const split = triggers.split(", ");
    if (split.find((s) => s === "url")) {
      htmx.swap(element, "url", {
        swapStyle: "outerHTML",
        swapDelay: 0,
        settleDelay: 0,
      });
    } else {
      for (let [key, values] of url.searchParams) {
        let eventName = "qs:" + key;
        if (triggers.includes(eventName)) {
          console.log("triggering", eventName);
          htmx.trigger(element, eventName, null);
          break;
        }
      }
    }
  });

  document.querySelectorAll("[hx-match-qp]").forEach((el) => {
    let hasMatch = false;
    for (let name of el.getAttributeNames()) {
      if (name.startsWith("hx-match-qp-mapping:")) {
        let match = name.replace("hx-match-qp-mapping:", "");
        let value = url.searchParams.get(match);
        if (value) {
          htmx.swap(el, el.getAttribute(name) ?? "", {
            swapStyle: "innerHTML",
            swapDelay: 0,
            settleDelay: 0,
          });
          hasMatch = true;
          break;
        }
      }
    }
    if (!hasMatch) {
      let defaultKey = el.getAttribute("hx-match-qp-default");
      if (defaultKey) {
        htmx.swap(
          el,
          el.getAttribute("hx-match-qp-mapping:" + defaultKey) ?? "",
          { swapStyle: "innerHTML", swapDelay: 0, settleDelay: 0 },
        );
      }
    }
  });
}
