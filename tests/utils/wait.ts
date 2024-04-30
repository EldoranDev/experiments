export async function wait(delay: number): Promise<void> {
	await new Promise<void>((resolve) => {
		setTimeout(() => {
			resolve();
		}, delay)
	})
}
