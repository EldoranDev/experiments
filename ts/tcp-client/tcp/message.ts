import { Command } from "../command"

const VERSION = 1
const HEADER_SIZE = 4

export class TCPMessage {

	constructor(
		private command: number,
		private data: any
	) {
	}

	marshallBinary(): Buffer {
		let data = Buffer.from(this.data, 'utf8')

		let buffer = Buffer.alloc(HEADER_SIZE + data.byteLength)

		buffer.writeUint8(VERSION, 0)
		buffer.writeUint8(this.command, 1)
		buffer.writeUInt16BE(data.byteLength, 2)

		data.copy(buffer, HEADER_SIZE)

		return buffer
	}
}
