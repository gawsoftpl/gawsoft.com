var fs = require('fs');
var https = require('https');
var options = {
    hostname: 'localhost',
    port: 3100,
    path: '/',
    method: 'GET',
    key: fs.readFileSync('mail-server/private.key'),
    cert: fs.readFileSync('mail-server/server-wrong.cert'),
    ca: fs.readFileSync('ca/ca.cert.pem'),
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
