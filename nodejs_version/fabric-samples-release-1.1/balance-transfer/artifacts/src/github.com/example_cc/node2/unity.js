'use strict';
const shim = require('fabric-shim')
const util = require('util')

let Chaincode = class {
    
    async Init(stub){
        console.info('====== Instantiated unity chaincode =======');
        return shim.success();
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
    
    //初始化各個腳色及數據
    async initLedger(stub, args, thisClass){
        console.info('====== START : initialize Ledger ======');
        var time = this.Now();
        var timeFullString = time[0]+time[1]+time[2]+time[3]+time[4]+time[5];
        var ymd = time[0]+time[1]+time[2];
        var hms = time[3]+time[4]+time[5];
        let Materials = [];
        //MaterialID 在 putstate時再加入在Key    Key : value
        Materials.push({
            Name: 'gochi',
            Efficacy: 'good for body',
            Color: 'red',
            OwnerID: 'POC0',
            TimeStampDate: ymd,
            TimeStampTime: hms,
        });
        Materials.push({
            Name: '山藥',
            Efficacy: '保健美容',
            Color: 'purple',
            OwnerID: 'POC1',
            TimeStampDate: ymd,
            TimeStampTime: hms,
        });
        Materials.push({
            Name: '蓮子',
            Efficacy: '美白',
            Color: 'white',
            OwnerID: 'POC2',
            TimeStampDate: ymd,
            TimeStampTime: hms,
        });
        
        let TMC = [{
            Name : 'TMC',
            Token : '29850000',
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
        }];
        
        
        for (let i = 0; i < Materials.length; i++) {
          Materials[i].docType = 'material';
          await stub.putState('MATERIAL' + i, Buffer.from(JSON.stringify(Materials[i])));
          console.info('Added <--> ', Materials[i]);
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
    
        console.info('============= END : initialize Ledger ===========');
    }

    //上傳數據(食品藥材 Material)
    async uploadData(stub, args, thisClass){
        console.info('====== START : upload Data ======');
        var time = this.Now();
        var timeFullString = time.[0]+time[1]+time[2]+time[3]+time[4]+time[5];
        
        if(args.length !=4){
            throw new Error('Incorrect number of arguments. Expecting 4');
        }

        var material = {
            docType: 'material',
            Name: args[0],
            Efficacy: args[1],
            Color: args[2],
            OwnerID: args[3]
        };
        let method = thisClass['queryAllByKey'];
        let queryResults = await method(stub,'MATERIAL',thisClass);
        //未必對
        let queryResultsParse = JSON.parse(queryResults);
        await stub.putState("MATERIAL"+queryResultsParse.length+1, Buffer.from(JSON.stringify(material)));
        
        //檢查是否超出一日100額度，並儲值10 token
        //因為couchdb才能使用rich query，這邊先使用Key的query
        let queryPocResults = await stub.getState(args[3]);
        let queryPocResultsParse = JSON.parse(queryPocResults);
        
        var time= this.Now();
        let queryTransactionRecordResults = method(stub,'TRANSACTIONRECORD',thisClass);
        let queryTransactionRecordResultsParse = JSON.parse(queryTransactionRecordResults);
        let todayRecord = [];
        for(var i=0;i<queryTransactionRecordResultsParse.length;i++){
            if(queryTransactionRecordResultsParse[i].Action=="upload" && queryTransactionRecordResultsParse[i].TimeStampDate==time[0]+time[1]+time[2] && queryTransactionRecordResultsParse[i].Participant==args[3]){
                todayRecord.append(queryTransactionRecordResultsParse[i]);
            }
        }
        if(todayRecord.length<10){//如果當日小於10比上傳，加10 token
            queryPocResultsParse.Token += 10;
            stub.putState(args[3], Buffer.from(JSON.stringify(queryPocResultsParse)));
            
            //數據銀行Token減10
            let queryTMCResults = await stub.getState("TMC0");
            let queryTMCResultsParse = JSON.parse(queryTMCResults);
            queryTMCResultsParse.Token -= 10;
            stub.putState("TMC0",Buffer.from(JSON.stringify(queryTMCResultsParse)));
        }
        
        console.info('============= END : upload Data ===========');
    }
    
    async queryAllByKey(stub, args, thisClass){
        console.info('===== START : query All By Key =====');
        let startKey = args[0]+'0';
        let endKey = args[0]+'99999';
        
        let iterator = await stub.getStartByRange(startKey, endKey);
        let results = await method(iterator, false);
        
        return Buffer.from(JSON.stringify(results));
        
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
            stub.putState(args[1],Buffer.from(JSON.stringify(queryPocResultsParse)));
            
            let queryMaterialResults = await stub.getState(args[0]);
            let queryMaterialResultsParse = JSON.parse(queryMaterialResults);
            
            //查詢資料擁有者(PoC2) + 5 * 0.98 Token
            let queryPoc2Results = await stub.getState(queryMaterialResultsParse.Owner);
            let queryPoc2ResultsParse = JSON.parse(queryPoc2Results);
            queryPoc2ResultsParse.Token = queryPoc2ResultsParse.Token + 5 * 0.98;
            stub.putState(queryMaterialResultsParse.Owner,Buffer.from(JSON.stringify(queryPoc2ResultsParse)));
            
            //數據銀行收益2%
            let queryTMCResults = await stub.getState("TMC0");
            let queryTMCResultsParse = JSON.parse(queryTMCResults);
            queryTMCResultsParse.Token = queryTMCResultsParse.Token+ 5*0.02;
            stub.putState("TMC0", Buffer.from(JSON.stringify(queryTMCResultsParse)));
            
            //紀錄transactionRecord
            var time = this.Now();
            let method = thisClass['queryAllByKey'];
            let queryTransactionRecordResults = await method(iterator, false);
            let queryTransactionRecordResultsParse = JSON.parse(queryTransactionRecordResults);
            let transactionRecord = {
                Action: "get", fee:5, participant: queryPocResultsParse.Name, TimeStampymd: time[0]+time[1]+time[2], TimeStamphms:time[3]+time[4]+time[5]
            };
            stub.putState("TRANSACTIONRECORD"+queryTransactionRecordResultsParse.length+1,Buffer.from(JSON.stringify(transactionRecord)));
            
            return Buffer.from(JSON.stringify(queryMaterialResults));
            
        }else{ //積分不夠支付 5 token
            //引導到積分充值畫面
        }
        
        console.info('============= END : Upload Data ===========');
    }
    
    async tokenStore(stub, args, thisClass){
        console.info('====== START : token Store ======');
        if (args.length != 2) {//POC Key以及儲值多少
            throw new Error('Incorrect number of arguments. Expecting 2');
        }
        
        //支付後才進入儲值階段，支付是線下處理
        
        let queryPocResults = await stub.getState(args[0]);
        let queryPocResultsParse = JSON.parse(queryPocResults);
        queryPocResultsParse.Token += args[1];
        
        stub.putState(args[0], Buffer.from(JSON.stringify(queryPocResultParse)));
        
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
    
    async queryAllDatas(stub, args, thisClass) {
        console.info('============ START : queryAllDatas ============');
        let startKey = 'DATA0';
        let endKey = 'DATA99999';

        let iterator = await stub.getStateByRange(startKey, endKey);

        let method = thisClass['getAllResults'];
        let results = await method(iterator, false);

        return Buffer.from(JSON.stringify(results));
    } 
    
    //rich query需要有CouchDB 現在環境沒有CouchDB
    async queryCarBymake(stub, args, thisClass) {
    
        if (args.length < 1){
            throw new Error('Incorrect number of arguments. Expecting make.');
        }

        //let make = args[0].toLowerCase();
        let make = args[0];
        let queryString = {};
        queryString.selector = {};
        queryString.selector.docType = 'car';
        queryString.selector.make = make;
        let method = thisClass['getQueryResultForQueryString'];
        let queryResults = await method(stub, JSON.stringify(queryString), thisClass);
        return queryResults;
    }
    
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
    async getQueryResultForQueryString(stub, queryString, thisClass) {
  
      console.info('- getQueryResultForQueryString queryString:\n' + queryString)
      let resultsIterator = await stub.getQueryResult(queryString);
      let method = thisClass['getAllResults'];
  
      let results = await method(resultsIterator, false);
      //let results = await getAllResults(resultsIterator, false);
  
      return Buffer.from(JSON.stringify(results));
    }
    
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