var qmarket = require('./querycontroller.js');
var imarket = require('./invokecontroller.js');


module.exports = function(app){
    
    app.get('/queryAllDatas',function(req,res){
        console.log("I'm in routes queryAllDatas");
       qmarket.query_queryAllDatas(req,res); 
    });
    
    app.get('/queryData_Data1',function(req,res){
        console.log("I'm in routes queryData_Data1");
       qmarket.query_queryData_Data1(req,res); 
    });
    
    app.get('/queryData_Spe/:QSD',function(req,res){
        console.log("I'm in routes Specific Get");
       qmarket.query_queryData_Spe(req,res); 
    });
    
    app.get('/checkDataValid/:CDV',function(req,res){
        console.log("I'm in routes checkDataVaild!");
       qmarket.query_checkDataVaild(req,res); 
    });
    
    
    
    app.get('/initLedger',function(req,res){
        console.log("I'm in routes initLedger");
       imarket.initLedger(req,res); 
    });
    
    app.get('/invokeData_Data4/:TimeStamp',function(req,res){
        console.log("I'm in routes invokeData_Data4");
       imarket.invoke_invokeData_Data4(req,res); 
    });
    
    app.get('/invokeData_Spe/:BSD',function(req,res){
        console.log("I'm in routes Specific Post");
       imarket.invoke_invokeData_Spe(req,res); 
    });
}