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
        'red-light': 'var(--system-red-light)',
        'red-dark': 'var(--system-red-dark)',
        'orange-light': 'var(--system-orange-light)',
        'orange-dark': 'var(--system-orange-dark)',
        'yellow-light': 'var(--system-yellow-light)',
        'yellow-dark': 'var(--system-yellow-dark)',
        'green-light': 'var(--system-green-light)',
        'green-dark': 'var(--system-green-dark)',
        'mint-light': 'var(--system-mint-light)',
        'mint-dark': 'var(--system-mint-dark)',
        'teal-light': 'var(--system-teal-light)',
        'teal-dark': 'var(--system-teal-dark)',
        'cyan-light': 'var(--system-cyan-light)',
        'cyan-dark': 'var(--system-cyan-dark)',
        'blue-light': 'var(--system-blue-light)',
        'blue-dark': 'var(--system-blue-dark)',
        'indigo-light': 'var(--system-indigo-light)',
        'indigo-dark': 'var(--system-indigo-dark)',
        'purple-light': 'var(--system-purple-light)',
        'purple-dark': 'var(--system-purple-dark)',
        'pink-light': 'var(--system-pink-light)',
        'pink-dark': 'var(--system-pink-dark)',
        'brown-light': 'var(--system-brown-light)',
        'brown-dark': 'var(--system-brown-dark)',
        'black-light': 'var(--system-black-light)',
        'black-dark': 'var(--system-black-dark)',
        'grey-light': 'var(--system-grey-light)',
        'grey-dark': 'var(--system-grey-dark)',
        'grey2-light': 'var(--system-grey2-light)',
        'grey2-dark': 'var(--system-grey2-dark)',
        'grey3-light': 'var(--system-grey3-light)',
        'grey3-dark': 'var(--system-grey3-dark)',
        'grey4-light': 'var(--system-grey4-light)',
        'grey4-dark': 'var(--system-grey4-dark)',
        'grey5-light': 'var(--system-grey5-light)',
        'grey5-dark': 'var(--system-grey5-dark)',
        'grey6-light': 'var(--system-grey6-light)',
        'grey6-dark': 'var(--system-grey6-dark)',
        'white-light': 'var(--system-white-light)',
        'white-dark': 'var(--system-white-dark)',

        // Background Colors
        'bg-primary-light': 'var(--background-primary-light)',
        'bg-secondary-light': 'var(--background-secondary-light)',
        'bg-tertiary-light': 'var(--background-tertiary-light)',
        'bg-primary-dark': 'var(--background-primary-dark)',
        'bg-secondary-dark': 'var(--background-secondary-dark)',
        'bg-tertiary-dark': 'var(--background-tertiary-dark)',

        // Text Colors
        'primary-light': 'var(--text-primary-light)',
        'secondary-light': 'var(--text-secondary-light)',
        'tertiary-light': 'var(--text-tertiary-light)',
        'quartenary-light': 'var(--text-quartenary-light)',
        'primary-dark': 'var(--text-primary-dark)',
        'secondary-dark': 'var(--text-secondary-dark)',
        'tertiary-dark': 'var(--text-tertiary-dark)',
        'quartenary-dark': 'var(--text-quartenary-dark)',

        // input
        'input': 'var(--input-light)',
      },
      fontFamily: {
        sfpro: ['SF Pro Text', 'sans-serif'],
        sans: ['Josefin Sans', 'Inter', '-apple-system', 'BlinkMacSystemFont', 'Segoe UI', 'Roboto', 'Oxygen', 'Ubuntu', 'Cantarell', 'Fira Sans', 'Droid Sans', 'Helvetica Neue', 'sans-serif']
      }
    },
  },
  plugins: [],
}
