// ROT13：字母表轮换 13 位，加一次是加密，再加一次就解回来了。
// 用法: node rot13.js <文本>
//      echo hello | node rot13.js
const fs = require("fs");

function rot13(text) {
  return text.replace(/[a-zA-Z]/g, (ch) => {
    const base = ch <= "Z" ? 65 : 97;
    return String.fromCharCode(((ch.charCodeAt(0) - base + 13) % 26) + base);
  });
}

async function main() {
  let text;
  if (process.argv.length > 2) {
    text = process.argv.slice(2).join(" ");
  } else {
    text = fs.readFileSync(0, "utf8");
  }
  process.stdout.write(rot13(text));
}

main();
