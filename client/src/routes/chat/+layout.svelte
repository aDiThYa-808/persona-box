<script lang="ts">
	import Sidebar from '$lib/components/sidebar.svelte';
	import type { User } from '$lib/types/user';
	import type { PersonaList } from '$lib/types/persona';
	import { chats, personas } from '../../stores/store';
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';

	export let data: User;
	let fetchedPersonas: PersonaList[];

	// fetch users personas and set the personas store on mount
	onMount(async () => {
		fetchedPersonas = await getUsersPersonas();

		//store a sorted personas array (newest to oldest)
		personas.set(fetchedPersonas.sort((a, b) => b.created_at.localeCompare(a.created_at)));
	});

	async function getUsersPersonas(): Promise<PersonaList[]> {
		try {
			const resp = await fetch(`/api/personas`, {
				method: 'GET'
			});
			if (!resp.ok) {
				const data = await resp.json();
				throw new Error(`HTTP Error. status: ${resp.status}; message: ${data.error}`);
			}
			const personas = await resp.json();
			return personas;
		} catch (err) {
			console.log(err);
			return [];
		}
	}

	async function deletePersona(personaid: string) {
		if (
			confirm(
				'Are you sure you want to delete this persona? This will also delete all the chat sessions and messages.'
			)
		) {
			try {
				const res = await fetch(`/api/personas/${personaid}`, { 
					method : 'DELETE' 
				});
				if (!res.ok) {
					let data = await res.json();
					throw new Error(data);
				}
				console.log('deleted');

				personas.update((current)=> current.filter((persona)=> persona.persona_id !== personaid))
				goto(`/chat/new-persona`,{
					replaceState: false,
					noScroll:false
				})
			} catch (err) {
				console.log(err);
			}
		}
	}
</script>

<Sidebar name={data.name} email={data.email} {deletePersona} />
<slot />
