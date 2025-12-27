<script lang="ts">
	import type { Message } from '$lib/types/message';
	import { PUBLIC_AWS_INVOKE_URL } from '$env/static/public';

	const apiBaseUrl = PUBLIC_AWS_INVOKE_URL;

	let prompt = '';
	let loading = false;
	let messages: Message[] = [];
	let chatName = 'This is the name of this chat';

	async function sendPrompt(prompt: string) {
		messages = [...messages, { role: 'user', text: prompt }];

		loading = true;
		try {
			const response = await fetch(`${apiBaseUrl}/chat`, {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify({
					message: prompt
				})
			});

			if (!response.ok) {
				throw new Error(`HTTP error. Status: ${response.status}`);
			}

			const data = await response.json();

			messages = [...messages, { role: 'assistant', text: data.response }];
		} catch (err) {
			messages = [...messages, { role: 'assistant', text: (err as Error).message }];
		} finally {
			loading = false;
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

<div class="flex h-screen flex-col bg-background text-text lg:ml-72">
	<!-- Chat Header -->
	<div class="border-b border-white/10 bg-card px-6 py-4 pl-16 lg:pl-6">
		<h2 class="truncate text-lg font-medium text-text">{chatName}</h2>
	</div>

	<!-- Chat Messages -->
	<div class="flex-1 overflow-y-auto px-6 py-8 pb-32">
		<div class="mx-auto max-w-3xl">
			{#each messages as msg, i (i)}
				<div class="flex {msg.role === 'user' ? 'justify-end' : 'justify-start'} mb-4">
					<div
						class={`max-w-[85%] px-4 py-2.5 text-sm leading-relaxed ${
							msg.role === 'user' ? 'rounded-2xl bg-card text-text' : 'text-text'
						}`}
					>
						{msg.text}
					</div>
				</div>
			{/each}
			{#if loading}
				<div class="mb-4 flex justify-start">
					<div class="px-1 py-2 text-sm text-text-muted">
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

	<!-- Floating Input -->
	<div
		class="fixed right-0 bottom-0 left-0 bg-linear-to-t from-background via-background to-transparent px-4 pt-8 pb-6 lg:left-72"
	>
		<div
			class="mx-auto flex max-w-3xl items-end space-x-2 rounded-2xl border border-white/10 bg-card p-2 shadow-lg"
		>
			<textarea
				bind:value={prompt}
				rows="1"
				placeholder="Type a message..."
				class="flex-1 resize-none border-0 bg-transparent px-3 py-2.5 text-sm text-text placeholder-text-muted focus:ring-0 focus:outline-none"
				style="overflow: hidden;"
				on:input={(e) => {
					e.currentTarget.style.height = 'auto';
					e.currentTarget.style.height = e.currentTarget.scrollHeight + 'px';
				}}
				on:keydown={(e) => e.key === 'Enter' && !e.shiftKey && (e.preventDefault(), handleSend())}
			></textarea>
			<button
				on:click={handleSend}
				class="shrink-0 rounded-lg bg-text p-2.5 text-background transition hover:bg-text/90 disabled:cursor-not-allowed disabled:opacity-50"
				disabled={loading || !prompt.trim()}
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
				>
					<path d="M12 19V5M5 12l7-7 7 7" />
				</svg>
			</button>
		</div>
	</div>
</div>
