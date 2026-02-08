<script lang="ts">
	import { replaceState } from '$app/navigation';
	import { page } from '$app/stores';
	import Chat from '$lib/components/chat.svelte';
	import type { NewChatResponse } from '$lib/types/chat.js';
	import { onMount } from 'svelte';
	import { chats, messages, personas } from '../../../../stores/store.js';

	export let data;

	$: chatID = data.chatid;

	$: personaName = $personas.find((p) => p.persona_id === data.personaid)?.name;
	$: chatName = $chats.find((c) => c.session_id === chatID)?.title;
	$: title = personaName + ' - ' + chatName;


	$: sessionMessages = data.messages;
	$: if(sessionMessages){
		messages.set(sessionMessages);
	} // else show some error

	let loading: boolean = false;

	onMount(async () => {
		const firstMessage = $page.state?.message;
		if (firstMessage) {
			sendPrompt(firstMessage);
		} // else do something to show failure
	});

	async function sendPrompt(prompt: string) {
		const createdAt = new Date().toISOString()
		messages.update((current)=> [...current,{role:"user",message:prompt,created_at: createdAt}])
		loading = true;
		try {
			let body =
				chatID != 'new-chat' && chatID != ''
					? JSON.stringify({ session_id: chatID, persona_id: data.personaid, message: prompt, created_at:createdAt })
					: JSON.stringify({ persona_id: data.personaid, message: prompt,created_at : createdAt });

			console.log(body)
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

			if (chatData.session_id) {
				replaceState(`/chat/${data.personaid}/${chatData.session_id}`, $page.state);
				chatID = chatData.session_id
			}
			if (chatData.title) {
				chatName = chatData.title;
			}

			messages.update((current)=>[...current, { role: 'assistant', message: chatData.response,created_at:chatData.timestamp }])
		} catch (err) {
			messages.update((current)=> [
				...current,
				{
					role: 'assistant',
					message: 'Something went wrong on our end. Please try sending your message again.',
					created_at: createdAt
				}
			]);
		} finally {
			loading = false;
		}
	}
</script>

<Chat chatName={title} {sendPrompt} {loading} />
