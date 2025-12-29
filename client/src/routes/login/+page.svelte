<script lang="ts">
	import { goto } from '$app/navigation';
	import { PUBLIC_GOOGLE_CLIENT_ID } from '$env/static/public';
	import { error } from '@sveltejs/kit';
	import { onMount } from 'svelte';

	const clientId = PUBLIC_GOOGLE_CLIENT_ID;

	onMount(() => {
		window.google.accounts.id.initialize({
			client_id: clientId,
			callback: handleResponse
		});

		window.google.accounts.id.renderButton(document.getElementById('googleSignIn')!, {
			theme: 'outline',
			size: 'large'
		});
	});

	function handleResponse(response: { credential: string }) {
		const idToken = response.credential;
		verifyJWT(idToken);
	}

	async function verifyJWT(jwt: string) {
		try {
			const response = await fetch(`/api/auth/login`, {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify({
					id_token: jwt
				})
			});

			if (!response.ok) {
				throw error(500, 'Login failed');
			}

			await goto('/chat', {
				replaceState: true,
				noScroll: false
			});
		} catch (err: unknown) {
			console.log((err as Error).message);
		}
	}
</script>

<svelte:head>
	<link rel="preconnect" href="https://fonts.googleapis.com" />
	<link rel="preconnect" href="https://fonts.gstatic.com" crossorigin="anonymous" />
	<link
		href="https://fonts.googleapis.com/css2?family=Pixelify+Sans:wght@400..700&display=swap"
		rel="stylesheet"
	/>
</svelte:head>

<div class="flex min-h-screen flex-col bg-background text-text">
	<header class="px-8 py-6">
		<a
			href="/"
			class="text-lg font-medium tracking-tight transition-opacity hover:opacity-80"
			style="font-family: 'Pixelify Sans', sans-serif;"
		>
			PersonaBox
		</a>
	</header>
	<main class="flex flex-1 items-center justify-center px-8">
		<div class="w-full max-w-sm">
			<div class="mb-8 text-center">
				<h1 class="mb-2 text-3xl font-semibold">Welcome back</h1>
				<p class="text-sm text-text-muted">Sign in to continue to PersonaBox</p>
			</div>
			<div class="rounded-xl border border-white/10 bg-card p-8">
				<div id="googleSignIn" class="flex justify-center"></div>
			</div>
			<p class="mt-6 text-center text-xs text-white/40">
				By continuing, you agree to our Terms of Service and Privacy Policy
			</p>
		</div>
	</main>
</div>
