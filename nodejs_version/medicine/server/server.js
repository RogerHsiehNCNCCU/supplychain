// server.js
const express        = require('express');
//const MongoClient    = require('mongodb').MongoClient;
const bodyParser     = require('body-parser');

const app            = express();

const port = 8000;

//為了使用cookie以及mysql的部分
const cookieParser = require('cookie-parser');
const mysql = require("mysql");

//DataBase
let con = mysql.createConnection({
    host: "10.232.239.107",
    user: "root",
    password: "root",
    database: "medicine",
    port: "3306"
});

//db state
con.connect(function (err){
   if(err){
       console.log('connecting error');
       return;
   } 
    console.log('connecting success');
});


app.set('view engine', 'ejs');
app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: true }));//should see the body as an object in the terminal
app.use( express.static( "public" ) );//為了引入image、js等檔案

app.use(function (req, res, next){
    req.con=con;
    next();
});
app.use(cookieParser());

require('../app/routes')(app, {});
app.listen(port, () => {
  console.log('We are live on ' + port);
});