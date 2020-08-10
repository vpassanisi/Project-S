const materialPalette = require("./materialPalette.js");

module.exports = {
  corePlugins: {
    backgroundOpacity: false,
    borderOpacity: false,
  },
  purge: [
    "./src/**/*.html",
    "./src/**/*.vue",
    "./src/**/*.js",
    "./src/**/*.svelte",
    "./src/**/*.jsx",
    "./public/*.html",
  ],
  theme: {
    boxShadow: {
      default: "0px 4px 8px 0px rgba(0,0,0,0.10)",
    },
    colors: materialPalette,
    extend: {
      transitionProperty: {
        height: "height",
      },
    },
  },
  variants: {
    backgroundColor: ["dark", "hover", "focus"],
    textColor: ["dark"],
    placeholderColor: ["dark-placeholder"],
    borderColor: ["dark", "hover", "focus"],
  },
  plugins: [require("tailwindcss-dark-mode")()],
};
