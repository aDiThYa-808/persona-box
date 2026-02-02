import type { Load } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';
import type { Chat } from '$lib/types/chat';

export const load: PageServerLoad = async ({ params, fetch }) => {
	const personaid = params.personaid;
	if (personaid == 'new-persona') {
		return {
			chatSessions: []
		};
	}

	try {
		const res = await fetch(`/api/chat-session?personaid=${personaid}`);

		if (!res.ok) {
			const data = await res.json();
			console.log(`HTTP Error. Status: ${res.status}. message: ${JSON.stringify(data)}`);
			return {
				chatSessions: []
			};
		}

		const chatSessions: Chat[] = await res.json();

		return {
			chatSessions: chatSessions.sort((a, b) => b.updated_at.localeCompare(a.updated_at))
		};
	} catch (err) {
		console.log(err);
		return {
			chatSessions: []
		};
	}
};
