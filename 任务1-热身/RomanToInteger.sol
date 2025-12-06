// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

contract RomanToInteger {
    mapping[string]uint8 romans = {
        "I":1,
        "V":5,
        "X":10,
        "L":50,
        "C":100,
        "D":500,
        "M":1000,
        "IV":4,
        "IX":9,
        "XL":40,
        "XC":90,
        "CD":400,
        "CM":900
    }

    function romanToInteger (string memory str) public pure returns (uint){
        require(str != "" ,"param must not blank");

        uint result;
        if (length(str) == 2){
            result = romans[str];
        }
        
        if(result != 0){
            return result;
        }

        bytes strByte =  str.bytes;
        for(uint i=0;i<strByte.length;i++){
            result +=romans[strByte[i]];
        }
         return result;
    }

}