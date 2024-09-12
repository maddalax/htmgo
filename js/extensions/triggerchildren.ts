import * as htmx from "htmx.org";
import { HtmxEvent } from "htmx.org";

htmx.defineExtension("trigger-children", {
  onEvent: (name: HtmxEvent, evt: CustomEvent) => {
    const target = evt.detail.target as HTMLElement;
    if (target && target.children) {
      Array.from(target.children).forEach((e) => {
        htmx.trigger(e, name, null);
      });
    }
  },
});
