<script lang="ts">
    import { PUBLIC_GOOGLE_CLIENT_ID } from "$env/static/public";
    import {PUBLIC_BASE_URL} from "$env/static/public";
    import { onMount } from "svelte";

    const clientId = PUBLIC_GOOGLE_CLIENT_ID
    const baseURL = PUBLIC_BASE_URL
    //const baseURL = "http://localhost:5173"

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

            // const data = await response.json()

            // if(!response.ok){
            //     throw new Error(`Authentication failed : ${data.error}`)
            // }

            // console.log("JWT verified successfully. Authentication complete")
            // console.log(data.sub)
            // console.log(data.access_token)
            // console.log(data.email)

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