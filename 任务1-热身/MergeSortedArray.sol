// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

contract MergeSortedArray{
    function mergeSortedArray(uint[] memory nums1,uint[]  memory nums2) public pure returns (uint[] memory)  {
        uint[] memory mergedArray = new uint[](nums1.length + nums2.length);
        for (uint i = 0; i< nums1.length ; i++){
            mergedArray[i] = nums1[i];
        }
        for (uint j = 0; j< nums2.length ; j++){
            mergedArray[nums1.length + j] = nums2[j];
        }

        uint temp;
        for (uint i = 0; i< mergedArray.length-1 ; i++){
            for (uint j = nums1.length; j< mergedArray.length ; j++){
                if( i == j){
                    continue ;
                }
                if (mergedArray[i] < mergedArray[j]){
                    break ;
                }else{
                    temp = mergedArray[j];
                    mergedArray[j] = mergedArray[i];
                    mergedArray[i] = temp;
                }
            }
        }
        return mergedArray;

    }
}