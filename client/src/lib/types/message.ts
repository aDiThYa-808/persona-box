export interface Message {
	created_at: string
	role: 'user' | 'assistant'
	message: string
}

