// 随机密码生成器：默认 16 位，可指定长度。
//
// 运行: node password.js [长度，默认 16]
const crypto = require("crypto");

const length = parseInt(process.argv[2], 10) || 16;
const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*";
let out = "";
for (let i = 0; i < length; i++) {
  out += charset[crypto.randomInt(charset.length)];
}
console.log(out);
