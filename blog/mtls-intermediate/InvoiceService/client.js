var fs = require('fs');
var https = require('https');
const path = require('path');

const dirn = __dirname;
var options = {
    hostname: 'localhost',
    port: 3010,
    path: '/',
    method: 'GET',
    key: fs.readFileSync(path.join(dirn,'key.pem')),
    cert: fs.readFileSync(path.join(dirn,'cert.pem')),
    ca: fs.readFileSync(path.join(dirn, '../IntermediateCaCluster1/cert-chain.pem'))
};

var req = https.request(options, function(res) {
    res.on('data', function(data) {
        process.stdout.write(data);
    });
});
req.end();
req.on('error', function(e) {
    console.error(e);
});
