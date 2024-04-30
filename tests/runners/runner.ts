import { resolve } from 'node:path'
import {
	execSync,
	execFile,
	ChildProcess,
} from 'node:child_process'
import { wait } from '../utils/wait'

export abstract class AbstractRunner {
	public cwd: string

	protected running: boolean

	protected child: ChildProcess

	constructor(
		public readonly lang: string,
		public readonly project: string,
	) {
		this.cwd = resolve('..', lang, project)
	}

	protected abstract getCommand(): string

	protected abstract getArguments(): Array<string>

	/**
	 * Runs the given command
	 */
	public run(): string {
		if (this.running) {
			throw new Error('Runner is already active')
		}

		const cmd = `${this.getCommand()} ${this.getArguments().join(' ')}`

		// TODO: handle execution errors here so that not the full output is shown
		const output = execSync(cmd, {
			cwd: this.cwd
		})

		return output.toString()
	}

	public async start() {
		if (this.running) {
			throw new Error('Runner is already active')
		}

		this.running = true;

		this.child = execFile(
			this.getCommand(),
			this.getArguments(),
			{
				cwd: this.cwd,
				shell: true,
			}
		)

		this.child.on('error', (err) => {
			console.error(err)
		})

		return wait(400)
	}

	public async stop() {
		if (!this.running || !this.child.pid) {
			return;
		}

		this.running = false

		// Apparently it doesn't work just calling `kill` on the child process
		// because this also is running on windows
		if (process.platform == "win32") {
			execSync(`taskkill /PID ${this.child.pid} /T /F`)
		} else {
			process.kill(this.child.pid!, 'SIGTERM')
		}

		return Promise.resolve()
	}
}
