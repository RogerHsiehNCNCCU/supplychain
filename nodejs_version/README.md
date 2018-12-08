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

![GITHUB](https://github.com/a037580238/supplychain/blob/master/nodejs_version/home.PNG "Home")

然後要先註冊使用者以及Token才能進入到網頁中，而獲得的Token可以協助確認真實身分，
目前Token的動態替換尚未完成，需要使用手動的方式去程式修改Token才能使用。

![GITHUB](https://github.com/a037580238/supplychain/blob/master/nodejs_version/registerToken.PNG "registerToken")

![GITHUB](https://github.com/a037580238/supplychain/blob/master/nodejs_version/registerUser.PNG "registerUser")

登入完成後就可以使用網頁的功能，首先須InitToken讓Token發送到各個節點，
整體使用情境是中藥材的區塊鏈溯源，從上游的農夫到下游的消費者，讓整體的資訊到可以上傳到區塊鏈上，
完成資訊的透明化，且是可信任的資訊。每個節點(使用者)可以上傳資料，上傳資料可以獲得Token獎勵，
也可以查詢資料，查詢資料需要消耗Token。

![GITHUB](https://github.com/a037580238/supplychain/blob/master/nodejs_version/functions.PNG "functions")
