import { Socket } from 'node:net'

import { TCPMessage } from './tcp/message'
import { Command } from './command'

const client = new Socket()

client.connect(3000, '127.0.0.1', () => {
	const message = new TCPMessage(
		Command.Echo,
		"Hello Server!"
	)

	const data = message.marshallBinary()

	client.write(data)

	client.end()
});
