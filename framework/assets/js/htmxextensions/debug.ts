import htmx from "htmx.org";

htmx.defineExtension("debug", {
  // @ts-ignore
  onEvent: function (name, evt) {
    if (console.debug) {
      console.debug(name, evt);
    } else if (console) {
      console.log("DEBUG:", name, evt);
    } else {
      // noop
    }
  },
});
