<script lang="ts">
	import { chats } from '../../stores/store';

	export let startNewChat: (message: string) => void;
	export let personaName;

	let message = ""

	function handleNewChat() {
		startNewChat(message);
		message = ""
	}
</script>

<div class="flex h-screen flex-col bg-background text-text lg:ml-72">
	<!-- Header -->
	<div class="border-b border-white/10 bg-card px-6 py-4 pl-16 lg:pl-6">
		<h2 class="truncate text-lg font-medium text-text">{personaName}</h2>
	</div>

	<!-- Content -->
	<div class="flex-1 overflow-y-auto px-6 py-8">
		<div class="mx-auto max-w-3xl">
			<!-- Message Input Container -->
			<div class="mb-8 flex w-full gap-2">
				<input
					bind:value={message}
					type="text"
					placeholder="Type your message..."
					class="flex-1 rounded-2xl border border-white/10 bg-card px-4 py-3 text-lg text-text transition focus:border-white/20 focus:outline-none"
					on:keydown={(e) => e.key === 'Enter' && handleNewChat()}
				/>
				<button
					class="rounded-2xl border border-white/10 bg-card px-6 py-3 text-lg text-text transition hover:bg-card/80"
					on:click={handleNewChat}
				>
					Send
				</button>
			</div>

			{#if $chats.length > 0}
				<!-- Chats Label -->
				<h3 class="mb-4 text-sm font-medium text-text-muted">Chats</h3>
				<!-- Chat List -->
				<div>
					{#each $chats as chat}
						<div class="cursor-pointer border-b border-white/10 py-4 transition hover:opacity-80">
							<div class="text-lg font-semibold text-text">{chat.title}</div>
						</div>
					{/each}
				</div>
			{:else}
				<!-- Empty State -->
				<div class="mt-12 text-center">
					<p class="text-lg text-text-muted">No chats yet. Start a new chat to begin.</p>
				</div>
			{/if}
		</div>
	</div>
</div>
