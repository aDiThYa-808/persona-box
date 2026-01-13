<script lang="ts">
	import type {Chat} from '$lib/types/chat.ts'; 
    import Chatlist from '$lib/components/chatlist.svelte';
	import { personas } from '../../../stores/personas.js';

    export let data
    $: personaName = $personas.find(p=>p.persona_id === data.personaid)?.name
    let chatList: Chat[] = []
</script>

<Chatlist personaName={personaName} chatList={chatList}/>

