# Supplychain for medicine material with nodejs chaincode

annotation: Before you start the project, you must prepare the environment of Fabric. Please reference Hyperledger Fabric document.

1. Change your directory to supplychain/nodejs_version/fabric-1.2/balance-transfer/  
This is referenced to [fabric-samples/balancetransfer](https://github.com/hyperledger/fabric-samples/tree/release-1.2/balance-transfer)  
annotation: The latest version of Fabric is 1.4. We will update this project to 1.4 soon.

2. Turn on terminal. Enter the instruction ./runApp.sh to start the shell file to start fabric network.

3. Turn on another terminal. Enter ./mytestAPIs.sh -node to start channel, chaincode and restful api.
The shell file mytestAPIs.sh is rewrited to start my chaincode that is unity.js under /artifacts/src/github.com/example_cc/node2/.

4. Turn on another terminal and back to medicine directory. Enter instruction "npm run dev" to make a web which is connecting to the restful api we had started. So the web can communicate to the chaincode.

5. Start a web under your Ip at port 8000 successfully.

注: 進入到此project之前，需先做好Fabric的基本環境，請參考Fabric官網

1. 進入supplychain/nodejs_version/fabric-1.2/balance-transfer/目錄底下
此部分為參考fabric-samples github repo:
[fabric-samples/balancetransfer](https://github.com/hyperledger/fabric-samples/tree/release-1.2/balance-transfer)
注: 最新版本的Fabric目前是1.4，預計不久後將更新至1.4版本

2. 開啟Terminal，輸入./runApp.sh 啟動fabric網絡

3. 開啟另一個Terminal，輸入./mytestAPIs.sh -node 來啟動channel及chaincode並產生restful api
使用mytestAPIs.sh是為啟用我寫的chaincode，在/artifacts/src/github.com/example_cc/node2/ 目錄下的unity.js

4. 開啟另一個Terminal，回到medicine目錄下，使用npm run dev指令產生網頁與第3步產生的restful api連結，以與chaincode溝通

5. 成功產生網頁在IP下的8000 Port

![GITHUB](https://github.com/a037580238/supplychain/blob/master/nodejs_version/home.PNG "Home")

In order to enter the web, we have to register user and Token first. The token we get can verify our identity.
But the dynamic token is not yet accomplished. We have to modify our program manually, so we can enter the web.

然後要先註冊使用者以及Token才能進入到網頁中，而獲得的Token可以協助確認真實身分，
目前Token的動態替換尚未完成，需要使用手動的方式去程式修改Token才能使用。

![GITHUB](https://github.com/a037580238/supplychain/blob/master/nodejs_version/registerToken.PNG "registerToken")

![GITHUB](https://github.com/a037580238/supplychain/blob/master/nodejs_version/registerUser.PNG "registerUser")

After we login, we can use the functions of the web. At first we have to use InitToken let Token sent to every node.
This is a POC of medicine material supply chain with blockchain. Let whole information upload to blockchain from farmer to consumer.
Accomplish information transparent and trustable. Every node(user) can upload data to blockchain. When user upload data, user can gain rewards of Token. And every node(user) can query data. Querying data will cost Token.

登入完成後就可以使用網頁的功能，首先須InitToken讓Token發送到各個節點，
整體使用情境是中藥材的區塊鏈溯源，從上游的農夫到下游的消費者，讓整體的資訊到可以上傳到區塊鏈上，
完成資訊的透明化，且是可信任的資訊。每個節點(使用者)可以上傳資料，上傳資料可以獲得Token獎勵，
也可以查詢資料，查詢資料需要消耗Token。

![GITHUB](https://github.com/a037580238/supplychain/blob/master/nodejs_version/functions.PNG "functions")
