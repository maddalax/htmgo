const {join} = require("node:path");
/** @type {import('tailwindcss').Config} */
const root = join(__dirname, "../../");
const content = join(root, "**/*.go");

console.log(content)

module.exports = {
  content: [content],
  theme: {
    extend: {},
  },
  plugins: [],
};
