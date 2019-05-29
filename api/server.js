const express = require('express')
const app = express()
const port = 8090

const { createLogger, format, transports } = require('winston')
const { combine, timestamp, label, printf } = format

const log = createLogger({
  format: combine(
    label({ label: 'fargate' }),
    timestamp(),
    printf(({ level, message, label, timestamp }) => {
      return `${timestamp} [${label}] ${level}: ${message}`;
    })
  ),
  transports: [ new transports.Console() ]
})

app.get('/', (req, res) =>
  res.send('Hello World!')
)

app.listen(port, () => log.info(`API listening on port: ${port}`))
