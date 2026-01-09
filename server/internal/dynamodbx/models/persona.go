package models

type Persona struct {
	PersonaID          string   `json:"persona_id"` //SK
	UserID             string   `json:"user_id"`    //PK
	PersonaName        string   `json:"name"`
	PersonaDescription string   `json:"description"`
	Age                int      `json:"age"`
	Pronouns           []string `json:"pronouns"`
	Openness           float32  `json:"openness"`
	Conscientiousness  float32  `json:"conscientiousness"`
	Extraversion       float32  `json:"extraversion"`
	Agreeableness      float32  `json:"agreeableness"`
	Neuroticism        float32  `json:"neuroticism"`
	Intelligence       float32  `json:"intelligence"`
	ThinkingStyle      float32  `json:"thinking_style"`
	Tone               []string `json:"tone"`
	HumorLevel         float32  `json:"humor_level"`
	MoodFluctuation    float32  `json:"mood_fluctuation"`
	Likes              []string `json:"likes"`
	Dislikes           []string `json:"dislikes"`
	Formality          float32  `json:"formality"`
	Fluency            string   `json:"fluency"`
	EmojiUsage         string   `json:"emoji_usage"`
	ResponseLength     string   `json:"response_length"`
	CreatedAt          string   `json:"created_at"`
}

type PersonaList struct {
	PersonaID   string `json:"persona_id"`
	PersonaName string `json:"name"`
	CreatedAt   string `json:"created_at"`
}
