// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

// 题目描述：反转一个字符串。输入 "abcde"，输出 "edcba"

contract ReverseString{

    function reverse (string memory param) public pure returns (string memory) {
        bytes memory bytesString = bytes(param);
        uint length = bytesString.length;
        bytes memory outputBytes = new bytes(length);
        for (uint i = 0; i < length; i++) {
            outputBytes[i] = bytesString[length - i-1];
        }
        return string(outputBytes);
    }
}