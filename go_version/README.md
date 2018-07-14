# Supplychain for medicine material with golang chaincode

1. 依照heroes-service-1.0.5中的步驟完成fabric go sdk的部署、測試
此部分為參考chainHero github repo:
[chainHero](https://github.com/chainHero/heroes-service)
2. 到達make步驟後，不會產生網頁，而是產生了連結到我撰寫的chaincode的restful api，目前是連結到固定IP，使用時請更換為自己的IP
3. 切換到medicine目錄底下，輸入指令npm run dev，產生網頁連接到第2步產生的restful api

目前完成首先進入IP/tables頁面，並點選左邊Actions底下的InitLedger產生初始數據:

![GITHUB](https://github.com/a037580238/supplychain/blob/master/go_version/InitLedger.png "InitLedger")

在來回到IP/tables頁面，會看到產生的結果:

![GITHUB](https://github.com/a037580238/supplychain/blob/master/go_version/tables.png "tables")

接下來點選左邊的UploadDate可以進入UploadData畫面，但進入Chaincode進行上傳的動作目前尚在Debug:

![GITHUB](https://github.com/a037580238/supplychain/blob/master/go_version/UploadData.png "UploadData")
