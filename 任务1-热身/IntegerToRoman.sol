// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

contract IntegerToRoman{
    mapping(uint => bytes1)  private romans;

    constructor() {
        romans[1] = 'I';
        romans[5] = 'V';
        romans[10] = 'X';
        romans[50] = 'L';
        romans[100] = 'C';
        romans[500] = 'D';
        romans[1000] = 'M';
    }

    function integerToRoman(uint number) public view returns (string memory){
        require(number > 0 ,'invalid param');

        bytes memory bytesStr;

        uint i = number / 1000 ;
        if (i > 0){
            for (uint j = 0; j < i; j++){
                bytesStr = abi.encodePacked(bytesStr, romans[1000]);
            }
            number = number % 1000;
        }

        i = number / 100 ;
        if (i == 4){
            bytesStr = abi.encodePacked(bytesStr, 'CD');
            number = number % 100;
        }else if(i ==9){
            bytesStr = abi.encodePacked(bytesStr, 'CM');
            number = number % 100;
        }else if(i > 5){
            bytesStr = abi.encodePacked(bytesStr, romans[500]);
            number = number % 100;
            for (uint j = 0; j < i-5; j++){
                bytesStr = abi.encodePacked(bytesStr, romans[100]);
            }
        }else if(i >0 ){
            for (uint j = 0; j < i; j++){
                bytesStr = abi.encodePacked(bytesStr, romans[100]);
            }
            number = number % 100;
        }

         i = number / 10;
         if (i == 4){
            bytesStr = abi.encodePacked(bytesStr, 'XL');
            number = number % 10;
        }else if(i ==9){
            bytesStr = abi.encodePacked(bytesStr, 'XC');
            number = number % 10;
         }else if(i > 5){
            bytesStr = abi.encodePacked(bytesStr, romans[50]);
            number = number % 10;
            for (uint j = 0; j < i-5; j++){
                bytesStr = abi.encodePacked(bytesStr, romans[10]);
            }
        }else if(i >0 ){
            for (uint j = 0; j < i; j++){
                bytesStr = abi.encodePacked(bytesStr, romans[10]);
            }
            number = number % 10;
        }

         i = number;
         if (i == 4){
            bytesStr = abi.encodePacked(bytesStr, 'IV');
        }else if(i ==9){
            bytesStr = abi.encodePacked(bytesStr, 'IX');
         }else if(i > 5){
            bytesStr = abi.encodePacked(bytesStr, romans[5]);
        }else if(i >0 ){
            for (uint j = 0; j < i; j++){
                bytesStr = abi.encodePacked(bytesStr, romans[1]);
            }
        }

        return string(bytesStr);
    }
}