import { PUBLIC_AWS_INVOKE_URL } from '$env/static/public';
import type { PageServerLoad } from './$types';
import type { Message } from '$lib/types/message';

export const load: PageServerLoad = async ({ params, fetch }) => {
	try {
		const invokeUrl = PUBLIC_AWS_INVOKE_URL;
		const chatId = params.chatid;

		if (chatId == 'new-chat') {
			return {
				messages: []
			};
		}

		const res = await fetch(`/api/messages?sessionid=${chatId}`);

		if (!res.ok) {
			let data = res.json();
			console.log(`Failed to fetch messages. ${JSON.stringify(data)}`);
			return {
				messages: []
			};
		}

		const sessionMessages: Message[] = await res.json();
		console.log(sessionMessages)
		return {
			messages: sessionMessages.sort((a, b) => b.created_at.localeCompare(a.created_at))
		};
	} catch (err) {
		console.log(err);
		return {
			messages: []
		};
	}
};
