const express = require('express');
const amqp = require('amqplib');

const app = express();
const port = 3001;

app.post('/publish', async (req, res) => {
  const message = req.body.message;

  try {
    const connection = await amqp.connect('amqp://localhost:5672');
    const channel = await connection.createChannel();
    const exchange = 'logs';

    channel.assertExchange(exchange, 'fanout', { durable: false });
    channel.publish(exchange, '', Buffer.from(message));

    console.log(`Message sent: ${message}`);
    res.status(200).send('Message sent');
  } catch (error) {
    console.error(error);
    res.status(500).send('Internal Server Error');
  }
});

app.listen(port, () => {
  console.log(`Publisher listening at http://localhost:${port}`);
});
