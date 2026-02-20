<script lang="ts">
  export let isOpen: boolean = false;
  export let title: string;
  export let message: string;
  export let confirmButtonText: string
  export let cancelButtonText: string
  export let onConfirm = () => {};
  export let onCancel = () => {};
</script>

{#if isOpen}
	<!-- Overlay -->
	<div role="button" tabindex="0" class="fixed inset-0 z-50 bg-black/25 backdrop-blur-sm" on:click={onCancel} on:keypress={(e)=> e.key === 'Enter' && onCancel}></div>
	
	<!-- Modal -->
	<div class="fixed inset-0 z-50 flex items-center justify-center p-4 pointer-events-none">
		<div class="pointer-events-auto w-full max-w-md rounded-xl border border-white/10 bg-card p-6 shadow-2xl">
			<h3 class="mb-3 text-xl font-semibold text-text">{title}</h3>
			<p class="mb-6 text-text-muted">{message}</p>
			
			<div class="flex gap-3 justify-end">
				<button
					on:click={onCancel}
					class="rounded-lg border border-white/10 bg-background px-5 py-2.5 font-medium text-text transition hover:bg-card"
				>
					{cancelButtonText}
				</button>
				<button
					on:click={onConfirm}
					class="rounded-lg bg-button-danger px-5 py-2.5 font-medium text-background transition hover:opacity-90"
				>
					{confirmButtonText}
				</button>
			</div>
		</div>
	</div>
{/if}