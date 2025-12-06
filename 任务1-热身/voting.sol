// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;
contract Voting{

    // 一个mapping来存储候选人的得票数
    mapping (address acount => uint64 votes) private votesReceived;
    address[] private accounts;

    // 一个vote函数，允许用户投票给某个候选人
    function vote(address candidate) public returns (bool ){
        require(candidate != address(0),"param cannot be empty");
        require(msg.sender != candidate,"cannot choose yourself");

        if (0 == votesReceived[candidate]){
            accounts.push(candidate);
        }

        votesReceived [candidate] = votesReceived[candidate] +1;
        return true;
    }
    // 一个getVotes函数，返回某个候选人的得票数
    function getVotes(address candidate) public view returns (uint64){
        require(candidate != address(0),"param cannot be empty");
        return votesReceived[candidate];
    }

    // 一个resetVotes函数，重置所有候选人的得票数
    function resetVotes() public returns (bool){
        for (uint i=0;i< accounts.length; i++) {
            delete votesReceived[accounts[i]];
        }
        return true;
    }

}