import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';
import path from 'path'

export default defineConfig(() => {
  return {
    define: {
        'process.env': {},
      },
    build: {
      outDir: 'build',
    },
    plugins: [react({
        babel: {
            plugins: [
              'babel-plugin-macros'
            ]
          }
    })],
    resolve: {
        alias: {
            src: "/src",
      components: "/src/components",
      assets: "/src/assets",
      constants:"/src/constants",
      configs:"/src/configs",
      utils:"/src/utils",
      locales:"/src/locales",
      views:"/src/views",
      store:"/src/store",
      services:"/src/services"
          },
      },
  };
});