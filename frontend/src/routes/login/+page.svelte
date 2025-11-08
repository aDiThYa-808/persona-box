<script lang="ts">
    import { PUBLIC_GOOGLE_CLIENT_ID } from "$env/static/public";
    import { onMount } from "svelte";

    const clientId = PUBLIC_GOOGLE_CLIENT_ID

    onMount(()=>{
        window.google.accounts.id.initialize({
            client_id: clientId,
            callback: handleCredentialResponse
        })

        window.google.accounts.id.renderButton(
            document.getElementById("googleSignIn")!,
            {"theme":"outline","size":"large"}
        )
    })

    function handleCredentialResponse(response:{credential:string}){
        const idToken = response.credential
        console.log("Google id token: ",idToken)
    }

</script>

<svelte:head>
    <script src="https://accounts.google.com/gsi/client" async defer></script>
</svelte:head>


<div id="googleSignIn"></div>