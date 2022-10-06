const https = require('https');
const fs = require('fs');
const path = require('path');

const hostname = '0.0.0.0';
const port = 3010;

const dirn = __dirname;

const options = {
    ca: fs.readFileSync(path.join(dirn, '../IntermediateCaCluster1/cert-chain.pem')),
    cert: fs.readFileSync(path.join(dirn, 'cert.pem')),
    key: fs.readFileSync(path.join(dirn, 'key.pem')),
    requestCert: true, // Client have to send own certificate, otherwise request will be dropped
    rejectUnauthorized: true, // If client send certficate not signed by certificates in chain, drop connection
};

const server = https.createServer(options, (req, res) => {
    res.statusCode = 200;
    res.setHeader('Content-Type', 'application/json');
    res.end(JSON.stringify({
        firtname: 'John',
        lastname: 'Smith',
        order: [
            {
                name: 'Sofa',
                items: 1,
                price: {
                    amount: 100,
                    currency: 'USD'
                }
            }
        ]
    }));
});

server.on('error',(e)=>console.log(e));

server.listen(port, hostname, () => {
    console.log(`Server running at http://${hostname}:${port}/`);
});
