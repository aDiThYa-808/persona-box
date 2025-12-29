<script lang="ts">
	import Chat from '$lib/components/chat.svelte';
	import type { Message } from '$lib/types/message';

	let chatName = 'This is the chat name';
	let messages: Message[] = [];
	let loading: boolean = false;

	async function sendPrompt(prompt: string) {
		messages = [...messages, { role: 'user', text: prompt }];

		loading = true;
		try {
			const response = await fetch(`/api/chat`, {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify({
					message: prompt
				})
			});

			if (!response.ok) {
				throw new Error(`HTTP error. Status: ${response.status}`);
			}

			const data = await response.json();

			messages = [...messages, { role: 'assistant', text: data.response }];
		} catch (err) {
			messages = [...messages, { role: 'assistant', text: (err as Error).message }];
		} finally {
			loading = false;
		}
	}
</script>

<Chat {chatName} {messages} {sendPrompt} {loading} />
