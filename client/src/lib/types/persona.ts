import type { Chat } from './chat';

export interface Persona {
	id: string;
	name: string;
	chats: Chat[];
}

export interface PersonaData {
	// About
	name: string;
	description: string;
	age: number;
	pronouns: string;

	// Base Type (Big Five)
	openness: number;
	conscientiousness: number;
	extraversion: number;
	agreeableness: number;
	neuroticism: number;

	// More
	intelligence: number;
	thinking_style: number;
	tone: string;
	humor_level: number;
	mood_fluctuation: number;

	// Interests
	likes: string[];
	dislikes: string[];

	// Language
	formality: number;
	fluency: string;

	// Others
	emoji_usage: string;
	response_length: string;
}
