<script lang="ts">
</script>

<svelte:head>
	<link rel="preconnect" href="https://fonts.googleapis.com" />
	<link rel="preconnect" href="https://fonts.gstatic.com" crossorigin="anonymous" />
	<link
		href="https://fonts.googleapis.com/css2?family=Pixelify+Sans:wght@400..700&display=swap"
		rel="stylesheet"
	/>
</svelte:head>

<div
	class="flex min-h-screen flex-col bg-background text-text"
	style="font-family: 'Pixelify Sans', sans-serif;"
>
	<header class="flex items-center justify-between px-8 py-6">
		<div
			class="text-lg font-medium tracking-tight"
			style="font-family: 'Pixelify Sans', sans-serif;"
		>
			Persona Box
		</div>
		<a
			href="/login"
			class="rounded-md px-4 py-2 text-[15px] transition-colors hover:bg-card"
			style="font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', sans-serif;"
		>
			Login
		</a>
	</header>
	<main class="flex flex-1 items-center justify-center px-8 pb-32">
		<div
			class="max-w-2xl text-center"
			style="font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', sans-serif;"
		>
			<p class="mb-3 text-sm text-text-muted">Introducing PersonaBox</p>
			<h1 class="mb-8 text-5xl leading-tight font-semibold tracking-tight md:text-6xl">
				Create custom AI personas
			</h1>
			<a
				href="/chat"
				class="inline-block rounded-lg bg-text px-6 py-3 font-medium text-background transition-colors hover:bg-text/90"
			>
				Try PersonaBox
			</a>
		</div>
	</main>
</div>
