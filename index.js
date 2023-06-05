const crypto = require('crypto');
const readline = require('readline');

const read = readline.createInterface({
  input: process.stdin,
  output: process.stdout
});

read.question('Enter the data: ', (data) => {
  const hash = crypto.createHash('sha256').update(data).digest('hex');
  console.log('SHA256 hash:', hash);

  read.close();
});
