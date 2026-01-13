<script lang="ts">
	import type { Chat } from '$lib/types/chat';

    export let startNewChat : () => void
	export let personaName;
	export let chatList: Chat[]

    function handleNewChat(){
        startNewChat()
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
			<!-- New Chat Button -->
			<button
				class="mb-8 w-full rounded-2xl border border-white/10 bg-card px-4 py-3 text-center text-lg text-text transition hover:bg-card/80"
                on:click={handleNewChat}
                on:keydown={(e)=> e.key === "Enter" && handleNewChat()}
			>
				Start a new chat
			</button>

			{#if chatList.length > 0}
				<!-- Chats Label -->
				<h3 class="mb-4 text-sm font-medium text-text-muted">Chats</h3>
				<!-- Chat List -->
				<div>
					{#each chatList as chat}
						<div
							class="cursor-pointer border-b border-white/10 py-4 transition hover:opacity-80"
						>
							<div class="text-lg font-semibold text-text">{chat.title}</div>
							<div class="truncate text-sm text-text-muted">{chat.lastMessage}</div>
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
