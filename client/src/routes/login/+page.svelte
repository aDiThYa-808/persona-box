<script lang="ts">
	import { goto } from "$app/navigation";
    import { PUBLIC_GOOGLE_CLIENT_ID } from "$env/static/public";
    import {PUBLIC_BASE_URL} from "$env/static/public";
    import {PUBLIC_ENVIRONMENT} from "$env/static/public"
	import { error} from "@sveltejs/kit";
    import { onMount } from "svelte";

    const clientId = PUBLIC_GOOGLE_CLIENT_ID
    const baseURL = (PUBLIC_ENVIRONMENT.toLowerCase() == "local")? 'http://localhost:5173' : PUBLIC_BASE_URL

    onMount(()=>{
        window.google.accounts.id.initialize({
            client_id: clientId,
            callback: handleResponse
        })

        window.google.accounts.id.renderButton(
            document.getElementById("googleSignIn")!,
            {"theme":"outline","size":"large"}
        )
    })

    function handleResponse(response:{credential:string}){
        const idToken = response.credential
        verifyJWT(idToken)
    }

    async function verifyJWT(jwt: string){
        try{

            const response = await fetch(`${baseURL}/api/auth/login`,{
                method:'POST',
                headers: {
                    'Content-Type':'application/json'
                },
                body: JSON.stringify({
                    'id_token': jwt
                })
            })

           if(!response.ok){
            throw error(500,'Login failed')
           }

           await goto('/dashboard',{
            replaceState:true,
            noScroll:false
           })

        }catch(err:unknown){
            console.log((err as Error).message)
        }
    }

</script>

<main class="min-h-screen flex items-center justify-center text-white">
    <div class="bg-[#1a1a1a] border border-[#2a2a2a] rounded-xl p-8 w-[280px] text-center">
        <h1 class="text-xl font-medium mb-6">Login</h1>
        <div id="googleSignIn" class="flex justify-center"></div>
    </div>
</main>