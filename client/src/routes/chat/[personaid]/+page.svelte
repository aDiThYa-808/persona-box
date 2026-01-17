<script lang="ts">
	import type {Chat} from '$lib/types/chat.ts'; 
    import Chatlist from '$lib/components/chatlist.svelte';
	import { personas } from '../../../stores/personas.js';

    export let data
    $: personaName = $personas.find(p=>p.persona_id === data.personaid)?.name
    let chatList: Chat[] = []

    async function startNewChat(){
        console.log("Work in progress, lol")
        // try{
        //     const body = JSON.stringify({"persona_id":data.personaid})
        //     const res = await fetch('/api/start-chat',{
        //         method:'POST',
        //         headers: {
        //             'Content-Type':'application/json'
        //         },
        //         body: body
        //     })

        //     if(!res.ok){
        //         const data = await res.json()
        //         throw new Error(`HTTP error. status:${res.status}; message:${data.error}`)
        //     }

        //     // navigate to chat/[personaid]/[chatid] . use the chatid from response
        // }catch(err){
        //     console.log(err)
        // }
    }
</script>
<Chatlist personaName={personaName} chatList={chatList} {startNewChat}/>

