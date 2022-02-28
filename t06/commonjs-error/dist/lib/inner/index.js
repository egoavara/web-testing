"use strict";
//
Object.defineProperty(exports, "[whoami]", {value : "lib/inner/index.js"})

Object.defineProperty(exports, "__esModule", { value: true });
exports.calltest = exports.test = void 0;
function test() {
    return 1;
}
exports.test = test;
var lib_conflict = require("./conflict.js");
Object.defineProperty(exports, "calltest", { enumerable: true, get: function () { return lib_conflict.calltest; } });
