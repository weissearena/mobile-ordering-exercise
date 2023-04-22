import express from 'express';
import bodyParser from 'body-parser';

const app = express();
app.use(bodyParser.json());

const startTimestamp = Date.now();
let id = 0;

app.post('/orders', (req, res) => {
    const order = {
        id: id++,
        timestamp: Date.now() - startTimestamp,
        ...req.body,
    };

    console.log(order);
    res.status(201).send({ status: 'OK' });
});

app.listen(3000, () => {
    console.log('Server is running on port 3000');
});
