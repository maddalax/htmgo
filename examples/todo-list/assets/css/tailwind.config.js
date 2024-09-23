const {join} = require("node:path");
/** @type {import('tailwindcss').Config} */
const root = join(__dirname, "../../");
const contentGo = join(root, "**/*.go");
const contentJs = join(root, "**/pages/**/*.js");


module.exports = {
  content: [contentGo, contentJs],
  theme: {
    extend: {
      colors: {
        background: 'hsl(224, 71.4%, 4.1%)',
        foreground: 'hsl(0, 0%, 89%)',
        card: 'hsl(224, 71.4%, 4.1%)',
        cardForeground: 'hsl(0, 0%, 89%)',
        popover: 'hsl(224, 71.4%, 4.1%)',
        popoverForeground: 'hsl(0, 0%, 89%)',
        primary: 'hsl(0, 0%, 89%)',
        primaryForeground: 'hsl(220.9, 39.3%, 11%)',
        secondary: 'hsl(215, 27.9%, 16.9%)',
        secondaryForeground: 'hsl(0, 0%, 89%)',
        muted: 'hsl(215, 27.9%, 16.9%)',
        mutedForeground: 'hsl(217.9, 10.6%, 64.9%)',
        accent: 'hsl(215, 27.9%, 16.9%)',
        accentForeground: 'hsl(0, 0%, 89%)',
        destructive: 'hsl(0, 62.8%, 30.6%)',
        destructiveForeground: 'hsl(0, 0%, 89%)',
        border: 'hsl(215, 27.9%, 16.9%)',
        input: 'hsl(215, 27.9%, 16.9%)',
        ring: 'hsl(216, 12.2%, 83.9%)',
      },
    },
  },
  plugins: [],
};
