<script lang="ts">
	import type { Persona } from '$lib/types/persona';
	import { onMount } from 'svelte';

	export let name: string = 'John Doe';
	export let email: string = 'john@example.com';
	export let personas: Persona[];

	let isOpen = false;
	let expandedPersonas: string[] = [];
	let selectedChatId: string | null = null;

	// Check if we're on desktop on mount
	onMount(() => {
		const checkDesktop = () => {
			isOpen = window.innerWidth >= 1024; // lg breakpoint
		};

		checkDesktop();
		window.addEventListener('resize', checkDesktop);

		return () => window.removeEventListener('resize', checkDesktop);
	});

	function togglePersona(personaId: string) {
		if (expandedPersonas.includes(personaId)) {
			expandedPersonas = expandedPersonas.filter((id) => id !== personaId);
		} else {
			expandedPersonas = [...expandedPersonas, personaId];
		}
	}
</script>

<!-- Overlay for mobile -->
{#if isOpen}
	<button
		class="fixed inset-0 z-40 bg-black/50 lg:hidden"
		on:click={() => (isOpen = false)}
		tabindex="-1"
		aria-label="Close drawer"
	></button>
{/if}

<!-- Drawer -->
<aside
	class="fixed top-0 left-0 z-50 flex h-full w-72 flex-col border-r border-white/10 bg-card transition-transform duration-300 lg:translate-x-0 {isOpen
		? 'translate-x-0'
		: '-translate-x-full'}"
>
	<!-- Header -->
	<div class="flex items-center justify-between border-b border-white/10 px-4 py-4">
		<h1
			class="text-lg font-medium tracking-tight text-text"
			style="font-family: 'Pixelify Sans', sans-serif;"
		>
			Persona Box
		</h1>
		<button
			on:click={() => (isOpen = false)}
			class="rounded-md p-2 text-text transition-colors hover:bg-background lg:hidden"
			aria-label="Close drawer"
		>
			<svg
				width="20"
				height="20"
				viewBox="0 0 24 24"
				fill="none"
				stroke="currentColor"
				stroke-width="2"
			>
				<path d="M18 6L6 18M6 6l12 12" />
			</svg>
		</button>
	</div>

	<!-- New Persona Button -->
	<div class="border-b border-white/10 px-3 py-3">
		<button
			class="flex w-full items-center gap-3 rounded-lg px-3 py-2 text-sm text-text transition-colors hover:bg-background"
		>
			<svg
				width="18"
				height="18"
				viewBox="0 0 24 24"
				fill="none"
				stroke="currentColor"
				stroke-width="2"
			>
				<path d="M12 5v14M5 12h14" />
			</svg>
			New persona
		</button>
	</div>

	<!-- Personas List -->
	<div class="flex-1 overflow-y-auto px-3 py-3">
		<div class="px-3 py-2 text-xs font-medium text-text-muted">Recents</div>
		{#each personas as persona (persona.id)}
			<div class="mb-2">
				<!-- Persona Item -->
				<div
					class="group flex cursor-pointer items-center gap-2 rounded-lg px-3 py-2 transition-colors hover:bg-background"
				>
					<button
						on:click={() => togglePersona(persona.id)}
						class="flex flex-1 items-center gap-2 text-left text-sm text-text"
						aria-label="Toggle {persona.name}"
					>
						<svg
							width="14"
							height="14"
							viewBox="0 0 24 24"
							fill="none"
							stroke="currentColor"
							stroke-width="2"
							class="transition-transform {expandedPersonas.includes(persona.id)
								? 'rotate-90'
								: ''}"
						>
							<path d="M9 18l6-6-6-6" />
						</svg>
						<span class="truncate">{persona.name}</span>
					</button>
					<button
						class="rounded p-1 text-text opacity-0 transition-opacity group-hover:opacity-100 hover:bg-background"
						aria-label="Persona options"
					>
						<svg width="16" height="16" viewBox="0 0 24 24" fill="currentColor">
							<circle cx="12" cy="5" r="2" />
							<circle cx="12" cy="12" r="2" />
							<circle cx="12" cy="19" r="2" />
						</svg>
					</button>
				</div>

				<!-- Chats List (collapsible) -->
				{#if expandedPersonas.includes(persona.id)}
					<div class="mt-1 ml-6 space-y-1">
						{#each persona.chats as chat (chat.id)}
							<div
								class="group flex cursor-pointer items-center gap-2 rounded-lg px-3 py-1.5 transition-colors hover:bg-background {selectedChatId ===
								chat.id
									? 'bg-background-light'
									: ''}"
								role="button"
								tabindex="0"
								aria-label="Open chat: {chat.title}"
							>
								<span class="flex-1 truncate text-sm text-text-muted">{chat.title}</span>
								<button
									class="rounded p-1 text-text opacity-0 transition-opacity group-hover:opacity-100 hover:bg-background"
									aria-label="Chat options"
								>
									<svg width="14" height="14" viewBox="0 0 24 24" fill="currentColor">
										<circle cx="12" cy="5" r="2" />
										<circle cx="12" cy="12" r="2" />
										<circle cx="12" cy="19" r="2" />
									</svg>
								</button>
							</div>
						{/each}
					</div>
				{/if}
			</div>
		{/each}
	</div>

	<!-- User Section -->
	<div class="border-t border-white/10 p-3">
		<div
			class="flex cursor-pointer items-center gap-3 rounded-lg px-3 py-2 transition-colors hover:bg-background"
			role="button"
			tabindex="0"
			aria-label="User profile"
		>
			<div class="flex h-8 w-8 shrink-0 items-center justify-center rounded-full bg-background">
				<span class="text-sm text-text">{name.charAt(0).toUpperCase()}</span>
			</div>
			<div class="min-w-0 flex-1">
				<div class="truncate text-sm font-medium text-text">{name}</div>
				<div class="truncate text-xs text-text-muted">{email}</div>
			</div>
			<button
				class="rounded p-1 text-text transition-colors hover:bg-background"
				aria-label="Logout"
			>
				<svg
					width="16"
					height="16"
					viewBox="0 0 24 24"
					fill="none"
					stroke="currentColor"
					stroke-width="2"
				>
					<path d="M9 21H5a2 2 0 01-2-2V5a2 2 0 012-2h4M16 17l5-5-5-5M21 12H9" />
				</svg>
			</button>
		</div>
	</div>
</aside>

<!-- Toggle button for mobile (outside drawer) -->
{#if !isOpen}
	<button
		on:click={() => (isOpen = true)}
		class="fixed top-4 left-4 z-40 rounded-md border border-white/10 bg-card p-2 text-text transition-colors hover:bg-background lg:hidden"
		aria-label="Open drawer"
	>
		<svg
			width="20"
			height="20"
			viewBox="0 0 24 24"
			fill="none"
			stroke="currentColor"
			stroke-width="2"
		>
			<path d="M3 12h18M3 6h18M3 18h18" />
		</svg>
	</button>
{/if}
