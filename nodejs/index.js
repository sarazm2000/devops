var express = require('express')
var cors = require('cors')
var app = express()

app.use(cors())



app.listen(3000, function () {
  console.log('CORS-enabled web server listening on port 80')
})

var bodyParser = require('body-parser');
// app.use(bodyParser.json());
// app.use(bodyParser.urlencoded({ extended: false }));

app.use(bodyParser.json()); // support json encoded bodies
app.use(bodyParser.urlencoded({
    extended: true
  }));
// app.use('/', router);
var crypto = require('crypto');

const redis = require("redis");
const client = redis.createClient(6379,"redis");

app.post('/', function (req, res, next) {
console.log(req.body.input);
if (req.body.input.length < 8){
    return ;
}
var hashed = crypto.createHash('sha256').update(req.body.input).digest('base64');
client.set( hashed, req.body.input, redis.print);

// client.get("hashedInput", redis.print)
res.send({
    "message" : hashed,
});
})

app.get('/', function(req, res, next){
  const input = req.query.input;
  console.log(req.query.input)
  var hashedString;
  client.get(input , function(err, reply) {
    if ( reply == null && err != null){
      res.send({
        "message" : "nadarim"
      })
    } else if (reply == null) {
      res.send({
        "message" : "nadarim"
      })
    } else {
      hashedString = reply;
      console.log(hashedString);
      res.send({
        "message" : hashedString,
    });
    }
  });
 
})