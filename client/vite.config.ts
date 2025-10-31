import { defineConfig } from 'vite';
import solid from 'vite-plugin-solid';
import tailwindcss from '@tailwindcss/vite';

export default defineConfig({
    plugins: [tailwindcss(), solid()],
    server: {
        proxy: {
            '/graphql': {
                target: 'http://localhost:5000',
            },
        },
        port: 3000,
        open: true,
    },
    resolve: {
        alias: [
            {
                find: /^highlight.js\/lib\/languages\/(.*)/,
                replacement: 'highlight.js/lib/languages/$1',
            },
        ],
    },
});
