// 合并 JSON：把多个 JSON 对象文件合并成一个，后面的覆盖前面的。
// 用法: node merge_json.js a.json b.json [-o out.json]
const fs = require("fs");

function main() {
  const args = process.argv.slice(2);
  const outIdx = args.indexOf("-o");
  let output = null;
  let files = args;
  if (outIdx !== -1) {
    output = args[outIdx + 1];
    files = args.filter((_, i) => i !== outIdx && i !== outIdx + 1);
  }
  if (files.length === 0) {
    console.error("用法: node merge_json.js a.json b.json [-o out.json]");
    process.exit(1);
  }
  const merged = {};
  for (const path of files) {
    const data = JSON.parse(fs.readFileSync(path, "utf8"));
    if (typeof data !== "object" || data === null || Array.isArray(data)) {
      console.error(`${path} 不是 JSON 对象，跳过合并`);
      process.exit(1);
    }
    Object.assign(merged, data);
  }
  const text = JSON.stringify(merged, null, 2) + "\n";
  if (output) fs.writeFileSync(output, text);
  else process.stdout.write(text);
}

main();
