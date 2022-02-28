"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.calltest = exports.test = void 0;
function test() {
    return 1;
}
exports.test = test;
var conflict_js_1 = require("./conflict.js");
Object.defineProperty(exports, "calltest", { enumerable: true, get: function () { return conflict_js_1.calltest; } });
