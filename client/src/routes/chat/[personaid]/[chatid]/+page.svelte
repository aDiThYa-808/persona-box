<script lang="ts">
	import { replaceState } from '$app/navigation';
	import { page } from '$app/stores';
	import Chat from '$lib/components/chat.svelte';
	import type { NewChatResponse } from '$lib/types/chat.js';
	import type { Message } from '$lib/types/message';
	import { onMount } from 'svelte';
	import { personas } from '../../../../stores/personas.js';

	export let data;
	$: personaName = $personas.find((p) => p.persona_id === data.personaid)?.name;
	$: chatName = 'New Chat';

	$: title = personaName + " - "+ chatName

	$: chatID = data.chatid
	let messages: Message[] = [];
	let loading: boolean = false;

	onMount(async()=>{
		const firstMessage = $page.state?.message
		if(firstMessage){
			sendPrompt(firstMessage)
			console.log(firstMessage)
		} // else do something to show failure
	})

	async function sendPrompt(prompt: string) {
		messages = [...messages, { role: 'user', text: prompt }];

		loading = true;
		try {
			let body = (chatID != "new-chat" && chatID != "")
				? JSON.stringify({ session_id: chatID, persona_id: data.personaid, message: prompt })
				: JSON.stringify({ persona_id: data.personaid, message: prompt });

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

			const chatData: NewChatResponse = await response.json();

			if(chatData.session_id){
				replaceState(`/chat/${data.personaid}/${chatData.session_id}`, $page.state)
			}
			if(chatData.title){
				chatName = chatData.title
			}

			messages = [...messages, { role: 'assistant', text: chatData.response }];
		} catch (err) {
			messages = [...messages, { role: 'assistant', text: "Something went wrong on our end. Please try sending your message again." }];
		} finally {
			loading = false;
		}
	}
</script>

<Chat chatName={title} {messages} {sendPrompt} {loading} />
