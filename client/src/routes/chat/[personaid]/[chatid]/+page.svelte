<script lang="ts">
	import Chat from '$lib/components/chat.svelte';
	import type { Message } from '$lib/types/message';

    export let data

	let chatName = 'This is the chat name';
	let messages: Message[] = [];
	let loading: boolean = false;

	async function sendPrompt(prompt: string) {
		messages = [...messages, { role: 'user', text: prompt }];

		loading = true;
		try {
            let body = JSON.stringify({"session_id": data.chatid,"persona_id":data.personaid,"message":prompt})
			const response = await fetch(`/api/chat`, {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: body
			});

			if (!response.ok) {
				throw new Error(`HTTP error. Status: ${response.status}`);
			}

			const chatData = await response.json();
        

			messages = [...messages, { role: 'assistant', text: chatData.response }];
		} catch (err) {
			messages = [...messages, { role: 'assistant', text: (err as Error).message }];
		} finally {
			loading = false;
		}
	}
</script>

<Chat {chatName} {messages} {sendPrompt} {loading} />