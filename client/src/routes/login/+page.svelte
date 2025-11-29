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

<svelte:head>
  <link rel="preconnect" href="https://fonts.googleapis.com">
  <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin="anonymous">
  <link href="https://fonts.googleapis.com/css2?family=Pixelify+Sans:wght@400..700&display=swap" rel="stylesheet">
</svelte:head>

<div class="min-h-screen flex flex-col text-white">
  <header class="px-8 py-6">
    <a href="/" class="text-lg font-medium tracking-tight hover:opacity-80 transition-opacity" style="font-family: 'Pixelify Sans', sans-serif;">
      PersonaBox
    </a>
  </header>

  <main class="flex-1 flex items-center justify-center px-8">
    <div class="w-full max-w-sm">
      <div class="text-center mb-8">
        <h1 class="text-3xl font-semibold mb-2">Welcome back</h1>
        <p class="text-white/60 text-sm">Sign in to continue to PersonaBox</p>
      </div>

      <div class="bg-white/5 border border-white/10 rounded-xl p-8">
        <div id="googleSignIn" class="flex justify-center"></div>
      </div>

      <p class="text-center text-white/40 text-xs mt-6">
        By continuing, you agree to our Terms of Service and Privacy Policy
      </p>
    </div>
  </main>
</div>