<script lang="ts">
    const pageName = "home"
    let responseText = ""

    async function serverHealthCheck(){
        const response = await fetch("http://localhost:3000/health")
        const text = await response.text()
        responseText = text
    }
</script>

<h1>You are in {pageName} page</h1>
<button onclick={serverHealthCheck}>Server Health Check</button>
<h1>{responseText}</h1>