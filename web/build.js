const fs = require('fs');
const path = require('path');

const source = path.join(__dirname, 'src', 'index.html');
const outputDir = path.join(__dirname, 'dist');
fs.mkdirSync(outputDir, { recursive: true });
const html = fs.readFileSync(source, 'utf8');
fs.writeFileSync(path.join(outputDir, 'index.html'), html.replace('%%BUILD_MODE%%', 'production'));
process.stdout.write(`built ${path.relative(process.cwd(), path.join(outputDir, 'index.html'))}\n`);
