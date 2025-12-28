<script lang="ts">
	import Sidebar from '$lib/components/sidebar.svelte';
	import Chat from '$lib/components/chat.svelte';
	import type { Persona } from '$lib/types/persona.js';
	import type { User } from '$lib/types/user';
	import type { Message } from '$lib/types/message';

	export let data: User;

	let chatName = 'This is the chat name';
	let messages: Message[] = [];
	let loading: boolean = false;

	async function sendPrompt(prompt: string) {
		messages = [...messages, { role: 'user', text: prompt }];

		loading = true;
		try {
			const response = await fetch(`api/chat`, {
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

	// sample data, will be removed later
	let personas: Persona[] = [
		{
			id: '1',
			name: 'Helpful Assistant',
			chats: [
				{ id: 'c1', title: 'Getting started with AI' },
				{ id: 'c2', title: 'Code review help' },
				{ id: 'c3', title: 'Project planning discussion' }
			]
		},
		{
			id: '2',
			name: 'Creative Writer',
			chats: [
				{ id: 'c4', title: 'Story ideas brainstorm' },
				{ id: 'c5', title: 'Character development' }
			]
		},
		{
			id: '3',
			name: 'Technical Expert',
			chats: [
				{ id: 'c6', title: 'Database optimization' },
				{ id: 'c7', title: 'System architecture review' },
				{ id: 'c8', title: 'Security best practices' },
				{ id: 'c9', title: 'API design patterns' }
			]
		}
	];
</script>

<Sidebar name={data.name} email={data.email} {personas} />
<Chat {chatName} {messages} {sendPrompt} {loading} />
