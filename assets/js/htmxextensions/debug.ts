import htmx from "htmx.org";

htmx.defineExtension("debug", {
  // @ts-ignore
  onEvent: function (name, evt) {
    if (console.debug) {
      console.debug(name);
    } else if (console) {
      console.log("DEBUG:", name);
    } else {
      // noop
    }
  },
});
