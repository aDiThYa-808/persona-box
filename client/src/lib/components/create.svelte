<script lang="ts">
	import type { PersonaData } from '$lib/types/persona';

	export let createPersona: (data: PersonaData) => void;

	let formData = {
		// About
		name: '',
		description: '',
		age: '',
		pronouns: '',

		// Base Type (Big Five)
		openness: 5,
		conscientiousness: 5,
		extraversion: 5,
		agreeableness: 5,
		neuroticism: 5,

		// More
		intelligence: 5,
		thinking_style: 5,
		tone: '',
		humor_level: 5,
		mood_fluctuation: 5,

		// Interests
		likes: '',
		dislikes: '',

		// Language
		formality: 5,
		fluency: '',

		// Others
		emoji_usage: '',
		response_length: ''
	};

	$: isFormValid =
		formData.name.trim() !== '' &&
		formData.description.trim() !== '' &&
		formData.age !== '' &&
		formData.pronouns.trim() !== '' &&
		formData.tone.trim() !== '' &&
		formData.likes.trim() !== '' &&
		formData.dislikes.trim() !== '' &&
		formData.fluency.trim() !== '' &&
		formData.emoji_usage.trim() !== '' &&
		formData.response_length.trim() !== '';

	function handleSubmit() {
		if (!isFormValid) return;

		const personaData = {
			name: formData.name,
			description: formData.description,
			age: parseInt(formData.age),
			pronouns: formData.pronouns,
			openness: formData.openness,
			conscientiousness: formData.conscientiousness,
			extraversion: formData.extraversion,
			agreeableness: formData.agreeableness,
			neuroticism: formData.neuroticism,
			intelligence: formData.intelligence,
			thinking_style: formData.thinking_style,
			tone: formData.tone,
			humor_level: formData.humor_level,
			mood_fluctuation: formData.mood_fluctuation,
			likes: formData.likes
				.split(',')
				.map((item) => item.trim())
				.filter(Boolean),
			dislikes: formData.dislikes
				.split(',')
				.map((item) => item.trim())
				.filter(Boolean),
			formality: formData.formality,
			fluency: formData.fluency,
			emoji_usage: formData.emoji_usage,
			response_length: formData.response_length
		};

		createPersona(personaData);
	}
</script>

<div class="flex h-screen flex-col bg-background text-text lg:ml-72">
	<!-- Header -->
	<div class="border-b border-white/10 bg-card px-6 py-4 pl-16 lg:pl-6">
		<h2 class="text-lg font-medium text-text">Create New Persona</h2>
	</div>

	<!-- Content -->
	<div class="flex-1 overflow-y-auto px-6 py-8">
		<div class="mx-auto max-w-3xl">
			<!-- About Section -->
			<div class="mb-12">
				<h3 class="mb-2 text-xl font-semibold text-text">About</h3>
				<p class="mb-6 text-sm text-text-muted">
					Basic information about your persona's identity and background.
				</p>

				<div class="space-y-4">
					<div>
						<label for="name" class="mb-2 block text-sm font-medium text-text">Name *</label>
						<input
							id="name"
							type="text"
							bind:value={formData.name}
							class="w-full rounded-lg border border-white/10 bg-card px-4 py-2.5 text-text placeholder-text-muted focus:ring-2 focus:ring-white/20 focus:outline-none"
							placeholder="Enter persona name"
						/>
					</div>

					<div>
						<label for="description" class="mb-2 block text-sm font-medium text-text"
							>Description *</label
						>
						<textarea
							id="description"
							rows="3"
							bind:value={formData.description}
							class="w-full resize-none rounded-lg border border-white/10 bg-card px-4 py-2.5 text-text placeholder-text-muted focus:ring-2 focus:ring-white/20 focus:outline-none"
							placeholder="Describe your persona"
						></textarea>
					</div>

					<div class="grid grid-cols-2 gap-4">
						<div>
							<label for="age" class="mb-2 block text-sm font-medium text-text">Age *</label>
							<input
								id="age"
								type="number"
								min="1"
								bind:value={formData.age}
								class="w-full rounded-lg border border-white/10 bg-card px-4 py-2.5 text-text placeholder-text-muted focus:ring-2 focus:ring-white/20 focus:outline-none"
								placeholder="25"
							/>
						</div>

						<div>
							<label for="pronouns" class="mb-2 block text-sm font-medium text-text"
								>Pronouns *</label
							>
							<input
								id="pronouns"
								type="text"
								bind:value={formData.pronouns}
								class="w-full rounded-lg border border-white/10 bg-card px-4 py-2.5 text-text placeholder-text-muted focus:ring-2 focus:ring-white/20 focus:outline-none"
								placeholder="he/him"
							/>
						</div>
					</div>
				</div>
			</div>

			<!-- Base Type Section -->
			<div class="mb-12">
				<h3 class="mb-2 text-xl font-semibold text-text">Personality (Big Five)</h3>
				<p class="mb-6 text-sm text-text-muted">
					Core personality traits that define how your persona thinks and behaves.
				</p>

				<div class="mb-6">
					<label for="openness" class="mb-2 block text-sm font-medium text-text">
						Openness
						<span class="ml-2 text-text-muted">({formData.openness})</span>
					</label>
					<input
						id="openness"
						type="range"
						min="0"
						max="10"
						step="0.1"
						bind:value={formData.openness}
						class="h-2 w-full cursor-pointer appearance-none rounded-lg bg-card accent-text"
					/>
					<div class="mt-1 flex justify-between text-xs text-text-muted">
						<span>0</span>
						<span>10</span>
					</div>
				</div>

				<div class="mb-6">
					<label for="conscientiousness" class="mb-2 block text-sm font-medium text-text">
						Conscientiousness
						<span class="ml-2 text-text-muted">({formData.conscientiousness})</span>
					</label>
					<input
						id="conscientiousness"
						type="range"
						min="0"
						max="10"
						step="0.1"
						bind:value={formData.conscientiousness}
						class="h-2 w-full cursor-pointer appearance-none rounded-lg bg-card accent-text"
					/>
					<div class="mt-1 flex justify-between text-xs text-text-muted">
						<span>0</span>
						<span>10</span>
					</div>
				</div>

				<div class="mb-6">
					<label for="extraversion" class="mb-2 block text-sm font-medium text-text">
						Extraversion
						<span class="ml-2 text-text-muted">({formData.extraversion})</span>
					</label>
					<input
						id="extraversion"
						type="range"
						min="0"
						max="10"
						step="0.1"
						bind:value={formData.extraversion}
						class="h-2 w-full cursor-pointer appearance-none rounded-lg bg-card accent-text"
					/>
					<div class="mt-1 flex justify-between text-xs text-text-muted">
						<span>0</span>
						<span>10</span>
					</div>
				</div>

				<div class="mb-6">
					<label for="agreeableness" class="mb-2 block text-sm font-medium text-text">
						Agreeableness
						<span class="ml-2 text-text-muted">({formData.agreeableness})</span>
					</label>
					<input
						id="agreeableness"
						type="range"
						min="0"
						max="10"
						step="0.1"
						bind:value={formData.agreeableness}
						class="h-2 w-full cursor-pointer appearance-none rounded-lg bg-card accent-text"
					/>
					<div class="mt-1 flex justify-between text-xs text-text-muted">
						<span>0</span>
						<span>10</span>
					</div>
				</div>

				<div class="mb-6">
					<label for="neuroticism" class="mb-2 block text-sm font-medium text-text">
						Neuroticism
						<span class="ml-2 text-text-muted">({formData.neuroticism})</span>
					</label>
					<input
						id="neuroticism"
						type="range"
						min="0"
						max="10"
						step="0.1"
						bind:value={formData.neuroticism}
						class="h-2 w-full cursor-pointer appearance-none rounded-lg bg-card accent-text"
					/>
					<div class="mt-1 flex justify-between text-xs text-text-muted">
						<span>0</span>
						<span>10</span>
					</div>
				</div>
			</div>

			<!-- More Section -->
			<div class="mb-12">
				<h3 class="mb-2 text-xl font-semibold text-text">Additional Traits</h3>
				<p class="mb-6 text-sm text-text-muted">
					Fine-tune your persona's cognitive and emotional characteristics.
				</p>

				<div class="mb-6">
					<label for="intelligence" class="mb-2 block text-sm font-medium text-text">
						Intelligence
						<span class="ml-2 text-text-muted">({formData.intelligence})</span>
					</label>
					<input
						id="intelligence"
						type="range"
						min="0"
						max="10"
						step="0.1"
						bind:value={formData.intelligence}
						class="h-2 w-full cursor-pointer appearance-none rounded-lg bg-card accent-text"
					/>
					<div class="mt-1 flex justify-between text-xs text-text-muted">
						<span>0</span>
						<span>10</span>
					</div>
				</div>

				<div class="mb-6">
					<label for="thinking_style" class="mb-2 block text-sm font-medium text-text">
						Thinking Style
						<span class="ml-2 text-text-muted">({formData.thinking_style})</span>
					</label>
					<input
						id="thinking_style"
						type="range"
						min="0"
						max="10"
						step="0.1"
						bind:value={formData.thinking_style}
						class="h-2 w-full cursor-pointer appearance-none rounded-lg bg-card accent-text"
					/>
					<div class="mt-1 flex justify-between text-xs text-text-muted">
						<span>0</span>
						<span>10</span>
					</div>
				</div>

				<div class="mb-6">
					<label for="tone" class="mb-2 block text-sm font-medium text-text">Tone *</label>
					<input
						id="tone"
						type="text"
						bind:value={formData.tone}
						class="w-full rounded-lg border border-white/10 bg-card px-4 py-2.5 text-text placeholder-text-muted focus:ring-2 focus:ring-white/20 focus:outline-none"
						placeholder="e.g., friendly, professional, casual"
					/>
				</div>

				<div class="mb-6">
					<label for="humor_level" class="mb-2 block text-sm font-medium text-text">
						Humor Level
						<span class="ml-2 text-text-muted">({formData.humor_level})</span>
					</label>
					<input
						id="humor_level"
						type="range"
						min="0"
						max="10"
						step="0.1"
						bind:value={formData.humor_level}
						class="h-2 w-full cursor-pointer appearance-none rounded-lg bg-card accent-text"
					/>
					<div class="mt-1 flex justify-between text-xs text-text-muted">
						<span>0</span>
						<span>10</span>
					</div>
				</div>

				<div class="mb-6">
					<label for="mood_fluctuation" class="mb-2 block text-sm font-medium text-text">
						Mood Fluctuation
						<span class="ml-2 text-text-muted">({formData.mood_fluctuation})</span>
					</label>
					<input
						id="mood_fluctuation"
						type="range"
						min="0"
						max="10"
						step="0.1"
						bind:value={formData.mood_fluctuation}
						class="h-2 w-full cursor-pointer appearance-none rounded-lg bg-card accent-text"
					/>
					<div class="mt-1 flex justify-between text-xs text-text-muted">
						<span>0</span>
						<span>10</span>
					</div>
				</div>
			</div>

			<!-- Interests Section -->
			<div class="mb-12">
				<h3 class="mb-2 text-xl font-semibold text-text">Interests</h3>
				<p class="mb-6 text-sm text-text-muted">
					Define what your persona enjoys and dislikes. Separate items with commas.
				</p>

				<div class="space-y-4">
					<div>
						<label for="likes" class="mb-2 block text-sm font-medium text-text">Likes *</label>
						<input
							id="likes"
							type="text"
							bind:value={formData.likes}
							class="w-full rounded-lg border border-white/10 bg-card px-4 py-2.5 text-text placeholder-text-muted focus:ring-2 focus:ring-white/20 focus:outline-none"
							placeholder="reading, music, technology"
						/>
					</div>

					<div>
						<label for="dislikes" class="mb-2 block text-sm font-medium text-text">Dislikes *</label
						>
						<input
							id="dislikes"
							type="text"
							bind:value={formData.dislikes}
							class="w-full rounded-lg border border-white/10 bg-card px-4 py-2.5 text-text placeholder-text-muted focus:ring-2 focus:ring-white/20 focus:outline-none"
							placeholder="negativity, rudeness, spam"
						/>
					</div>
				</div>
			</div>

			<!-- Language Section -->
			<div class="mb-12">
				<h3 class="mb-2 text-xl font-semibold text-text">Language Style</h3>
				<p class="mb-6 text-sm text-text-muted">
					Configure how your persona communicates and expresses themselves.
				</p>

				<div class="mb-6">
					<label for="formality" class="mb-2 block text-sm font-medium text-text">
						Formality
						<span class="ml-2 text-text-muted">({formData.formality})</span>
					</label>
					<input
						id="formality"
						type="range"
						min="0"
						max="10"
						step="0.1"
						bind:value={formData.formality}
						class="h-2 w-full cursor-pointer appearance-none rounded-lg bg-card accent-text"
					/>
					<div class="mt-1 flex justify-between text-xs text-text-muted">
						<span>0</span>
						<span>10</span>
					</div>
				</div>

				<div class="mb-6">
					<label for="fluency" class="mb-2 block text-sm font-medium text-text">Fluency *</label>
					<input
						id="fluency"
						type="text"
						bind:value={formData.fluency}
						class="w-full rounded-lg border border-white/10 bg-card px-4 py-2.5 text-text placeholder-text-muted focus:ring-2 focus:ring-white/20 focus:outline-none"
						placeholder="e.g., native, fluent, conversational"
					/>
				</div>
			</div>

			<!-- Others Section -->
			<div class="mb-12">
				<h3 class="mb-2 text-xl font-semibold text-text">Other Settings</h3>
				<p class="mb-6 text-sm text-text-muted">
					Additional preferences for your persona's communication style.
				</p>

				<div class="space-y-4">
					<div>
						<label for="emoji_usage" class="mb-2 block text-sm font-medium text-text"
							>Emoji Usage *</label
						>
						<input
							id="emoji_usage"
							type="text"
							bind:value={formData.emoji_usage}
							class="w-full rounded-lg border border-white/10 bg-card px-4 py-2.5 text-text placeholder-text-muted focus:ring-2 focus:ring-white/20 focus:outline-none"
							placeholder="e.g., frequent, moderate, rare, none"
						/>
					</div>

					<div>
						<label for="response_length" class="mb-2 block text-sm font-medium text-text"
							>Response Length *</label
						>
						<input
							id="response_length"
							type="text"
							bind:value={formData.response_length}
							class="w-full rounded-lg border border-white/10 bg-card px-4 py-2.5 text-text placeholder-text-muted focus:ring-2 focus:ring-white/20 focus:outline-none"
							placeholder="e.g., concise, moderate, detailed"
						/>
					</div>
				</div>
			</div>

			<!-- Submit Button -->
			<div class="flex justify-end border-t border-white/10 pt-6">
				<button
					on:click={handleSubmit}
					disabled={!isFormValid}
					class="rounded-lg bg-text px-6 py-3 font-medium text-background transition hover:bg-text/90 disabled:cursor-not-allowed disabled:opacity-50"
				>
					Create Persona
				</button>
			</div>
		</div>
	</div>
</div>
