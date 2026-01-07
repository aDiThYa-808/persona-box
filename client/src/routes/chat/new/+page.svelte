<script lang="ts">
	import type { PersonaData } from '$lib/types/persona';
	import Create from '$lib/components/create.svelte';

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

			console.log(JSON.stringify(resData));
		} catch (err) {
			console.log(err);
		}
	}
</script>

<Create {createPersona} />
