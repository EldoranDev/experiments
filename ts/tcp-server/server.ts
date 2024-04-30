import { Server, createServer } from 'node:net'
import { TCPMessage } from './tcp/message'
import { Command } from './command'

const server: Server = createServer((socket) => {
	socket.on('data', (buffer: Buffer) => {
		const message = TCPMessage.unmarshallBinary(buffer)

		console.log(`${message.command} - ${message.data.toString()}`)
	})
})

server.listen(3000, '127.0.0.1')
