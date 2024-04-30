import { AbstractRunner } from './runner'

export class TsRunner extends AbstractRunner {
	public constructor(
		app: string
	) {
		super('ts', app)
	}

	protected getCommand(): string {
		// This fixes some issues with the way nodejs
		// is installed on the system currently
		if (process.platform === 'win32') {
			return 'npm.cmd'
		}

		return 'npm'
	}

	protected getArguments(): string[] {
		return ['run', 'start']
	}
}
