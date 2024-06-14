/** @type {import('tailwindcss').Config} */

const defaultTheme = require('tailwindcss/defaultTheme')
const colors = require('tailwindcss/colors')

module.exports = {
    darkMode: 'class',
    content: [
        "./web/**/*.html",
		"./**/*.go",
    ],
    theme: {
        extend: {
            fontFamily: {
                sans: ['Noto Sans', ...defaultTheme.fontFamily.sans],
                mono: ['Noto Sans Mono', 'monospace']
            },
            borderWidth: {
                DEFAULT: '1px',
                '1': '1px',
            },
			gridTemplateColumns: {
				'layout': 'min-content 1fr',
			},
        },
        colors: {
            transparent: 'transparent',
            current: 'currentColor',
            black: colors.black,
            white: colors.white,
            rose: colors.rose,
            pink: colors.pink,
            fuchsia: colors.fuchsia,
            purple: colors.purple,
            violet: colors.violet,
            indigo: colors.indigo,
            blue: colors.blue,
            sky: colors.sky, // As of Tailwind CSS v2.2, `lightBlue` has been renamed to `sky`
            cyan: colors.cyan,
            teal: colors.teal,
            emerald: colors.emerald,
            green: colors.green,
            lime: colors.lime,
            yellow: colors.yellow,
            amber: colors.amber,
            orange: colors.orange,
            red: colors.red,
            slate: colors.slate,
            zinc: colors.zinc,
            gray: colors.zinc,
            neutral: colors.blueGray,
            stone: colors.stone,
        }
    },
    plugins: [
    ],

}
