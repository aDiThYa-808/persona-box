<script lang="ts">
	import { goto } from '$app/navigation';
	import { resolve } from '$app/paths';
	import { onMount } from 'svelte';
	import { personas } from '../../stores/store';
	import Trash from '$lib/assets/icons/trash.png'

	export let name: string;
	export let email: string;
	export let deletePersona: (personaid: string)=> void

	let isOpen = false;

	// Check if we on desktop on mount
	onMount(() => {
		const checkDesktop = () => {
			isOpen = window.innerWidth >= 1024; // lg breakpoint
		};

		checkDesktop();
		window.addEventListener('resize', checkDesktop);

		return () => window.removeEventListener('resize', checkDesktop);
	});

	function goToPersonaPage(id: string) {
		goto(`/chat/${id}`,{
			replaceState:false,
			noScroll:false
		});
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
			on:click={() => {
				goto(resolve('/chat/new-persona'));
				isOpen = false;
			}}
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
		{#each $personas as persona (persona.persona_id)}
			<div class="mb-2">
				<div
					class="group flex cursor-pointer items-center gap-2 rounded-lg px-3 py-2 transition-colors hover:bg-background"
					role="button"
					tabindex="0"
					on:click={() => {
						goToPersonaPage(persona.persona_id);
						isOpen = false;
					}}
					on:keydown={(e) => {
						e.key === 'Enter' && goToPersonaPage(persona.persona_id);
						isOpen = false;
					}}
				>
					<span class="flex-1 truncate text-sm text-text">{persona.name}</span>
					<button
						class="rounded p-1 text-text opacity-0 transition-opacity group-hover:opacity-100 hover:bg-card"
						aria-label="Delete persona"
						on:click={(e) =>{
							e.stopPropagation()
							deletePersona(persona.persona_id)
						} }
					>
						<img src={Trash} alt="Delete" class="h-4 w-4">
					</button>
				</div>
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
