<script lang="ts">
	import Chatlist from '$lib/components/chatlist.svelte';
	import Create from '$lib/components/create.svelte';
	import Modal from '$lib/components/modal.svelte';
	import { chats, personas } from '../../../stores/store.js';
	import { goto } from '$app/navigation';
	import type { PersonaData, PersonaList } from '$lib/types/persona.js';

	export let data;
	$: personaid = data.personaid;
	$: personaName = $personas.find((p) => p.persona_id === data.personaid)?.name;
	$: chatSessions = data.chatSessions;
	$: if (chatSessions) chats.set(chatSessions);

	let showConfirmModal = false;
	let deleteSessionId = '' // id of the session that has to be deleted

	let createPersonaLoading = false
	let createChatLoading = false

	async function createPersona(data: PersonaData) {
		createPersonaLoading = true
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
		}finally{
			createPersonaLoading = false
		}
	}

	async function startNewChat(message: string) {
		createChatLoading = true
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

	function deleteChatSession(sessionid: string) {
		deleteSessionId = sessionid
		showConfirmModal = true;
	}

	async function confirmDelete() {
		try {
			const res = await fetch(`/api/sessions/${personaid}/${deleteSessionId}`, { method: 'DELETE' });
			const data = await res.json();
			if (!res.ok) {
				throw new Error(data);
			}
			console.log('deleted');
			chats.update((current) => current.filter((c) => c.session_id !== deleteSessionId));
		} catch (err) {
			console.log(err);
		} finally{
			deleteSessionId = ""
			showConfirmModal = false
		}
	}

	function cancelDelete(){
		deleteSessionId = ""
		showConfirmModal = false
	}
</script>

{#if data.personaid == 'new-persona'}
	<Create {createPersona} isLoading= {createPersonaLoading}/>
{:else}
	<Chatlist {personaName} {startNewChat} {openChatSession} {deleteChatSession} isLoading = {createChatLoading} />
{/if}

<Modal 
	isOpen={showConfirmModal}
	title="Confirm"
	message="Are you sure you want to delete this chat session?"
	confirmButtonText="Delete"
	cancelButtonText="Cancel"
	onConfirm={confirmDelete}
	onCancel={cancelDelete}
/>
