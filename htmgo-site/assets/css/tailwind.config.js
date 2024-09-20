const {join} = require("node:path");
/** @type {import('tailwindcss').Config} */
const root = join(__dirname, "../../");
const contentGo = join(root, "**/*.go");
const contentJs = join(root, "**/pages/**/*.js");


module.exports = {
  content: [contentGo, contentJs],
  plugins: [ require('@tailwindcss/typography')],
};
