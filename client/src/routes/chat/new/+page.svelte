<script lang="ts">
	import type { PersonaData, PersonaList } from '$lib/types/persona';
	import Create from '$lib/components/create.svelte';
	import { goto } from '$app/navigation';
	import { personas } from '../../../stores/personas';

	async function createPersona(data: PersonaData) {
		try {
			const res = await fetch(`/api/create-persona`, {
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
			personas.update((current) =>[...current,persona].sort((a,b)=>b.created_at.localeCompare(a.created_at)))

			await goto(`/chat/${persona.persona_id}`,{
				replaceState:true,
				noScroll:false
			})
			
		} catch (err) {
			console.log(err);
		}
	}
</script>

<Create {createPersona} />
