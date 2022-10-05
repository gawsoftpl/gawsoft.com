const https = require('https');
const fs = require('fs');

const hostname = '0.0.0.0';
const port = 3100;

const options = { 
    ca: fs.readFileSync('ca/ca.cert.pem'), 
    cert: fs.readFileSync('invoice-server/server.cert'), 
    key: fs.readFileSync('invoice-server/private.key'), 
    rejectUnauthorized: true,
    requestCert: true,
}; 

const server = https.createServer(options, (req, res) => {
  res.statusCode = 200;
  res.setHeader('Content-Type', 'application/json');
  const userData = JSON.stringify({
     firstname: 'John',
     lastname: 'Ji',
     email: 'john.ji@example.com',
     invoice_number: "00095234",
     invoice_price: "10USD",
  });
  res.end(userData);
});

server.on('error',(e)=>console.log(e));

server.listen(port, hostname, () => {
  console.log(`Server running at http://${hostname}:${port}/`);
});

