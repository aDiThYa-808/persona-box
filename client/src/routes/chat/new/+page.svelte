<script lang="ts">
	import type { PersonaData } from '$lib/types/persona';
	import Create from '$lib/components/create.svelte';

	function createPersona(data: PersonaData) {
		console.log(data);
	}
</script>

<Create {createPersona} />
