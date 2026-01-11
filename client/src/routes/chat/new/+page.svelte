<script lang="ts">
	import type { PersonaData } from '$lib/types/persona';
	import Create from '$lib/components/create.svelte';
	import { goto } from '$app/navigation';

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

			const resData = await res.json();

			await goto(`/chat/${resData.persona_id}`,{
				replaceState:true,
				noScroll:false
			})
			
		} catch (err) {
			console.log(err);
		}
	}
</script>

<Create {createPersona} />
