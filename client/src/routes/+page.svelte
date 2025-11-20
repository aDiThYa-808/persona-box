<script lang="ts">
    import {PUBLIC_AWS_INVOKE_URL} from "$env/static/public"

    const apiBaseUrl = PUBLIC_AWS_INVOKE_URL

    const pageName = "home"
    let responseText = ""

    async function serverHealthCheck(){
        const response = await fetch(`${apiBaseUrl}/health`)
        const text = await response.text()
        responseText = text
    }
</script>

<h1>You are in {pageName} page</h1>
<button onclick={serverHealthCheck}>Server Health Check</button>
<h1>{responseText}</h1>
<br><br>
<h1>THE FRONTEND IS STILL A WORK IN PROGRESS. DONT JUDGE ME SEEING THIS!!!!</h1>