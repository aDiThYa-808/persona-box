export interface PersonaList {
	persona_id: string;
	name: string;
	created_at: string;
}

export interface PersonaData {
	// About
	name: string;
	description: string;
	age: number;
	pronouns: string[];

	// Base Type
	openness: number;
	conscientiousness: number;
	extraversion: number;
	agreeableness: number;
	neuroticism: number;

	// More
	intelligence: number;
	thinking_style: number;
	tone: string[];
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

export interface TraitInfo {
	description: string
	lowLabel : string
	highLabel : string
}