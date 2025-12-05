package models

type Persona struct {
	PersonaID string `json:"persona_id"`
	UserID    string `json:"user_id"`
	About
	BaseType
	More
	Interests
	Language
	Others
}

type About struct {
	PersonaName        string   `json:"persona_name"`
	PersonaDescription string   `json:"persona_description"`
	Age                int      `json:"age"`
	Pronouns           []string `json:"pronouns"`
}

type BaseType struct {
	Openness          float32 `json:"openness"`
	Conscientiousness float32 `json:"conscientiousness"`
	Extraversion      float32 `json:"extraversion"`
	Agreeableness     float32 `json:"agreeableness"`
	Neuroticism       float32 `json:"neuroticism"`
}

type More struct {
	Intelligence    float32 `json:"intelligence"`
	ThinkingStyle   float32 `json:"thinking_style"`
	Tone            string  `json:"tone"`
	HumorLevel      float32 `json:"humor_level"`
	MoodFluctuation float32 `json:"mood_fluctuation"`
}

type Interests struct {
	Likes    []string `json:"likes"`
	Dislikes []string `json:"dislikes"`
}

type Language struct {
	Formality float32 `json:"language_formality"`
	Fluency   string  `json:"language_fluency"`
}

type Others struct {
	EmojiUsage     string `json:"emoji_usage"`
	ResponseLength string `json:"response_length"`
}
