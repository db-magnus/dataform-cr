'use strict';
const express = require('express');
const app = express();

app.get('/', (req, res) => {
    const { spawn } = require('child_process');
    const ls = spawn('bash',['./script.sh']);
    ls.stdout.on('data', (data) => {
        console.log(`stdout: ${data}`);
    });

    ls.stderr.on('data', (data) => {
        console.log(`stderr: ${data}`);
    });

    ls.on('close', (code) => {
        console.log(`child process exited with code ${code}`);
    });
    ls.on('exit', function(code, signal) {
        console.log('child process finished');
            res.send(`request finished`);
    });
    
});

const port = process.env.PORT || 8080;
app.listen(port, () => {
    console.log(`webserver startet: listening on port ${port}`);
});