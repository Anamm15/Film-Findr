/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      fontFamily: {
        geist: ['Geist', 'sans-serif'],
      },
      colors: {
        'primary': '#007BFF',
        'secondary': '#6C757D',
        'tertiary': '#FF1493',
        'text': '#121212',
        'background': '#FFFFFF',
      },
      backgroundImage: {
        'gradient-primary': 'linear-gradient(135deg, #007BFF 20%, #8A2BE2 80%)',
        'gradient-secondary': 'linear-gradient(135deg, #007BFF 30%, #FF1493 70%)',
      },
    },
  },
  plugins: [],
}
