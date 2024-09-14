import { defineConfig } from "tsup";

export default defineConfig({
  format: ["esm"],
  entry: ["htmgo.ts"],
  outDir: "./../dist",
  dts: false,
  shims: true,
  skipNodeModulesBundle: true,
  clean: false,
  target: "esnext",
  treeshake: true,
  platform: "browser",
  outExtension: () => {
    return {
      js: ".js",
    };
  },
  minify: false,
  bundle: true,
  splitting: true,
  // https://github.com/egoist/tsup/issues/619
  noExternal: [/(.*)/],
});
