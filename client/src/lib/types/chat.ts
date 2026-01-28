export interface Chat {
	session_id: string;
	title: string;
	lastMessage: string
}

export interface NewChatResponse {
	session_id?: string
	title?: string
	response: string
}
