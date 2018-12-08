# Supplychain for medicine material with nodejs chaincode

注: 進入到此project之前，需先做好Fabric的基本環境，請參考Fabric官網

1. 進入supplychain/nodejs_version/fabric-1.2/balance-transfer/目錄底下
此部分為參考fabric-samples github repo:
[fabric-samples/balancetransfer](https://github.com/hyperledger/fabric-samples/tree/release-1.2/balance-transfer)
注: 最新版本的Fabric目前是1.3，預計不久後將更新至1.3版本

2. 開啟Terminal，輸入./runApp.sh 啟動fabric網絡

3. 開啟另一個Terminal，輸入./mytestAPIs.sh -node 來啟動channel及chaincode並產生restful api
使用mytestAPIs.sh是為啟用我寫的chaincode，在/artifacts/src/github.com/example_cc/node2/ 目錄下的unity.js

4. 開啟另一個Terminal，回到medicine目錄下，使用npm run dev指令產生網頁與第3步產生的restful api連結，以與chaincode溝通

5. 成功產生網頁在IP下的8000 Port


