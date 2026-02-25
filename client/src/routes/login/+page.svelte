<script lang="ts">
	/**
	 * Google OAuth sign-in page.
	 * Renders Google Sign-In button, handles authentication, and redirects to new persona page.
	 * Shows loading state during authentication process.
	 *
	 * Initializes Google OAuth on mount and handles ID token verification through /api/auth/login.
	 */
	import { goto } from '$app/navigation';
	import { PUBLIC_GOOGLE_CLIENT_ID } from '$env/static/public';
	import { error } from '@sveltejs/kit';
	import { onMount } from 'svelte';

	const clientId = PUBLIC_GOOGLE_CLIENT_ID;
	let isLoading = false;

	onMount(() => {
		window.google.accounts.id.initialize({
			client_id: clientId,
			callback: handleResponse
		});

		window.google.accounts.id.renderButton(document.getElementById('googleSignIn')!, {
			theme: 'filled_blue',
			size: 'large',
			width: 250,
			text: 'continue_with',
			shape: 'square',
			logo_alignment: 'left'
		});
	});

	function handleResponse(response: { credential: string }) {
		const idToken = response.credential;
		verifyJWT(idToken);
	}

	async function verifyJWT(jwt: string) {
		isLoading = true;
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

			await goto('/chat/new-persona', {
				replaceState: false,
				noScroll: false
			});
		} catch (err: unknown) {
			console.log((err as Error).message);
		} finally {
			isLoading = false;
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
	<title>Sign in - Personabox</title>
</svelte:head>

<div class="flex min-h-screen flex-col bg-background text-text">
	<!-- Header -->
	<header class="px-8 py-8">
		<a
			href="/"
			class="text-xl font-semibold tracking-tight transition-opacity hover:opacity-80"
			style="font-family: 'Pixelify Sans', sans-serif;"
		>
			Personabox
		</a>
	</header>

	<!-- Main Content -->
	<main class="flex flex-1 items-center justify-center px-6 py-12">
		<div class="w-full max-w-md">
			<!-- Header Text -->
			<div class="mb-10 text-center">
				<h1 class="mb-3 text-3xl font-bold tracking-tight md:text-4xl">Welcome back</h1>
				<p class="text-sm text-text-muted">
					{isLoading ? 'Signing you in...' : 'Sign in to create and chat with AI personas'}
				</p>
			</div>

			<!-- Sign In Card -->
			<div class="rounded-lg border border-white/10 bg-card p-10 shadow-2xl shadow-black/10">
				{#if isLoading}
					<!-- Loading State -->
					<div class="flex flex-col items-center justify-center py-4">
						<div
							class="mb-3 h-10 w-10 animate-spin rounded-full border-4 border-white/10 border-t-text"
						></div>
						<p class="text-sm text-text-muted">Signing in...</p>
					</div>
				{:else}
					<!-- Google Sign In Button -->
					<div id="googleSignIn" class="flex justify-center"></div>
				{/if}
			</div>

			<!-- Terms Notice -->
			<p class="mt-8 text-center text-sm leading-relaxed text-text-muted">
				By continuing, you agree to our
				<a href="/terms" class="underline transition-colors hover:text-text">Terms of Service</a>
				and
				<a href="/privacy" class="underline transition-colors hover:text-text">Privacy Policy</a>
			</p>
		</div>
	</main>

	<!-- Footer -->
	<footer class="border-t border-white/10 px-8 py-8">
		<div
			class="mx-auto flex w-full max-w-6xl flex-col items-center justify-between gap-4 text-center sm:flex-row sm:text-left"
		>
			<p class="text-sm text-text-muted">© 2024 PersonaBox</p>
			<div class="flex gap-6 text-sm">
				<a href="/terms" class="font-medium text-text-muted transition-colors hover:text-text"
					>Terms</a
				>
				<a href="/privacy" class="font-medium text-text-muted transition-colors hover:text-text"
					>Privacy</a
				>
				<a href="/contact" class="font-medium text-text-muted transition-colors hover:text-text"
					>Contact</a
				>
			</div>
		</div>
	</footer>
</div>
