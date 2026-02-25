<script lang="ts">
	/**
	 * Chat list and new chat starter component.
	 * Displays recent chat sessions with relative timestamps and handles new chat creation.
	 * Shows desktop and mobile layouts with different input positioning.
	 *
	 * @component
	 * @prop {function} startNewChat - Creates a new chat session
	 * @prop {function} openChatSession - Opens existing chat
	 * @prop {function} deleteChatSession - Deletes a chat
	 * @prop {string} personaName - Persona name displayed in header
	 * @prop {boolean} isLoading - Shows spinner in start button when true
	 */
	import { chats } from '../../stores/store';
	import Trash from '$lib/assets/icons/trash.png';

	export let startNewChat: (message: string) => void;
	export let openChatSession: (sessionid: string) => void;
	export let deleteChatSession: (sessionid: string) => void;
	export let personaName;
	export let isLoading: boolean;

	let message = '';

	function handleNewChat() {
		startNewChat(message);
		message = '';
	}

	function formatTimestamp(isoString: string): string {
		const date = new Date(isoString);
		const now = new Date();
		const diffMs = now.getTime() - date.getTime();
		const diffMins = Math.floor(diffMs / 60000);
		const diffHours = Math.floor(diffMs / 3600000);
		const diffDays = Math.floor(diffMs / 86400000);

		if (diffMins < 1) return 'Just now';
		if (diffMins < 60) return `${diffMins} min ago`;
		if (diffHours < 24) return `${diffHours} hour${diffHours > 1 ? 's' : ''} ago`;
		if (diffDays === 1) return 'Yesterday';
		if (diffDays < 7) return `${diffDays} days ago`;

		// For older dates, show the actual date
		return date.toLocaleDateString('en-US', { month: 'short', day: 'numeric', year: 'numeric' });
	}
</script>

<div class="flex h-dvh flex-col bg-background text-text lg:ml-72">
	<!-- Header -->
	<div class="sticky top-0 z-10 border-b border-white/10 bg-card px-4 py-4 pl-16 sm:px-6 lg:pl-6">
		<h2 class="truncate text-lg font-medium text-text">{personaName}</h2>
	</div>

	<!-- Desktop input -->
	<div class="hidden px-6 py-6 lg:block">
		<div class="mx-auto max-w-3xl">
			<div class="flex w-full gap-3">
				<input
					bind:value={message}
					type="text"
					placeholder="Type your message to start a new chat..."
					class="flex-1 rounded-lg border border-white/10 bg-card px-4 py-4 text-text transition focus:border-white/20 focus:outline-none"
					on:keydown={(e) => e.key === 'Enter' && handleNewChat()}
				/>
				<button
					class="rounded-lg bg-text px-8 py-4 font-medium text-background transition hover:opacity-90"
					on:click={handleNewChat}
				>
					{#if isLoading}
						<div
							class="h-4 w-4 animate-spin rounded-full border-2 border-current border-r-transparent"
						></div>
					{:else}
						Start Chat
					{/if}
				</button>
			</div>
		</div>
	</div>

	<!-- Content -->
	<div class="flex-1 overflow-y-auto px-4 py-6 pb-32 lg:px-6 lg:py-8 lg:pb-8">
		<div class="mx-auto max-w-3xl">
			{#if $chats.length > 0}
				<!-- Chats Section -->
				<div>
					<h3 class="mb-4 text-sm font-medium text-text-muted">Recent Chats</h3>

					<div class="space-y-3">
						{#each $chats as chat}
							<div
								class="group relative cursor-pointer rounded-xl border border-white/10 bg-card/30 p-5 transition hover:bg-card/50"
								on:click={() => openChatSession(chat.session_id)}
								on:keydown={(e) => e.key === 'Enter' && openChatSession(chat.session_id)}
								role="button"
								tabindex="0"
							>
								<div class="flex items-center justify-between gap-4">
									<div class="min-w-0 flex-1">
										<h4 class="mb-1 truncate text-base font-medium text-text">
											{chat.title}
										</h4>
										<p class="text-xs text-text-muted">
											{formatTimestamp(chat.updated_at)}
										</p>
									</div>

									<button
										class="flex-shrink-0 rounded-lg p-2 text-text-muted transition hover:bg-background hover:text-text"
										aria-label="Delete chat"
										on:click={(e) => {
											e.stopPropagation();
											deleteChatSession(chat.session_id);
										}}
									>
										<img src={Trash} alt="" class="h-4 w-4" />
									</button>
								</div>
							</div>
						{/each}
					</div>
				</div>
			{:else}
				<!-- Empty State -->
				<div class="mt-16 text-center">
					<div
						class="mx-auto mb-4 flex h-16 w-16 items-center justify-center rounded-full bg-card/50"
					>
						<svg
							class="h-8 w-8 text-text-muted"
							fill="none"
							stroke="currentColor"
							viewBox="0 0 24 24"
						>
							<path
								stroke-linecap="round"
								stroke-linejoin="round"
								stroke-width="2"
								d="M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z"
							/>
						</svg>
					</div>
					<p class="mb-2 text-lg font-medium text-text">No conversations yet</p>
					<p class="text-sm text-text-muted">
						Start chatting with {personaName} to begin your first conversation
					</p>
				</div>
			{/if}
		</div>
	</div>

	<!-- Mobile Floating Input (hidden on desktop) -->
	<div class="fixed right-0 bottom-0 left-0 bg-background shadow-2xl lg:hidden">
		<div class="mx-auto max-w-3xl px-4 py-4">
			<div class="flex w-full gap-2">
				<input
					bind:value={message}
					type="text"
					placeholder="Start a new chat..."
					class="flex-1 rounded-lg border border-white/10 bg-card px-3 py-3.5 text-sm text-text transition focus:border-white/20 focus:outline-none"
					on:keydown={(e) => e.key === 'Enter' && handleNewChat()}
				/>
				<button
					class="rounded-lg bg-text px-5 py-3.5 text-sm font-medium text-background transition hover:opacity-90"
					on:click={handleNewChat}
				>
					{#if isLoading}
						<div
							class="h-4 w-4 animate-spin rounded-full border-2 border-current border-r-transparent"
						></div>
					{:else}
						Start
					{/if}
				</button>
			</div>
		</div>
	</div>
</div>
