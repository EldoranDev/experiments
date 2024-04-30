export const VERSION = 1
export const HEADER_SIZE = 4

export class TCPMessage {

	public constructor(
		public command: number,
		public data: Buffer,
	) {
	}

	static unmarshallBinary(buffer: Buffer) {
		if (buffer.readInt8(0) != VERSION) {
			throw new Error("Version missmatch")
		}

		const command = buffer.readInt8(1)
		const length = buffer.readUint16BE(2)

		let data = Buffer.alloc(length)

		buffer.copy(data, 0, HEADER_SIZE, HEADER_SIZE + length)

		return new TCPMessage(command, data)
	}
}
