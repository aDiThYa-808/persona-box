<script lang="ts">
	/**
	 * Chat interface component with message display and input.
	 * Auto-scrolls to bottom on new messages and handles message limits.
	 *
	 * @component
	 * @prop {string} chatName - Display name shown in the header
	 * @prop {boolean} loading - Shows typing indicator when true
	 * @prop {boolean} limitReached - Disables input and shows limit warning
	 * @prop {function} sendPrompt - Callback to send messages, receives prompt string
	 */
	import { messages } from '../../stores/store';
	import { tick } from 'svelte';

	export let chatName: string;
	export let loading: boolean;
	export let limitReached: boolean;
	export let sendPrompt: (prompt: string) => void;

	let prompt = '';
	let messagesContainer: HTMLDivElement;

	$: if ($messages || loading) {
		scrollToBottom();
	}

	async function scrollToBottom() {
		await tick();
		if (messagesContainer) {
			messagesContainer.scrollTop = messagesContainer.scrollHeight;
		}
	}

	function handleSend() {
		if (!prompt.trim() || loading) return;
		sendPrompt(prompt);
		prompt = '';

		const textarea = document.querySelector('textarea');
		if (textarea) {
			textarea.style.height = 'auto';
		}
	}
</script>

<div class="flex h-dvh flex-col bg-background text-text lg:ml-72">
	<!-- Chat Header -->
	<div class="sticky top-0 z-10 border-b border-white/10 bg-card px-4 py-4 pl-16 sm:px-6 lg:pl-6">
		<h2 class="truncate text-lg font-medium text-text">{chatName}</h2>
	</div>

	<!-- Chat Messages -->
	<div bind:this={messagesContainer} class="flex-1 overflow-y-auto px-4 py-4 sm:px-6 sm:py-8">
		<div class="mx-auto max-w-3xl pb-4">
			{#each $messages.sort((a, b) => a.created_at.localeCompare(b.created_at)) as msg, i (i)}
				<div class="flex {msg.role === 'user' ? 'justify-end' : 'justify-start'} mb-4 sm:mb-6">
					<div
						class={`max-w-[90%] px-3 py-2 text-base leading-relaxed break-words sm:max-w-[80%] sm:px-4 sm:py-2.5 sm:text-lg ${
							msg.role === 'user' ? 'rounded-2xl bg-card text-text' : 'text-text'
						}`}
					>
						{msg.message}
					</div>
				</div>
			{/each}
			{#if loading}
				<div class="mb-4 flex justify-start sm:mb-6">
					<div class="px-1 py-2 text-base text-text-muted sm:text-lg">
						<span class="inline-flex items-center gap-1">
							<span class="animate-pulse">●</span>
							<span class="animate-pulse delay-75">●</span>
							<span class="animate-pulse delay-150">●</span>
						</span>
					</div>
				</div>
			{/if}
		</div>
	</div>

	<!-- Input - Sticky to bottom of container -->
	<div class="sticky bottom-0 border-t border-white/10 bg-background px-4 py-3 sm:px-6 sm:py-4">
		{#if limitReached}
			<!-- Limit Reached Banner -->
			<div
				class="mx-auto mb-3 max-w-3xl rounded-xl border border-red-500/20 bg-red-500/10 px-4 py-3 sm:px-6"
			>
				<div class="flex items-start gap-3">
					<svg
						class="mt-0.5 h-5 w-5 shrink-0 text-red-400"
						viewBox="0 0 24 24"
						fill="none"
						stroke="currentColor"
						stroke-width="2"
					>
						<path d="M12 9v4m0 4h.01M21 12a9 9 0 1 1-18 0 9 9 0 0 1 18 0z" />
					</svg>
					<div class="flex-1 text-sm sm:text-base">
						<p class="font-medium text-red-400">Chat limit reached</p>
						<p class="mt-1 text-red-300/80">
							You've reached the message limit for this conversation. Please create a new chat to
							continue.
						</p>
					</div>
				</div>
			</div>
		{/if}

		<div
			class="mx-auto flex max-w-3xl items-end gap-2 rounded-2xl border border-white/10 bg-card p-1.5 shadow-lg sm:p-2 {limitReached
				? 'opacity-50'
				: ''}"
		>
			<textarea
				bind:value={prompt}
				rows="1"
				placeholder={limitReached ? 'Chat limit reached...' : 'Type a message...'}
				disabled={limitReached}
				class="max-h-[200px] min-h-[44px] flex-1 resize-none border-0 bg-transparent px-3 py-2 text-base text-text placeholder-text-muted focus:ring-0 focus:outline-none disabled:cursor-not-allowed sm:py-2.5 sm:text-lg"
				style="overflow-y: auto; field-sizing: content;"
				on:input={(e) => {
					e.currentTarget.style.height = 'auto';
					e.currentTarget.style.height = Math.min(e.currentTarget.scrollHeight, 200) + 'px';
				}}
				on:keydown={(e) => e.key === 'Enter' && !e.shiftKey && (e.preventDefault(), handleSend())}
			></textarea>
			<button
				on:click={handleSend}
				class="shrink-0 rounded-lg bg-text p-2 text-background transition hover:bg-text/90 disabled:cursor-not-allowed disabled:opacity-50 sm:p-2.5"
				disabled={loading || !prompt.trim() || limitReached}
				aria-label="Send message"
			>
				<svg
					width="18"
					height="18"
					viewBox="0 0 24 24"
					fill="none"
					stroke="currentColor"
					stroke-width="2"
					stroke-linecap="round"
					stroke-linejoin="round"
					class="h-4 w-4 sm:h-[18px] sm:w-[18px]"
				>
					<path d="M12 19V5M5 12l7-7 7 7" />
				</svg>
			</button>
		</div>
	</div>
</div>
