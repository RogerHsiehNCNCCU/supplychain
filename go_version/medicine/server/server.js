// server.js
const express        = require('express');
//const MongoClient    = require('mongodb').MongoClient;
const bodyParser     = require('body-parser');

const app            = express();

const port = 8000;


app.set('view engine', 'ejs');
app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: true }));//should see the body as an object in the terminal
app.use( express.static( "public" ) );//為了引入image、js等檔案

require('../app/routes')(app, {});
app.listen(port, () => {
  console.log('We are live on ' + port);
});