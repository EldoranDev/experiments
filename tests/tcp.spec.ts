import { describe, test, beforeAll } from 'vitest'

import { GoRunner, TsRunner, AbstractRunner } from './runners'

const clients: Array<[AbstractRunner]> = [
	[new TsRunner('tcp-client')],
	[new GoRunner('tcp-client')],
]

describe('Go TCP-Server works with ', () => {
	beforeAll(async () => {
		const runner = new GoRunner('tcp-server')

		await runner.start()

		return async () => {
			await runner.stop()
		}
	})

	test.each(clients)('TCP-Client ($lang)', (runner: AbstractRunner) => {
		runner.run()
	})
})

describe('TS TCP-Server works with ', () => {
	beforeAll(async () => {
		const runner = new TsRunner('tcp-server')

		await runner.start()

		// Give the server time to boot up

		return async () => {
			await runner.stop()
		}
	})

	test.each(clients)('TCP-Client ($lang)', async (runner: AbstractRunner) => {
		runner.run()
	})
})
