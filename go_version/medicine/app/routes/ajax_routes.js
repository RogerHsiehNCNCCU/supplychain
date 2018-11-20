const fetch = require('node-fetch');

var http = require('http');
var querystring = require('querystring');

var $ = require('jQuery');

const url = "140.119.19.108"

module.exports = function(app, db) {
    var options = {
        host: url,
        port: 3000,
        path: '/api/Bike',
        method: 'GET'
    };
    
    /*app.get('/', function(req, res) {
        res.render('header');
    })*/
    app.get('/', function(req, res) {
       
        var data={
            "peers": ["peer0.org1.example.com","peer1.org1.example.com"],
            "fcn" : "initLedger",
            "args" : [""]
        };

        $.ajax({
            url: 'http://localhost:4000/channels/mychannel/chaincodes/unity',
            method: "POST",
            dataType: "json",
            contentType: "application/json; charset=utf-8",
            data: JSON.stringify(data),
            beforeSend: function(xhr) {
                 xhr.setRequestHeader("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MzE1ODk4NDEsInVzZXJuYW1lIjoiSmltIiwib3JnTmFtZSI6Ik9yZzEiLCJpYXQiOjE1MzE1NTM4NDF9.3nOt7dMXP-A6HPPd9xYQD7D_yn_Ou2wcZwox-mzw1L4")
            }, success: function(data){
                alert(data);
                //process the JSON data etc
                res.send(data);
            }
        })

        // curl -s -X POST \
        // http://localhost:4000/channels/mychannel/chaincodes/mycc \
        // -H "authorization: Bearer <put JSON Web Token here>" \
        // -H "content-type: application/json" \
        // -d '{
        //     "peers": ["peer0.org1.example.com","peer1.org1.example.com"],
        //     "fcn":"move",
        //     "args":["a","b","10"]
        // }'


        //res.send('hello world');
    })
    //查詢所有腳踏車
    app.get('/tables', (req, res) => {
        /*http.request(options, function(res) {
            console.log('STATUS: ' + res.statusCode);
            console.log('HEADERS: ' + JSON.stringify(res.headers));
            res.setEncoding('utf8');
            res.on('data', function (chunk) {
                console.log('BODY: ' + chunk);
            });
        }).end();*/
        fetch('http://140.119.19.108:3000/QueryAllPoCs')
            .then(res => res.text())
            .then((query_responses) => {
                console.log("query_responses is ", query_responses);
                console.log("Type is ", typeof query_responses);
                console.log("length : ", query_responses.length);
                let result = JSON.parse(query_responses.toString());
                console.log("Response is ", result);
                console.log("Type is ", typeof result);
                console.log("length : ", result.length);
                //第一次parse後還是string只是去掉了/ 所以再進行一次parse
                let result2 = JSON.parse(result.toString());
                console.log("Response is ", result2);
                console.log("Type is ", typeof result2);
                console.log("length : ", result2.length);

                result2.sort(function(a,b){
                    var nameA=a.Record.Name.toLowerCase(), nameB=b.Record.Name.toLowerCase()
                    if (nameA < nameB) //sort string ascending
                        return -1 
                    if (nameA > nameB)
                        return 1
                    return 0 //default return value (no sorting)
                })
                res.render("tables",{data: result2});
            });
            //.then(body => console.log(body));
    });

    app.get('/InitLedger', function(req, res) {
        fetch('http://140.119.19.108:3000/InitLedger')
            .then(res => res.json())
            .then(json => console.log(json));

        res.redirect('/tables');
    });
    //兩個名字相同變成一直redirect無限迴圈 get 跟 redirect
    //render是到那個ejs檔 不用加/ 是檔名    redirect會是 get 命令
    app.get('/UploadData', function(req, res) {

        res.render('uploadDataPage');
    });
    app.post('/UploadData', function(req, res) {
        var args={ Name:String(req.body.Name), Efficacy:String(req.body.Efficacy), Color:String(req.body.Color),
            Owner:String(req.body.Owner)};
        fetch('http://140.119.19.108:3000/UploadData', { 
            method: 'POST',
            body:    JSON.stringify(args),
            headers: { 'Content-Type': 'application/json' },
        })
           .then(res => res.json())
           .then(json => console.log(json))
           .catch(console.log(res)); //catch 有res error function可用嗎

        res.redirect('/tables');
    });
    //新增一台腳踏車
    app.post('/bike', (req, res) => {
        var args={ $class:String(req.body.$class), status:String(req.body.status), provider:String(req.body.provider),
             remark:String(req.body.remark), aid:String(req.body.aid)};
        fetch('http://140.119.19.108:3000/api/Bike', { 
            method: 'POST',
            body:    JSON.stringify(args),
            headers: { 'Content-Type': 'application/json' },
        })
            .then(res => res.json())
            .then(json => console.log(json));

        res.redirect('/bike');
    });
    //修改一台腳踏車
    app.get('/update_bike/:bike_Id', function(req, res) {
        console.log(req.params.bike_Id);
        fetch('http://140.119.19.108:3000/api/Bike/'+req.params.bike_Id)
            .then(res => res.text())
            .then((query_responses) => {
                let result = JSON.parse(query_responses.toString());
                console.log("Response is ", result);
                console.log("length : ", result.length);
                res.render("update_bike",{data: result});
        });
    });
    app.post('/update_bike/:bike_Id', (req, res) => {
        var args={ $class:String(req.body.$class), status:String(req.body.status), provider:String(req.body.provider),
            remark:String(req.body.remark), aid:String(req.body.aid)};
            fetch('http://140.119.19.108:3000/api/Bike/'+args.aid, { 
                method: 'PUT',
                body:    JSON.stringify(args),
                headers: { 'Content-Type': 'application/json' },
            })
                .then(res => res.json())
                .then(json => console.log(json));
    
            res.redirect('/bike');    
    });

    //刪除一台腳踏車
    app.get('/delete_bike/:bike_Id', function(req, res) {
        console.log(req.params.bike_Id);
        fetch('http://140.119.19.108:3000/api/Bike/'+req.params.bike_Id)
            .then(res => res.text())
            .then((query_responses) => {
                let result = JSON.parse(query_responses.toString());
                console.log("Response is ", result);
                console.log("length : ", result.length);
                res.render("delete_bike",{data: result});
        });
    });
    app.post('/delete_bike/:bike_Id', (req, res) => {
        var args={ $class:String(req.body.$class), status:String(req.body.status), provider:String(req.body.provider),
            remark:String(req.body.remark), aid:String(req.body.aid)};
            fetch('http://140.119.19.108:3000/api/Bike/'+args.aid, { 
                method: 'DELETE',
                //body:    JSON.stringify(args),
                headers: { 'Content-Type': 'application/json' },
            })
                .then(res => res.json())
                .then(json => console.log(json));
    
            res.redirect('/bike');    
    });

};