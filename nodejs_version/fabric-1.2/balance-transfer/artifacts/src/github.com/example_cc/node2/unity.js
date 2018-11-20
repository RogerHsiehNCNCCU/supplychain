'use strict';
const shim = require('fabric-shim')
const util = require('util')

let Chaincode = class {
    
    async Init(stub){
        console.info('====== Instantiated unity chaincode =======');
        let ret = stub.getFunctionAndParameters();
        console.info(ret);
        let args = ret.params;
        // initialise only if 4 parameters passed.
        if (args.length != 4) {
          return shim.error('Incorrect number of arguments. Expecting 4');
        }
//\"args\":[\"a\",\"100\",\"b\",\"200\"
        let A = args[0];
        let B = args[2];
        let Aval = args[1];
        let Bval = args[3];

        if (typeof parseInt(Aval) !== 'number' || typeof parseInt(Bval) !== 'number') {
          return shim.error('Expecting integer value for asset holding');
        }

        try {
          await stub.putState(A, Buffer.from(Aval));
          try {
            await stub.putState(B, Buffer.from(Bval));
            return shim.success();
          } catch (err) {
            return shim.error(err);
          }
        } catch (err) {
          return shim.error(err);
        }
    }
    
    // The Invoke method is called as a result of an application request to run the Smart Contract
    // function to be called, with arguments
    async Invoke(stub){
        let ret = stub.getFunctionAndParameters();
        console.info(ret);
        
        let method = this[ret.fcn];
        if (!method){
            console.error('no function of name:' + ret.fcn + ' found');
            throw new Error('Received unknown function' + ret.fcn + ' invocation');
        }
        try {
            let payload = await method(stub, ret.params, this);
            return shim.success(payload);
        } catch (err){
            console.log(err);
            return shim.error(err);
        }
    }
    
    //測試用的function
    async queryTest(stub, args, thisClass){
      console.info('============ START : queryTest ============');
      if(args.length !=1){
          throw new Error('Incorrect number of arguments. Expecting 1');
      }
        let valueAsBytes = await stub.getState("a");//get the key's value from chaincode state
        if(!valueAsBytes || valueAsBytes.toString().length <= 0){
            throw new Error(args[0] + ' does not exist: ');
        }
        console.log(valueAsBytes.toString());
        return valueAsBytes;
    }
    
    //初始化各個腳色及數據
    async initLedger(stub, args, thisClass){
        console.info('====== START : initialize Ledger ======');
        var time = thisClass.Now();
        var timeFullString = time[0]+time[1]+time[2]+time[3]+time[4]+time[5];
        var ymd = time[0]+time[1]+time[2];
        var hms = time[3]+":"+time[4]+":"+time[5];
        let Materials = [];
        //MaterialID 在 putstate時再加入在Key    Key : value

        Materials.push({
            Name: 'gochi',
            Efficacy: 'good for body',
            Color: 'red',
            HarvestBatch: 'B1',
            Action: '採收、初加工',
            Place: '四川',
            Weather: '雨後',
            Number: '10',
            Unit: '公斤',
            Temperature: '27',
            Fertilizer: '過磷酸鈣',
            FirstBatch: 'F1',
            Skill: '曝曬',
            InspectBatch: '',
            Inspecter: '',
            Inspect: '',
            TotalGrey: '',
            SO2: '',
            PurchaseBatch: '',
            ContractNo: '',
            ProduceBatch: '',
            ProductNo: '',
            ProductName: '',
            OwnerID: 'POC0',
            TimeStampDate: ymd,
            TimeStampTime: hms,
        });
        Materials.push({
            Name: '山藥',
            Efficacy: '保健美容',
            Color: 'purple',
            HarvestBatch: 'B2',
            Action: '採收、初加工',
            Place: '南投',
            Weather: '雨前',
            Number: '50',
            Unit: '斤',
            Temperature: '29',
            Fertilizer: '硫酸鉀',
            FirstBatch: 'F2',
            Skill: '去皮',
            InspectBatch: '',
            Inspecter: '',
            Inspect: '',
            TotalGrey: '',
            SO2: '',
            PurchaseBatch: '',
            ContractNo: '',
            ProduceBatch: '',
            ProductNo: '',
            ProductName: '',
            OwnerID: 'POC1',
            TimeStampDate: ymd,
            TimeStampTime: hms,
        });
        Materials.push({
            Name: '蓮子',
            Efficacy: '美白',
            Color: 'white',
            HarvestBatch: 'B3',
            Action: '採收、初加工',
            Place: '釜山',
            Weather: '晴天',
            Number: '1',
            Unit: '噸',
            Temperature: '23',
            Fertilizer: '過磷酸鈣',
            FirstBatch: 'F3',
            Skill: '過濾雜質',
            InspectBatch: '',
            Inspecter: '',
            Inspect: '',
            TotalGrey: '',
            SO2: '',
            PurchaseBatch: '',
            ContractNo: '',
            ProduceBatch: '',
            ProductNo: '',
            ProductName: '',
            OwnerID: 'POC3',
            TimeStampDate: ymd,
            TimeStampTime: hms,
        });
        Materials.push({
            Name: '',
            Efficacy: '',
            Color: '',
            HarvestBatch: 'B1',
            Action: '抽樣檢驗',
            Place: '武漢',
            Weather: '',
            Number: '',
            Unit: '',
            Temperature: '',
            Fertilizer: '',
            FirstBatch: '',
            Skill: '',
            InspectBatch: 'I1',
            Inspecter: '輔仁堂製藥公司',
            Inspect: '合格',
            TotalGrey: '1.50%',
            SO2: '5mg/kg',
            PurchaseBatch: '',
            ContractNo: '',
            ProduceBatch: '',
            ProductNo: '',
            ProductName: '',
            OwnerID: 'POC4',
            TimeStampDate: ymd,
            TimeStampTime: hms,
        });
        Materials.push({
            Name: '',
            Efficacy: '',
            Color: '',
            HarvestBatch: 'B2',
            Action: '抽樣檢驗',
            Place: '台灣',
            Weather: '',
            Number: '',
            Unit: '',
            Temperature: '',
            Fertilizer: '',
            FirstBatch: '',
            Skill: '',
            InspectBatch: 'I2',
            Inspecter: '禾伸堂製藥公司',
            Inspect: '不合格',
            TotalGrey: '2.50%',
            SO2: '10mg/kg',
            PurchaseBatch: '',
            ContractNo: '',
            ProduceBatch: '',
            ProductNo: '',
            ProductName: '',
            OwnerID: 'POC5',
            TimeStampDate: ymd,
            TimeStampTime: hms,
        });
        Materials.push({
            Name: '',
            Efficacy: '',
            Color: '',
            HarvestBatch: 'B1',
            Action: '採購管理',
            Place: '玄武輔仁堂中藥材原材料庫1',
            Weather: '',
            Number: '10',
            Unit: '公斤',
            Temperature: '',
            Fertilizer: '',
            FirstBatch: '',
            Skill: '',
            InspectBatch: '',
            Inspecter: '',
            Inspect: '',
            TotalGrey: '',
            SO2: '',
            PurchaseBatch: 'Pur1',
            ContractNo: 'C1',
            ProduceBatch: '',
            ProductNo: '',
            ProductName: '',
            OwnerID: 'POC5',
            TimeStampDate: ymd,
            TimeStampTime: hms,
        });
        Materials.push({
            Name: '',
            Efficacy: '',
            Color: '',
            HarvestBatch: 'B1',
            Action: '加工生產',
            Place: '玄武輔仁堂中藥材炮製車間',
            Weather: '',
            Number: '10',
            Unit: '公斤',
            Temperature: '',
            Fertilizer: '',
            FirstBatch: '',
            Skill: '低溫烘乾-等級分選-靜電除染',
            InspectBatch: '',
            Inspecter: '',
            Inspect: '',
            TotalGrey: '',
            SO2: '',
            PurchaseBatch: 'Pur1',
            ContractNo: '',
            ProduceBatch: 'Pro1',
            ProductNo: '',
            ProductName: '',
            OwnerID: 'POC6',
            TimeStampDate: ymd,
            TimeStampTime: hms,
        });
        Materials.push({
            Name: '',
            Efficacy: '',
            Color: '',
            HarvestBatch: 'B1',
            Action: '銷售管理',
            Place: '上海錦秀堂中藥店',
            Weather: '',
            Number: '3',
            Unit: '件',
            Temperature: '',
            Fertilizer: '',
            FirstBatch: '',
            Skill: '',
            InspectBatch: '',
            Inspecter: '',
            Inspect: '',
            TotalGrey: '',
            SO2: '',
            PurchaseBatch: '',
            ContractNo: '',
            ProduceBatch: 'Pro1',
            ProductNo: '6925802110699',
            ProductName: '寧安堡黑枸杞',
            OwnerID: 'POC7',
            TimeStampDate: ymd,
            TimeStampTime: hms,
        });

        let TMC = [{
            Name : 'TMC',
            Token : 29850000,
            TimeStampDate: ymd,
            TimeStampTime: hms,
        }];  
    
        
        let PoC = [{
            Name : 'farmer1',
            Token : 10000,
            Role : 'farmer',
            TimeStampDate: ymd,
            TimeStampTime: hms,
        },{
            Name : 'farmer2',
            Token : 10000,
            Role : 'farmer',
            TimeStampDate: ymd,
            TimeStampTime: hms, 
        },{
            Name : 'farmer3',
            Token : 10000,
            Role : 'farmer',
            TimeStampDate: ymd,
            TimeStampTime: hms, 
        },{
            Name : 'farmer4',
            Token : 10000,
            Role : 'farmer',
            TimeStampDate: ymd,
            TimeStampTime: hms,
        },{
            Name : 'farmer5',
            Token : 10000,
            Role : 'farmer',
            TimeStampDate: ymd,
            TimeStampTime: hms, 
        },{
            Name : 'plantBase1',
            Token : 10000,
            Role : 'plantBase',
            TimeStampDate: ymd,
            TimeStampTime: hms, 
        },{
            Name : 'plantBase2',
            Token : 10000,
            Role : 'plantBase',
            TimeStampDate: ymd,
            TimeStampTime: hms,
        },{
            Name : 'wholesaleMarket1',
            Token : 10000,
            Role : 'wholesaleMarket',
            TimeStampDate: ymd,
            TimeStampTime: hms, 
        },{
            Name : 'wholesaleMarket2',
            Token : 10000,
            Role : 'wholesaleMarket',
            TimeStampDate: ymd,
            TimeStampTime: hms, 
        },{
            Name : 'materialCompony1',
            Token : 10000,
            Role : 'materialCompony',
            TimeStampDate: ymd,
            TimeStampTime: hms, 
        },{
            Name : 'wholesaleCompony1',
            Token : 10000,
            Role : 'wholesaleCompony',
            TimeStampDate: ymd,
            TimeStampTime: hms,
        },{
            Name : 'drugstore1',
            Token : 10000,
            Role : 'drugstore',
            TimeStampDate: ymd,
            TimeStampTime: hms,
        },{
            Name : 'drugstore2',
            Token : 10000,
            Role : 'drugstore',
            TimeStampDate: ymd,
            TimeStampTime: hms,
        },{
            Name : 'hospital1',
            Token : 10000,
            Role : 'hospital',
            TimeStampDate: ymd,
            TimeStampTime: hms,
        },{
            Name : 'hospital2',
            Token : 10000,
            Role : 'hospital',
            TimeStampDate: ymd,
            TimeStampTime: hms,
        }];
        
        
        let TransactionRecord = [{
            Action : 'upload',
            Fee : 5,
            Participant : 'POC1',
            TimeStampDate: ymd,
            TimeStampTime: hms,
        },{
            Action : 'upload',
            Fee : 5,
            Participant : 'POC0',
            TimeStampDate: ymd,
            TimeStampTime: hms,
        },{
            Action : 'get',
            Fee : 5,
            Participant : 'POC2',
            TimeStampDate: ymd,
            TimeStampTime: hms,
        },{
            Action : 'get',
            Fee : 5,
            Participant : 'POC0',
            TimeStampDate: ymd,
            TimeStampTime: hms,
        },{
            Action : 'upload',
            Fee : 5,
            Participant : 'POC0',
            TimeStampDate: ymd,
            TimeStampTime: hms,
        }];
        
        
        for (let i = 0; i < Materials.length; i++) {
          Materials[i].docType = 'material';
          await stub.putState('MATERIAL' + i, Buffer.from(JSON.stringify(Materials[i])));
          console.info('Added <--> ', Materials[i]);
            
          //建立composite key，可以進行不同值的query
          let indexName = 'HarvestBatch~MATERIAL';
          let HarvestBatchIndexKey = await stub.createCompositeKey(indexName, [Materials[i].HarvestBatch, 'MATERIAL'+i]);
          console.info(HarvestBatchIndexKey);
          //  Save index entry to state. Only the key name is needed, no need to store a duplicate copy of the marble.
          //  Note - passing a 'nil' value will effectively delete the key from state, therefore we pass null character as value
          //但只存nil好像行不通
          //await stub.putState(HarvestBatchIndexKey, Buffer.from('\u0000'));
          await stub.putState(HarvestBatchIndexKey, Buffer.from(JSON.stringify(Materials[i])));
            
          //建立Product Batch composite key，可以進行不同值的query
          let indexName2 = 'ProductNo~MATERIAL';
          if(Materials[i].ProductNo != ''){
              let ProductNoIndexKey = await stub.createCompositeKey(indexName2, [Materials[i].ProductNo, 'MATERIAL'+i]);
              console.info(ProductNoIndexKey);
              await stub.putState(ProductNoIndexKey, Buffer.from(JSON.stringify(Materials[i])));
          }
        }
        
        
        
        for (let i=0; i< TMC.length; i++){
          TMC[i].docType = 'tmc';
          await stub.putState('TMC' + i, Buffer.from(JSON.stringify(TMC[i])));
          console.info('Added <-->, ', TMC[i]);
        }  
        
        for (let i=0; i< PoC.length; i++){
          PoC[i].docType = 'poc';
          await stub.putState('POC' + i, Buffer.from(JSON.stringify(PoC[i])));
          console.info('Added <-->, ', PoC[i]);
        }  
        //測試函數用
        for (let i=0; i< TransactionRecord.length; i++){
          TransactionRecord[i].docType = 'transactionRecord';
          await stub.putState('TRANSACTIONRECORD' + i, Buffer.from(JSON.stringify(TransactionRecord[i])));
          console.info('Added <-->, ', TransactionRecord[i]);
        }  
    
        console.info('============= END : initialize Ledger ===========');
    }

    //上傳數據(食品藥材 Material)
    async uploadData(stub, args, thisClass){
        console.info('====== START : upload Data ======');
        var time = thisClass.Now();
        var timeFullString = time[0]+time[1]+time[2]+time[3]+time[4]+time[5];
        var ymd = time[0]+time[1]+time[2];
        var hms = time[3]+":"+time[4]+":"+time[5];
        
        if(args.length !=24){
            throw new Error('Incorrect number of arguments. Expecting 24');
        }

        var material = {
            docType: 'material',
            Name: args[0],
            Efficacy: args[1],
            Color: args[2],
            HarvestBatch: args[3],
            Action: args[4],
            Place: args[5],
            Weather: args[6],
            Number: args[7],
            Unit: args[8],
            Temperature: args[9],
            Fertilizer: args[10],
            FirstBatch: args[11],
            Skill: args[12],
            InspectBatch: args[13],
            Inspecter: args[14],
            Inspect: args[15],
            TotalGrey: args[16],
            SO2: args[17],
            PurchaseBatch: args[18],
            ContractNo: args[19],
            ProduceBatch: args[20],
            ProductNo: args[21],
            ProductName: args[22],
            OwnerID: args[23],
            TimeStampDate: ymd,
            TimeStampTime: hms
        };

        //因為用queryAllByKey一直讀不到他的length及值
        //所以改用getAllResults自己來getStateByRange
        let method = thisClass['getAllResults'];
        //let queryResults = await method(stub,'MATERIAL',thisClass);
        let querystartKey = 'MATERIAL'+'0';
        let queryendKey = 'MATERIAL'+'99999';
        
        let queryIterator = await stub.getStateByRange(querystartKey, queryendKey);
        let queryResults = await method(queryIterator, false);
        //未必對
        //let queryResultsParse = JSON.parse(queryResults);
        await stub.putState("MATERIAL"+queryResults.length, Buffer.from(JSON.stringify(material)));
        
        //建立composite key，可以進行不同值的query
        let indexName = 'HarvestBatch~MATERIAL';
        let HarvestBatchIndexKey = await stub.createCompositeKey(indexName, [material.HarvestBatch, 'MATERIAL'+queryResults.length]);
        console.info(HarvestBatchIndexKey);
        //  Save index entry to state. Only the key name is needed, no need to store a duplicate copy of the marble.
        //  Note - passing a 'nil' value will effectively delete the key from state, therefore we pass null character as value
        //await stub.putState(HarvestBatchIndexKey, Buffer.from('\u0000'));
        await stub.putState(HarvestBatchIndexKey, Buffer.from(JSON.stringify(material)));
        
        //建立Product Batch composite key，可以進行不同值的query
        indexName = 'ProductNo~MATERIAL';
        let ProductNoIndexKey = await stub.createCompositeKey(indexName, [material.ProductNo, 'MATERIAL'+queryResults.length]);
        console.info(ProductNoIndexKey);
        await stub.putState(ProductNoIndexKey, Buffer.from(JSON.stringify(material)));
        
        //檢查是否超出一日100額度，並儲值10 token
        //因為couchdb才能使用rich query，這邊先使用Key的query
        let queryPocResults = await stub.getState(args[23]);
        let queryPocResultsParse = JSON.parse(queryPocResults);
        
        //運用async chaincode要有await不然會有unexpected token 0 錯誤
        //let queryTransactionRecordResults = await method(stub,'TRANSACTIONRECORD',thisClass);
        //let queryTransactionRecordResultsParse = JSON.parse(queryTransactionRecordResults);
        
        let startKey = 'TRANSACTIONRECORD'+'0';
        let endKey = 'TRANSACTIONRECORD'+'99999';
        
        let iterator = await stub.getStateByRange(startKey, endKey);
        //method = thisClass['getAllResults'];
        let results = await method(iterator, false);
        
        let todayRecord = [];
        for(var i=0;i<results.length;i++){
            if(results[i].Record.Action=="upload" && results[i].Record.TimeStampDate==time[0]+time[1]+time[2] && results[i].Record.Participant==args[23]){
                todayRecord.push(results[i]);
            }
        }
        if(todayRecord.length<10){//如果當日小於10比上傳，加10 token
            queryPocResultsParse.Token += 10;
            await stub.putState(args[23], Buffer.from(JSON.stringify(queryPocResultsParse)));
            
            //數據銀行Token減10
            let queryTMCResults = await stub.getState("TMC0");
            let queryTMCResultsParse = JSON.parse(queryTMCResults);
            queryTMCResultsParse.Token -= 10;
            await stub.putState("TMC0",Buffer.from(JSON.stringify(queryTMCResultsParse)));
        }
        //JSON.stringify(response_payloads[0].toString('utf8'))
        //紀錄transactionRecord
        let transactionRecord = {
            Action: "upload", Fee:0, Participant: args[23], TimeStampDate: time[0]+time[1]+time[2], TimeStampTime:time[3]+time[4]+time[5]
        };
        await stub.putState("TRANSACTIONRECORD"+results.length,Buffer.from(JSON.stringify(transactionRecord)));
        
        console.info('============= END : upload Data ===========');
    }
    
    async queryAllByKey(stub, args, thisClass){
        console.info('===== START : query All By Key =====');
        let startKey = args[0]+'0';
        let endKey = args[0]+'99999';
        
        let iterator = await stub.getStateByRange(startKey, endKey);
        let method = thisClass['getAllResults'];
        let results = await method(iterator, false);
        
        return Buffer.from(JSON.stringify(results));
        
        console.info('===== END : query All By Key =====');
    }
    async queryAllByKeyWithoutBuffer(stub, args, thisClass){
        console.info('===== START : query All By Key =====');
        let startKey = args[0]+'0';
        let endKey = args[0]+'99999';
        
        let iterator = await stub.getStateByRange(startKey, endKey);
        let method = thisClass['getAllResults'];
        let results = await method(iterator, false);
        
        return results;
        
        console.info('===== END : query All By Key =====');
    }
    
    //數據查詢
    async queryData(stub, args, thisClass){
        console.info('====== START : query Data ======');
        if (args.length != 2) {//Material Key , Poc Key(查詢者的Key)
            throw new Error('Incorrect number of arguments. Expecting 2');
        }
        
        //需要先檢查是否滿足最低積分要求，並有扣除5積分及2%手續費流程
        //如果檢查不足，引導至積分充值畫面
        let queryPocResults = await stub.getState(args[1]);
        let queryPocResultsParse = JSON.parse(queryPocResults);
        if(queryPocResultsParse.Token>=5){
            queryPocResultsParse.Token -=5;
            await stub.putState(args[1],Buffer.from(JSON.stringify(queryPocResultsParse)));
            
            let queryMaterialResults = await stub.getState(args[0]);
            let queryMaterialResultsParse = JSON.parse(queryMaterialResults);
            
//            let method = thisClass['queryDataInsert'];
//            let queryDataInsert = await method(stub,args,thisClass);
            
            //查詢資料擁有者(PoC2) + 5 * 0.98 Token
            let queryPoc2Results = await stub.getState(queryMaterialResultsParse.OwnerID);
            let queryPoc2ResultsParse = JSON.parse(queryPoc2Results);
            queryPoc2ResultsParse.Token = queryPoc2ResultsParse.Token + 5 * 0.98;
            await stub.putState(queryMaterialResultsParse.OwnerID,Buffer.from(JSON.stringify(queryPoc2ResultsParse)));
            
            //數據銀行收益2%
            let queryTMCResults = await stub.getState("TMC0");
            let queryTMCResultsParse = JSON.parse(queryTMCResults);
            queryTMCResultsParse.Token = queryTMCResultsParse.Token+ 5*0.02;
            await stub.putState("TMC0", Buffer.from(JSON.stringify(queryTMCResultsParse)));
            
            //紀錄transactionRecord
            var time = thisClass.Now();
            //let method = thisClass['queryAllByKeyWithoutBuffer'];
            
            let startKey = 'TRANSACTIONRECORD'+'0';
            let endKey = 'TRANSACTIONRECORD'+'99999';

            let iterator = await stub.getStateByRange(startKey, endKey);
            let method = thisClass['getAllResults'];
            let results = await method(iterator, false);
            
            //let queryTransactionRecordResults = await method(stub,'TRANSACTIONRECORD', thisClass);
            //let queryTransactionRecordResultsParse = JSON.parse(queryTransactionRecordResults);
            let transactionRecord = {
                Action: "get", Fee:5, Participant: args[1], TimeStampDate: time[0]+time[1]+time[2], TimeStampTime:time[3]+time[4]+time[5]
            };
            await stub.putState("TRANSACTIONRECORD"+results.length,Buffer.from(JSON.stringify(transactionRecord)));
            
            return queryMaterialResults;
            
        }else{ //積分不夠支付 5 token
            //引導到積分充值畫面
            console.log("引導到積分充值畫面");
        }
        
        console.info('============= END : Upload Data ===========');
    }
    
    //數據查詢 with partial composite key
    async queryHarvestBatch(stub, args, thisClass){
        console.info('====== START : query Data HarvestBatch ======');
        if (args.length != 2) {//HarvestBatch , Poc Key(查詢者的Key)
            throw new Error('Incorrect number of arguments. Expecting 2');
        }
        
        //需要先檢查是否滿足最低積分要求，並有扣除5積分及2%手續費流程
        //如果檢查不足，引導至積分充值畫面
        let queryPocResults = await stub.getState(args[1]);
        let queryPocResultsParse = JSON.parse(queryPocResults);
        if(queryPocResultsParse.Token>=5){
            
            //let queryMaterialResults = await stub.getState(args[0]);
            //let queryMaterialResultsParse = JSON.parse(queryMaterialResults);
            
            // Query the HarvestBatch Key index by color
            // This will execute a key range query on all keys starting with 'color'
            let HarvestBatchResultsIterator = await stub.getStateByPartialCompositeKey('HarvestBatch~MATERIAL', [args[0]]);
            
            let method = thisClass['getAllResults'];
            let HarvestBatchResults = await method(HarvestBatchResultsIterator, false);
            
            //query出複數筆，需要扣token複數次
            for(let i=0; i< HarvestBatchResults.length; i++){
                queryPocResultsParse.Token -=5;
                await stub.putState(args[1],Buffer.from(JSON.stringify(queryPocResultsParse)));
            }

            //查詢資料擁有者(PoC2) + 5 * 0.98 Token
            //因為要增加的對象，是複數個，所以要這樣特別處理
            let tempOwnerID = new Array();
            for(let i=0; i< HarvestBatchResults.length; i++){
                //let queryPoc2Results = await stub.getState(HarvestBatchResults[i].Record.OwnerID);
                //let queryPoc2ResultsParse = JSON.parse(queryPoc2Results); 
                tempOwnerID.push(HarvestBatchResults[i].Record.OwnerID);
//                queryPoc2ResultsParse.Token = queryPoc2ResultsParse.Token + 5 * 0.98;
//                await stub.putState(HarvestBatchResults[i].Record.OwnerID,Buffer.from(JSON.stringify(queryPoc2ResultsParse)));
            }
            let tempPoc2 = new Array();
            for(let i=0; i< tempOwnerID.length; i++){
                let whotochange=i;
                let count=0;
                let countTokenChange=0;
                //要增加的對象重複時，為了讓要增加的對象可以增加Token複數次
                //把getState的值使用在第一次出現的那次
                for(let j=0; j< tempOwnerID.length; j++){
                    if(tempOwnerID[i] == tempOwnerID[j] && count==0){
                        whotochange=j;
                        count++;
                    }
                }
                let queryPoc2Results = await stub.getState(tempOwnerID[i]);
                let queryPoc2ResultsParse = JSON.parse(queryPoc2Results);
                //這邊進行query時，上面減少token那邊的狀態，似乎還來不及寫進區塊鏈，成為world state，所以沿用上面的Token的值
                if(args[1] == tempOwnerID[i] && countTokenChange==0){
                    countTokenChange++;
                    queryPoc2ResultsParse.Token = queryPocResultsParse.Token;
                }
                tempPoc2.push(queryPoc2ResultsParse);
//                if(i == whotochange){
//                    tempPoc2[i].Token = tempPoc2[i].Token + 5 * 0.98;
//                    await stub.putState(tempOwnerID[i],Buffer.from(JSON.stringify(tempPoc2[i])));
//                }else{
                    tempPoc2[whotochange].Token = tempPoc2[whotochange].Token + 5 * 0.98;
                    await stub.putState(tempOwnerID[whotochange],Buffer.from(JSON.stringify(tempPoc2[whotochange])));
//                }
            }
            
            //應該是因為區塊鏈的狀態寫入速度關係，一次只能更改一次
            //在還沒寫入前，query出來還是未更改前的值，所以看起來只寫入一次
            //for迴圈沒辦法增加多次token，後來把query拿出來，解決了!
            
            let queryTMCResults = await stub.getState("TMC0");
            let queryTMCResultsParse = JSON.parse(queryTMCResults);
            //數據銀行收益2%
            for(let i=0; i< HarvestBatchResults.length; i++){
                queryTMCResultsParse.Token = queryTMCResultsParse.Token+ 5*0.02;
                await stub.putState("TMC0", Buffer.from(JSON.stringify(queryTMCResultsParse)));
            }
            
            //紀錄transactionRecord
            var time = thisClass.Now();
            //let method = thisClass['queryAllByKeyWithoutBuffer'];
            
            let startKey = 'TRANSACTIONRECORD'+'0';
            let endKey = 'TRANSACTIONRECORD'+'99999';

            let iterator = await stub.getStateByRange(startKey, endKey);
            //let method = thisClass['getAllResults'];
            let results = await method(iterator, false);
            
            //let queryTransactionRecordResults = await method(stub,'TRANSACTIONRECORD', thisClass);
            //let queryTransactionRecordResultsParse = JSON.parse(queryTransactionRecordResults);
            let transactionRecord = {
                Action: "getHarvestBatch", Fee:5, Participant: args[1], TimeStampDate: time[0]+time[1]+time[2], TimeStampTime:time[3]+time[4]+time[5]
            };
            await stub.putState("TRANSACTIONRECORD"+results.length,Buffer.from(JSON.stringify(transactionRecord)));
            
            return Buffer.from(JSON.stringify(HarvestBatchResults));
            
        }else{ //積分不夠支付 5 token
            //引導到積分充值畫面
            console.log("引導到積分充值畫面");
        }
        
        console.info('============= END : query Data HarvestBatch ===========');
    }
    
    //QRcode查詢with產品No(partial composite key)
    async queryProductNo(stub, args, thisClass){
        console.info('====== START : QRcode query Data Product No ======');
        
        if (args.length != 2) {//Product No , Poc Key(查詢者的Key)
            throw new Error('Incorrect number of arguments. Expecting 2');
        }
        
        let ProductNoResultsIterator = await stub.getStateByPartialCompositeKey('ProductNo~MATERIAL', [args[0]]);

        let method = thisClass['getAllResults'];
        let ProductNoResults = await method(ProductNoResultsIterator, false);

        let HB="";
        //應該只可能有一個，ProductNo是唯一的
        console.debug(ProductNoResults[0].Record.HarvestBatch);
        for(let i=0; i< ProductNoResults.length; i++){
            HB = ProductNoResults[i].Record.HarvestBatch;
        }
        
        //let arguments = new Array(HB,args[1]);
        //let arguments = new Array();
        //arguments.push(HB); arguments.push(args[1]);
        let sentargus = [HB, args[1]];
        //console.debug(sentargus); //會有錯 應該是不能debug array跟object?
        console.info(sentargus); 
        let method2 = thisClass['queryHarvestBatch'];
        let queryResults = await method2(stub, sentargus, thisClass);
        return queryResults;
        
        console.info('====== END : QRcode query Data Product No ======');
    }
    
    async queryProductNoOnly(stub, args, thisClass){
        console.info('====== START : QRcode query Data Product No Only ======');
        
        if (args.length != 2) {//Product No , Poc Key(查詢者的Key)
            throw new Error('Incorrect number of arguments. Expecting 2');
        }
        
        //需要先檢查是否滿足最低積分要求，並有扣除5積分及2%手續費流程
        //如果檢查不足，引導至積分充值畫面
        let queryPocResults = await stub.getState(args[1]);
        let queryPocResultsParse = JSON.parse(queryPocResults);
        if(queryPocResultsParse.Token>=5){
            queryPocResultsParse.Token -=5;
            await stub.putState(args[1],Buffer.from(JSON.stringify(queryPocResultsParse)));
            
            let ProductNoResultsIterator = await stub.getStateByPartialCompositeKey('ProductNo~MATERIAL', [args[0]]);

            let method = thisClass['getAllResults'];
            let ProductNoResults = await method(ProductNoResultsIterator, false);
            
            //查詢資料擁有者(PoC2) + 5 * 0.98 Token
            // 用 0 是因為一個產品只會有一筆
            let queryPoc2Results = await stub.getState(ProductNoResults[0].Record.OwnerID);
            let queryPoc2ResultsParse = JSON.parse(queryPoc2Results);
            queryPoc2ResultsParse.Token = queryPoc2ResultsParse.Token + 5 * 0.98;
            await stub.putState(ProductNoResults[0].Record.OwnerID,Buffer.from(JSON.stringify(queryPoc2ResultsParse)));
            
            //數據銀行收益2%
            let queryTMCResults = await stub.getState("TMC0");
            let queryTMCResultsParse = JSON.parse(queryTMCResults);
            queryTMCResultsParse.Token = queryTMCResultsParse.Token+ 5*0.02;
            await stub.putState("TMC0", Buffer.from(JSON.stringify(queryTMCResultsParse)));
            
            //紀錄transactionRecord
            var time = thisClass.Now();
            //let method = thisClass['queryAllByKeyWithoutBuffer'];
            
            let startKey = 'TRANSACTIONRECORD'+'0';
            let endKey = 'TRANSACTIONRECORD'+'99999';

            let iterator = await stub.getStateByRange(startKey, endKey);
            method = thisClass['getAllResults'];
            let results = await method(iterator, false);
            
            let transactionRecord = {
                Action: "get", Fee:5, Participant: args[1], TimeStampDate: time[0]+time[1]+time[2], TimeStampTime:time[3]+time[4]+time[5]
            };
            await stub.putState("TRANSACTIONRECORD"+results.length,Buffer.from(JSON.stringify(transactionRecord)));
            
            return Buffer.from(JSON.stringify(ProductNoResults));
            
        }else{ //積分不夠支付 5 token
            //引導到積分充值畫面
            console.log("引導到積分充值畫面");
        }
        
        console.info('====== END : QRcode query Data Product No Only ======');
    }
    
    async tokenStore(stub, args, thisClass){
        console.info('====== START : token Store ======');
        if (args.length != 2) {//POC Key以及儲值多少
            throw new Error('Incorrect number of arguments. Expecting 2');
        }
        
        //支付後才進入儲值階段，支付是線下處理
        
        //數據銀行的Token要扣掉(像銀行發錢一樣)
        let queryTMCResults = await stub.getState("TMC0");
        let queryTMCResultsParse = JSON.parse(queryTMCResults);
        queryTMCResultsParse.Token = queryTMCResultsParse.Token - parseInt(args[1]);
        await stub.putState("TMC0", Buffer.from(JSON.stringify(queryTMCResultsParse)));
        
        //POC儲值Token
        let queryPocResults = await stub.getState(args[0]);
        let queryPocResultsParse = JSON.parse(queryPocResults);
        queryPocResultsParse.Token += parseInt(args[1]);
        
        await stub.putState(args[0], Buffer.from(JSON.stringify(queryPocResultsParse)));
        
        //紀錄transactionRecord
        var time = thisClass.Now();
        //let method = thisClass['queryAllByKeyWithoutBuffer'];

        let startKey = 'TRANSACTIONRECORD'+'0';
        let endKey = 'TRANSACTIONRECORD'+'99999';

        let iterator = await stub.getStateByRange(startKey, endKey);
        let method = thisClass['getAllResults'];
        let results = await method(iterator, false);

        let transactionRecord = {
            Action: "store", Fee:0, Participant: args[0], TimeStampDate: time[0]+time[1]+time[2], TimeStampTime:time[3]+time[4]+time[5]
        };
        await stub.putState("TRANSACTIONRECORD"+results.length,Buffer.from(JSON.stringify(transactionRecord)));
        
        console.info('============= END : token Store ===========');
    }
    //可用queryByKey代替
//    async queryTokenByPoc(stub, args, thisClass){
//        console.info('====== START : query Token By Poc ======');
//        
//        console.info('============= END : query Token By Poc ===========');
//    }
    //可用queryHistoryByKey代替
//    async queryHistoryPoc(stub, args, thisClass){
//        console.info('====== START : query History Poc ======');
//        
//        console.info('============= END : query History Poc ===========');
//    }
    
    async queryHistoryByKey(stub, args, thisClass){
      console.info('============ START : getHistoryCar ============');
      if(args.length !=1){
          throw new Error('Incorrect number of arguments. Expecting 1');
      }

      let keyNumber = args[0];

      let resultsIterator = await stub.getHistoryForKey(keyNumber); //get the car from chaincode state
      let method = thisClass['getAllResults'];
      let results = await method(resultsIterator, true);

      return Buffer.from(JSON.stringify(results));
    }
    //可用queryAllByKey代替
//    async queryAllPocs(stub, args, thisClass){
//        console.info('====== START : query All Pocs ======');
//        
//        console.info('============= END : query Alll Pocs ===========');
//    }
    //js 中可呼叫原本的function不用多寫internal的
//    async internalQueryAllPocs(){
//        
//    }
    
//    async queryTransactionByKey(stub, args, thisClass){
//        console.info('====== START : query Transaction By Key ======');
//        
//        console.info('============= END : query Transaction By Key ===========');
//    }
    
//    async queryAllTransactions(stub, args, thisClass){
//        console.info('====== START : query All Transactions ======');
//        
//        console.info('============= END : query All Transactions ===========');
//    }
    
//    async internalQueryAllTransactions(){
//        
//    }
    
    async queryByKey(stub, args, thisClass){
        if(args.length !=1){
            throw new Error('Incorrect number of arguments. Expecting Key ex:POC1');
        }
        let keyNumber = args[0];
        
        let valueAsBytes = await stub.getState(keyNumber);//get the key's value from chaincode state
        if(!valueAsBytes || valueAsBytes.toString().length <= 0){
            throw new Error(keyNumber + ' does not exist: ');
        }
        console.log(valueAsBytes.toString());
        return valueAsBytes;
    }
    
    async createUser(stub, args, thisClass) {
        console.info('============= START : Create user ===========');
        if (args.length != 4) {
          throw new Error('Incorrect number of arguments. Expecting 4');
        }

        var user = {
          docType: 'User',
          name: args[1],
          phone: args[2],
          role: args[3]
        };

        await stub.putState(args[0], Buffer.from(JSON.stringify(user)));
        console.info('============= END : Create user ===========');
    }
    
//    async queryAllDatas(stub, args, thisClass) {
//        console.info('============ START : queryAllDatas ============');
//        let startKey = 'DATA0';
//        let endKey = 'DATA99999';
//
//        let iterator = await stub.getStateByRange(startKey, endKey);
//
//        let method = thisClass['getAllResults'];
//        let results = await method(iterator, false);
//
//        return Buffer.from(JSON.stringify(results));
//    } 
    
//    //rich query需要有CouchDB 現在環境沒有CouchDB
//    async queryCarBymake(stub, args, thisClass) {
//    
//        if (args.length < 1){
//            throw new Error('Incorrect number of arguments. Expecting make.');
//        }
//
//        //let make = args[0].toLowerCase();
//        let make = args[0];
//        let queryString = {};
//        queryString.selector = {};
//        queryString.selector.docType = 'car';
//        queryString.selector.make = make;
//        let method = thisClass['getQueryResultForQueryString'];
//        let queryResults = await method(stub, JSON.stringify(queryString), thisClass);
//        return queryResults;
//    }
    
    async getAllResults(iterator, isHistory) {
        let allResults = [];
        while (true) {
          let res = await iterator.next();

          if (res.value && res.value.value.toString()) {
            let jsonRes = {};
            console.log(res.value.value.toString('utf8'));

            if (isHistory && isHistory === true) {
              jsonRes.TxId = res.value.tx_id;
              jsonRes.Timestamp = res.value.timestamp;
              jsonRes.IsDelete = res.value.is_delete.toString();
              try {
                jsonRes.Value = JSON.parse(res.value.value.toString('utf8'));
              } catch (err) {
                console.log(err);
                jsonRes.Value = res.value.value.toString('utf8');
              }
            } else {
              jsonRes.Key = res.value.key;
              try {
                jsonRes.Record = JSON.parse(res.value.value.toString('utf8'));
              } catch (err) {
                console.log(err);
                jsonRes.Record = res.value.value.toString('utf8');
              }
            }
            allResults.push(jsonRes);
          }
          if (res.done) {
            console.log('end of data');
            await iterator.close();
            console.info(allResults);
            return allResults;
          }
        }
    }
    
    // =========================================================================================
    // getQueryResultForQueryString executes the passed in query string.
    // Result set is built and returned as a byte array containing the JSON results.
    // =========================================================================================
//    async getQueryResultForQueryString(stub, queryString, thisClass) {
//  
//      console.info('- getQueryResultForQueryString queryString:\n' + queryString)
//      let resultsIterator = await stub.getQueryResult(queryString);
//      let method = thisClass['getAllResults'];
//  
//      let results = await method(resultsIterator, false);
//      //let results = await getAllResults(resultsIterator, false);
//  
//      return Buffer.from(JSON.stringify(results));
//    }
    
    Now(){
        var y,mon,d,h,min,s;
        var Now=new Date();
        y= String(Now.getFullYear());
        mon= String(Now.getMonth()+1);
        d= String(Now.getDate());
        h= String(Now.getHours());
        min= String(Now.getMinutes());
        s= String(Now.getSeconds());
        //console.log(mon+" type: "+ typeof mon+" length: "+mon.length);
        if(mon.length==1) mon = "0"+mon;
        if(d.length==1) d = "0"+d;
        if(h.length==1) h = "0"+h;
        if(min.length==1) min="0"+min;
        if(s.length==1) s="0"+s;
            
        var all = [y,mon,d,h,min,s];
        
        return all
    }
    
};

shim.start(new Chaincode());