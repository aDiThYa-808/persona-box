export interface Chat {
	session_id: string;
	title: string;
	updated_at: string
}

export interface NewChatResponse {
	session_id?: string
	title?: string
	response: string
}
