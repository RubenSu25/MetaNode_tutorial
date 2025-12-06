// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

contract BinarySearch{

    function binarySearch(uint[] memory arr, uint target) public pure returns (uint){
        if(arr[0] == target){
            return 1;
        }
        if(arr[arr.length-1] == target){
            return arr.length;
        }
        uint i = arr.length / 2;
        while (i >0) {
            if (arr[i] > target){
                i = i / 2;
            }else if(arr[i] < target){
                i += (arr.length -i )/ 2;
            }else{
                return i+1;
            }
        }
        return i+1;
    }

}