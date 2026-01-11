<script lang="ts">
	import Sidebar from '$lib/components/sidebar.svelte';
	import type { User } from '$lib/types/user';
	import type { PersonaList } from '$lib/types/persona';
	import { personas } from '../../stores/personas';
	import { onMount } from 'svelte';

	export let data: User;
	let fetchedPersonas: PersonaList[]

	onMount(async ()=>{
		fetchedPersonas = await getUsersPersonas()
		personas.set(fetchedPersonas)
	})

	async function getUsersPersonas():Promise<PersonaList[]>{
		try{

			const resp = await fetch(`/api/personas`)
			if(!resp.ok){
				const data = await resp.json()
				throw new Error(`HTTP Error. status: ${resp.status}; message: ${data.error}`)

			}
			const personas = await resp.json()
			return personas
		}catch(err){
			console.log(err)
			return []
		}
	} 
</script>

<Sidebar name={data.name} email={data.email}/>
<slot />
