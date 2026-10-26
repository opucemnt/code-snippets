// 命令行 Todo 小工具, 数据存在当前目录的 todo.json。
//
// 运行:
//   node todo.js add "买牛奶"
//   node todo.js list
//   node todo.js done 1
const fs = require("fs");
const DB = "./todo.json";

function load() {
  try {
    return JSON.parse(fs.readFileSync(DB, "utf8"));
  } catch {
    return [];
  }
}

function save(items) {
  fs.writeFileSync(DB, JSON.stringify(items, null, 2));
}

const [cmd, ...rest] = process.argv.slice(2);
const items = load();

if (cmd === "add") {
  items.push({ text: rest.join(" "), done: false });
  save(items);
  console.log("已添加");
} else if (cmd === "list") {
  items.forEach((it, i) => {
    console.log(`${i + 1}. [${it.done ? "x" : " "}] ${it.text}`);
  });
} else if (cmd === "done") {
  const idx = parseInt(rest[0], 10) - 1;
  if (items[idx]) {
    items[idx].done = true;
    save(items);
    console.log("已完成");
  } else {
    console.log("编号不存在");
  }
} else {
  console.log("用法: node todo.js <add|list|done> [参数]");
}
