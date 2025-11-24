#!/usr/bin/env node

import * as esbuild from 'esbuild';
import { readdirSync, statSync, existsSync } from 'fs';
import { join, extname, basename } from 'path';

const args = process.argv.slice(2);
const target = args.find(arg => !arg.startsWith('--'));
const isWatch = args.includes('--watch');

// Build configuration
const config = {
  datatable: {
    entryPoints: ['js/datatable/src/index.ts'],
    bundle: true,
    format: 'iife',
    globalName: 'DataTable',
    target: 'es2020',
  },
  basecoat: {
    inputDir: 'js/basecoat',
    outputDir: 'dist/basecoat',
    bundle: false,
    format: 'iife',
    target: 'es2020',
  },
  css: {
    inputDir: 'css',
    outputDir: 'dist/css',
  }
};

// Utility function to get JS files in a directory
function getJSFiles(dir) {
  if (!existsSync(dir)) return [];
  
  return readdirSync(dir)
    .filter(file => extname(file) === '.js')
    .map(file => join(dir, file));
}

// Utility function to get CSS files in a directory
function getCSSFiles(dir) {
  if (!existsSync(dir)) return [];
  
  return readdirSync(dir)
    .filter(file => extname(file) === '.css')
    .map(file => join(dir, file));
}

// Process CSS files with Tailwind CLI
async function processCSSFile(inputFile, outputDir) {
  const fileName = basename(inputFile, '.css');
  const fs = await import('fs');
  const path = await import('path');
  const { execSync } = await import('child_process');
  
  try {
    // Ensure output directory exists
    await fs.mkdirSync(outputDir, { recursive: true });
    
    const outputFile = path.join(outputDir, `${fileName}.css`);
    const minifiedOutputFile = path.join(outputDir, `${fileName}.min.css`);
    
    // Build unminified version with Tailwind CLI
    execSync(`pnpm dlx @tailwindcss/cli -i ${inputFile} -o ${outputFile}`, { 
      stdio: 'inherit',
      cwd: process.cwd()
    });
    
    // Build minified version with Tailwind CLI
    execSync(`pnpm dlx @tailwindcss/cli -i ${inputFile} -o ${minifiedOutputFile} --minify`, { 
      stdio: 'inherit',
      cwd: process.cwd()
    });
    
  } catch (error) {
    console.error(`❌ Error processing ${inputFile}:`, error);
    throw error;
  }
}

// Build datatable as both minified and unminified files
async function buildDataTable() {
  console.log('🔨 Building datatable...');
  
  try {
    // Build unminified version
    await esbuild.build({
      ...config.datatable,
      outfile: 'dist/datatable.js',
      minify: false,
      sourcemap: true,
    });
    
    // Build minified version
    await esbuild.build({
      ...config.datatable,
      outfile: 'dist/datatable.min.js',
      minify: true,
      sourcemap: false,
    });
    
    console.log('✅ DataTable built successfully (both versions)');
  } catch (error) {
    console.error('❌ DataTable build failed:', error);
    process.exit(1);
  }
}

// Build basecoat components individually (both minified and unminified)
async function buildBasecoat() {
  console.log('🔨 Building basecoat components...');
  
  const jsFiles = getJSFiles(config.basecoat.inputDir);
  
  if (jsFiles.length === 0) {
    console.log('ℹ️ No JS files found in basecoat directory');
    return;
  }

  try {
    // Build each component individually
    for (const file of jsFiles) {
      const fileName = basename(file, '.js');
      
      // Build unminified version
      await esbuild.build({
        entryPoints: [file],
        outfile: `${config.basecoat.outputDir}/${fileName}.js`,
        bundle: config.basecoat.bundle,
        minify: false,
        format: config.basecoat.format,
        target: config.basecoat.target,
        sourcemap: true,
      });
      
      // Build minified version
      await esbuild.build({
        entryPoints: [file],
        outfile: `${config.basecoat.outputDir}/${fileName}.min.js`,
        bundle: config.basecoat.bundle,
        minify: true,
        format: config.basecoat.format,
        target: config.basecoat.target,
        sourcemap: false,
      });
      
      console.log(`✅ Built ${fileName}.js and ${fileName}.min.js`);
    }
    
    console.log('✅ All basecoat components built successfully');
  } catch (error) {
    console.error('❌ Basecoat build failed:', error);
    process.exit(1);
  }
}

// Build CSS files
async function buildCSS() {
  console.log('🎨 Building CSS files...');
  
  const cssFiles = getCSSFiles(config.css.inputDir);
  
  if (cssFiles.length === 0) {
    console.log('ℹ️ No CSS files found in css directory');
    return;
  }

  try {
    for (const file of cssFiles) {
      const fileName = basename(file, '.css');
      
      // Skip basecoat.css as it's imported by base.css
      if (fileName === 'basecoat') {
        console.log(`ℹ️ Skipping ${fileName}.css (imported by base.css)`);
        continue;
      }
      
      await processCSSFile(file, config.css.outputDir);
      console.log(`✅ Built ${fileName}.css and ${fileName}.min.css`);
    }
    
    console.log('✅ All CSS files built successfully');
  } catch (error) {
    console.error('❌ CSS build failed:', error);
    process.exit(1);
  }
}

// Build all components
async function buildAll() {
  await buildDataTable();
  await buildBasecoat();
  await buildCSS();
}

// Watch mode
async function startWatch() {
  console.log('👀 Starting watch mode...');
  
  // Watch datatable
  const datatableCtx = await esbuild.context({
    ...config.datatable,
    plugins: [{
      name: 'rebuild-notify',
      setup(build) {
        build.onEnd(result => {
          if (result.errors.length === 0) {
            console.log('🔄 DataTable rebuilt');
          } else {
            console.error('❌ DataTable rebuild failed');
          }
        });
      },
    }],
  });
  
  await datatableCtx.watch();
  
  // Watch basecoat files
  const jsFiles = getJSFiles(config.basecoat.inputDir);
  const basecoatContexts = [];
  
  for (const file of jsFiles) {
    const fileName = basename(file, '.js');
    
    const ctx = await esbuild.context({
      entryPoints: [file],
      outfile: `${config.basecoat.outputDir}/${fileName}.min.js`,
      bundle: config.basecoat.bundle,
      minify: config.basecoat.minify,
      format: config.basecoat.format,
      target: config.basecoat.target,
      sourcemap: config.basecoat.sourcemap,
      plugins: [{
        name: 'rebuild-notify',
        setup(build) {
          build.onEnd(result => {
            if (result.errors.length === 0) {
              console.log(`🔄 ${fileName}.min.js rebuilt`);
            } else {
              console.error(`❌ ${fileName}.min.js rebuild failed`);
            }
          });
        },
      }],
    });
    
    await ctx.watch();
    basecoatContexts.push(ctx);
  }
  
  console.log('✅ Watch mode started. Press Ctrl+C to stop.');
  
  // Cleanup on exit
  process.on('SIGINT', async () => {
    console.log('\n🔄 Cleaning up...');
    await datatableCtx.dispose();
    await Promise.all(basecoatContexts.map(ctx => ctx.dispose()));
    process.exit(0);
  });
}

// Main execution
async function main() {
  console.log('📦 JS Components Builder');
  console.log('=======================');
  
  if (isWatch) {
    await startWatch();
  } else if (target === 'datatable') {
    await buildDataTable();
  } else if (target === 'basecoat') {
    await buildBasecoat();
  } else if (target === 'css') {
    await buildCSS();
  } else {
    await buildAll();
  }
}

main().catch(error => {
  console.error('💥 Build failed:', error);
  process.exit(1);
});