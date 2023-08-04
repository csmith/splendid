import adapter from '@sveltejs/adapter-node';

/** @type {import('@sveltejs/kit').Config} */
const config = {
    kit: {
        adapter: adapter(),
        files: {
            assets: 'src/frontend/static',
            routes: 'src/frontend/routes',
            appTemplate: 'src/frontend/app.html'
        },
    },
};

export default config;