<script lang="ts">
	import type {Chat, NewChatResponse} from '$lib/types/chat.ts'; 
    import Chatlist from '$lib/components/chatlist.svelte';
	import { personas } from '../../../stores/personas.js';
	import { goto } from '$app/navigation';
	import { resolve } from '$app/paths';

    export let data
    $: personaName = $personas.find(p=>p.persona_id === data.personaid)?.name
    let chatList: Chat[] = []

    async function startNewChat(message: string){
        try{
            const body = JSON.stringify({"persona_id":data.personaid, "message":message})
            const res = await fetch('/api/chat',{
                method:'POST',
                headers: {
                    'Content-Type':'application/json'
                },
                body: body
            })

            if(!res.ok){
                const errorData = await res.json()
                throw new Error(`HTTP error. status:${res.status}; message:${errorData.error}`)
            }

            const chatResponse : NewChatResponse = await res.json()
            console.log(chatResponse)

            goto(resolve(`/chat/${data.personaid}/${chatResponse.session_id}`),{
				replaceState:true,
				noScroll:false
			})

        }catch(err){
            console.log(err)
        }
    }
</script>
<Chatlist personaName={personaName} chatList={chatList} {startNewChat}/>

