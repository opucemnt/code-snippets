// 简单的 HTTP 服务, 返回 JSON。
//
// 运行: node server.js
// 访问: curl http://localhost:3000/
const http = require("http");

const server = http.createServer((req, res) => {
  res.writeHead(200, { "Content-Type": "application/json; charset=utf-8" });
  res.end(JSON.stringify({ message: "hello", path: req.url }));
});

server.listen(3000, () => {
  console.log("listening on http://localhost:3000");
});
