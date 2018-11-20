# Supplychain for medicine material with nodejs chaincode

1. 進入supplychain/nodejs_version/fabric-samples-release-1.1/balance-transfer/目錄底下
此部分為參考fabric-samples github repo:
[fabric-samples/balancetransfer](https://github.com/hyperledger/fabric-samples/tree/release-1.2/balance-transfer)

2. 輸入./runApp.sh 啟動fabric網絡

3. 輸入./mytestAPIs.sh 來啟動channel及chaincode並產生restful api
使用mytestAPIs.sh是為啟用我寫的chaincode，在/artifacts/src/github.com/example_cc/node2/目錄下

4. 回到medicine目錄下，使用npm run dev指令產生網頁與第3步產生的restful api連結，以與chaincode溝通

5. 這部分還沒實作完成，會遇到一些錯誤
