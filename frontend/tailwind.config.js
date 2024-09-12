/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./index.html",
    "./src/**/*.{vue,js,ts,jsx,tsx}",
  ],
  darkMode: 'class',
  theme: {
    extend: {
      colors: {
        'background': 'var(--color-background)',
        'background-soft': 'var(--color-background-soft)',
        'background-mute': 'var(--color-background-mute)',
        'border': 'var(--color-border)',
        'border-hover': 'var(--color-border-hover)',
        'heading': 'var(--color-heading)',
        'text': 'var(--color-text)',
      },
      fontFamily: {
        sans: ['Josefin Sans', 'Inter', '-apple-system', 'BlinkMacSystemFont', 'Segoe UI', 'Roboto', 'Oxygen', 'Ubuntu', 'Cantarell', 'Fira Sans', 'Droid Sans', 'Helvetica Neue', 'sans-serif']
      }
    },
  },
  plugins: [],
}
