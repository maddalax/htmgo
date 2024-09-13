import htmx from "htmx.org";

function dependsOn(pathSpec: any, url: string) {
  if (pathSpec === "ignore") {
    return false;
  }
  const dependencyPath = pathSpec.split("/");
  const urlPath = url.split("/");
  for (let i = 0; i < urlPath.length; i++) {
    const dependencyElement = dependencyPath.shift();
    const pathElement = urlPath[i];
    if (dependencyElement !== pathElement && dependencyElement !== "*") {
      return false;
    }
    if (
      dependencyPath.length === 0 ||
      (dependencyPath.length === 1 && dependencyPath[0] === "")
    ) {
      return true;
    }
  }
  return false;
}

function refreshPath(path: string) {
  const eltsWithDeps = htmx.findAll("[path-deps]");
  for (let i = 0; i < eltsWithDeps.length; i++) {
    const elt = eltsWithDeps[i];
    if (dependsOn(elt.getAttribute("path-deps"), path)) {
      htmx.trigger(elt, "path-deps", null);
    }
  }
}

htmx.defineExtension("path-deps", {
  // @ts-ignore
  onEvent: function (name, evt) {
    if (!(evt instanceof CustomEvent)) {
      return false;
    }
    if (name === "htmx:beforeOnLoad") {
      const config = evt.detail.requestConfig;
      // mutating call
      if (
        config &&
        config.verb !== "get" &&
        evt.target != null &&
        evt.target instanceof Element &&
        evt.target.getAttribute("path-deps") !== "ignore"
      ) {
        refreshPath(config.path);
      }
    }
  },
});
