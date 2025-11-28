<script lang="ts">
	export let data

    const pageName = "dashboard"
    const email = data.user.email
</script>

<h1>You are in {pageName} page</h1>
<p>Logged in using : {email}</p>