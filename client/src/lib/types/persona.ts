import type { Chat } from './chat';

export interface Persona {
	id: string;
	name: string;
	chats: Chat[];
}
