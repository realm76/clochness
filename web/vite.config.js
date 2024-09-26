import { defineConfig } from 'vite';
import path from 'path';

export default defineConfig({
    build: {
        outDir: "../internal/app/assets",
        lib: {
            entry: path.resolve(__dirname, 'src/index.ts'), // Your library entry point
            name: 'Clochness', // Global name for UMD/IIFE builds
            fileName: (format) => `clochness.${format}.js`,
        },
    }
});
