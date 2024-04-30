import { AbstractRunner } from './runner'

export class GoRunner extends AbstractRunner {
	constructor(
		app: string
	) {
		super('go', app)
	}

	protected getCommand(): string {
		return 'go'
	}

	protected getArguments(): string[] {
		return ['run', '.']
	}
}
