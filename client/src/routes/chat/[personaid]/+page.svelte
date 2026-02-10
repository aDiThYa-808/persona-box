<script lang="ts">
	import Chatlist from '$lib/components/chatlist.svelte';
	import Create from '$lib/components/create.svelte';
	import { chats, personas } from '../../../stores/store.js';
	import { goto } from '$app/navigation';
	import type { PersonaData, PersonaList } from '$lib/types/persona.js';

	export let data;
	$: personaid = data.personaid;
	$: personaName = $personas.find((p) => p.persona_id === data.personaid)?.name;
	$: chatSessions = data.chatSessions;
	$: if (chatSessions) chats.set(chatSessions);

	async function createPersona(data: PersonaData) {
		try {
			const res = await fetch(`/api/personas`, {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify(data)
			});

			if (!res.ok) {
				throw new Error(`HTTP Error. Status: ${res.status}`);
			}

			const persona: PersonaList = await res.json();

			// add the new persona to the personas store and sort the entire array using created_at
			personas.update((current) =>
				[...current, persona].sort((a, b) => b.created_at.localeCompare(a.created_at))
			);

			await goto(`/chat/${persona.persona_id}`, {
				noScroll: false,
				replaceState: false
			});
		} catch (err) {
			console.log(err);
		}
	}

	async function startNewChat(message: string) {
		goto(`/chat/${data.personaid}/new-chat`, {
			replaceState: false,
			noScroll: false,
			state: {
				message: message
			}
		});
	}

	function openChatSession(sessionid: string) {
		goto(`/chat/${data.personaid}/${sessionid}`, {
			replaceState: false,
			noScroll: false
		});
	}

	async function deleteChatSession(sessionid: string) {
		if (confirm('Are you sure you want to delete this chat session?')) {
			try {
				const res = await fetch(`/api/sessions/${personaid}/${sessionid}`, { method: 'DELETE' });
				const data = await res.json();
				if (!res.ok) {
					throw new Error(data);
				}
				console.log('deleted');
				chats.update((current) => current.filter((c) => c.session_id !== sessionid));
			} catch (err) {
				console.log(err);
			}
		}
	}
	
</script>

{#if data.personaid == 'new-persona'}
	<Create {createPersona} />
{:else}
	<Chatlist {personaName} {startNewChat} {openChatSession} {deleteChatSession} />
{/if}
