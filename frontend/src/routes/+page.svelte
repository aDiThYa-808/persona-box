<script lang="ts">
    const pageName = "home"
    let responseText = ""

    async function serverHealthCheck(){
        const response = await fetch("https://0m6lq7o0ne.execute-api.eu-north-1.amazonaws.com/health")
        const text = await response.text()
        responseText = text
    }
</script>

<h1>You are in {pageName} page</h1>
<button onclick={serverHealthCheck}>Server Health Check</button>
<h1>{responseText}</h1>