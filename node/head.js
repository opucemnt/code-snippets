// head：打印文件的前 N 行，默认 10 行。
// 用法: node head.js <文件> [-n 行数]
const fs = require("fs");

function main() {
  const args = process.argv.slice(2);
  const nIdx = args.indexOf("-n");
  let n = 10;
  let files = args;
  if (nIdx !== -1) {
    n = parseInt(args[nIdx + 1], 10) || 10;
    files = args.filter((_, i) => i !== nIdx && i !== nIdx + 1);
  }
  if (files.length === 0) {
    console.error("用法: node head.js <文件> [-n 行数]");
    process.exit(1);
  }
  const lines = fs.readFileSync(files[0], "utf8").split("\n");
  console.log(lines.slice(0, n).join("\n"));
}

main();
