import * as htmx from "htmx.org";

htmx.defineExtension("debug", {
  onEvent: function (name, evt) {
    if (console.debug) {
      console.debug(name, evt.target, evt);
    } else if (console) {
      console.log("DEBUG:", name, evt.target, evt);
    } else {
      // noop
    }
  },
});
