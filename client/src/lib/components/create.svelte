<script lang="ts">
	import type { PersonaData, TraitInfo } from '$lib/types/persona';

	export let createPersona: (data: PersonaData) => void;
	export let isLoading: boolean;

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
			pronouns: formData.pronouns
				.split('/')
				.map((item) => item.trim())
				.filter(Boolean),
			openness: formData.openness,
			conscientiousness: formData.conscientiousness,
			extraversion: formData.extraversion,
			agreeableness: formData.agreeableness,
			neuroticism: formData.neuroticism,
			intelligence: formData.intelligence,
			thinking_style: formData.thinking_style,
			tone: formData.tone
				.split(',')
				.map((item) => item.trim())
				.filter(Boolean),
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

	const oceanTraits: Record<string, TraitInfo> = {
		openness: {
			description: 'Measures creativity, curiosity, and willingness to entertain new ideas.',
			lowLabel: 'Practical',
			highLabel: 'Imaginative'
		},
		conscientiousness: {
			description: 'Measures self-control, diligence, and attention to detail.',
			lowLabel: 'Spontaneous',
			highLabel: 'Disciplined'
		},
		extraversion: {
			description: 'Measures boldness, energy, and social interactivity.',
			lowLabel: 'Reserved',
			highLabel: 'Outgoing'
		},
		agreeableness: {
			description: 'Measures kindness, helpfulness, and willingness to cooperate.',
			lowLabel: 'Challenging',
			highLabel: 'Compassionate'
		},
		neuroticism: {
			description: 'Measures depression, irritability, and proneness to anxiety.',
			lowLabel: 'Resilient',
			highLabel: 'Sensitive'
		}
	};

	const additionalTraits: Record<string, TraitInfo> = {
		intelligence: {
			description: 'How knowledgeable and analytical the persona appears in conversations.',
			lowLabel: 'Simple',
			highLabel: 'Sophisticated'
		},
		thinking_style: {
			description: 'How the persona processes information and makes decisions.',
			lowLabel: 'Intuitive',
			highLabel: 'Analytical'
		},
		humor_level: {
			description: 'How often the persona uses jokes, wit, and playful language.',
			lowLabel: 'Serious',
			highLabel: 'Playful'
		},
		mood_fluctuation: {
			description: "How much the persona's emotional state varies during conversations.",
			lowLabel: 'Stable',
			highLabel: 'Variable'
		}
	};

	function handleSliderFocus(event: Event) {
		// Force blur any active text input when slider is touched
		const activeElement = document.activeElement;
		if (
			activeElement &&
			(activeElement instanceof HTMLInputElement ||
				activeElement instanceof HTMLTextAreaElement ||
				activeElement instanceof HTMLSelectElement)
		) {
			activeElement.blur();
		}
	}
</script>

<div class="flex h-screen flex-col bg-background text-text lg:ml-72">
	<!-- Header -->
	<div class="sticky top-0 z-10 border-b border-white/10 bg-card px-4 py-4 pl-16 sm:px-6 lg:pl-6">
		<h2 class="truncate text-lg font-medium text-text">Create New Persona</h2>
	</div>

	<!-- Content -->
	<div class="flex-1 overflow-y-auto px-6 py-8">
		<div class="mx-auto max-w-3xl">
			<!-- About Section -->
			<div class="mb-12">
				<h3 class="mb-2 text-xl font-semibold text-text">About</h3>
				<p class="mb-6 text-sm text-text-muted">
					Define your persona's basic identity. This information helps establish their character and
					how they'll introduce themselves.
				</p>

				<div class="space-y-6">
					<div>
						<label for="name" class="mb-2 block text-sm font-medium text-text">Name *</label>
						<input
							id="name"
							type="text"
							bind:value={formData.name}
							class="w-full rounded-lg border border-white/10 bg-card px-4 py-2.5 text-text placeholder-text-muted focus:ring-2 focus:ring-white/20 focus:outline-none"
							placeholder="e.g., Alex Chen"
						/>
					</div>

					<div>
						<label for="description" class="mb-2 block text-sm font-medium text-text">
							Description *
						</label>
						<textarea
							id="description"
							rows="3"
							bind:value={formData.description}
							class="w-full resize-none rounded-lg border border-white/10 bg-card px-4 py-2.5 text-text placeholder-text-muted focus:ring-2 focus:ring-white/20 focus:outline-none"
							placeholder="e.g., A friendly software engineer who loves teaching others and has a quirky sense of humor"
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
							<label for="pronouns" class="mb-2 block text-sm font-medium text-text">
								Pronouns *
							</label>
							<input
								id="pronouns"
								type="text"
								bind:value={formData.pronouns}
								class="w-full rounded-lg border border-white/10 bg-card px-4 py-2.5 text-text placeholder-text-muted focus:ring-2 focus:ring-white/20 focus:outline-none"
								placeholder="e.g., he/him"
							/>
						</div>
					</div>
				</div>
			</div>

			<div class="my-12 border-t border-white/10"></div>

			<!-- Big Five traits Section -->
			<div class="mb-12">
				<h3 class="mb-2 text-xl font-semibold text-text">Big Five Personality Traits</h3>
				<p class="mb-6 text-sm text-text-muted">
					Adjust these traits to shape how your persona thinks and communicates.
				</p>

				<!-- Openness -->
				<div class="mb-8">
					<label for="openness" class="mb-2 block text-sm font-medium text-text">
						Openness
						<span class="ml-2 text-text-muted">({formData.openness.toFixed(1)})</span>
					</label>
					<p class="mb-3 text-xs text-text-muted">{oceanTraits.openness.description}</p>

					<input
						id="openness"
						type="range"
						min="0"
						max="10"
						step="0.1"
						bind:value={formData.openness}
						on:pointerdown={handleSliderFocus}
						on:touchstart={handleSliderFocus}
						class="h-2 w-full cursor-pointer appearance-none rounded-lg bg-card accent-text"
					/>
					<div class="mt-1 flex justify-between text-xs text-text-muted">
						<span>{oceanTraits.openness.lowLabel}</span>
						<span>{oceanTraits.openness.highLabel}</span>
					</div>
				</div>

				<!-- Conscientiousness -->
				<div class="mb-8">
					<label for="conscientiousness" class="mb-2 block text-sm font-medium text-text">
						Conscientiousness
						<span class="ml-2 text-text-muted">({formData.conscientiousness.toFixed(1)})</span>
					</label>
					<p class="mb-3 text-xs text-text-muted">{oceanTraits.conscientiousness.description}</p>

					<input
						id="conscientiousness"
						type="range"
						min="0"
						max="10"
						step="0.1"
						bind:value={formData.conscientiousness}
						on:pointerdown={handleSliderFocus}
						on:touchstart={handleSliderFocus}
						class="h-2 w-full cursor-pointer appearance-none rounded-lg bg-card accent-text"
					/>
					<div class="mt-1 flex justify-between text-xs text-text-muted">
						<span>{oceanTraits.conscientiousness.lowLabel}</span>
						<span>{oceanTraits.conscientiousness.highLabel}</span>
					</div>
				</div>

				<!-- Extraversion -->
				<div class="mb-8">
					<label for="extraversion" class="mb-2 block text-sm font-medium text-text">
						Extraversion
						<span class="ml-2 text-text-muted">({formData.extraversion.toFixed(1)})</span>
					</label>
					<p class="mb-3 text-xs text-text-muted">{oceanTraits.extraversion.description}</p>

					<input
						id="extraversion"
						type="range"
						min="0"
						max="10"
						step="0.1"
						bind:value={formData.extraversion}
						on:pointerdown={handleSliderFocus}
						on:touchstart={handleSliderFocus}
						class="h-2 w-full cursor-pointer appearance-none rounded-lg bg-card accent-text"
					/>
					<div class="mt-1 flex justify-between text-xs text-text-muted">
						<span>{oceanTraits.extraversion.lowLabel}</span>
						<span>{oceanTraits.extraversion.highLabel}</span>
					</div>
				</div>

				<!-- Agreeableness -->
				<div class="mb-8">
					<label for="agreeableness" class="mb-2 block text-sm font-medium text-text">
						Agreeableness
						<span class="ml-2 text-text-muted">({formData.agreeableness.toFixed(1)})</span>
					</label>
					<p class="mb-3 text-xs text-text-muted">{oceanTraits.agreeableness.description}</p>

					<input
						id="agreeableness"
						type="range"
						min="0"
						max="10"
						step="0.1"
						bind:value={formData.agreeableness}
						on:pointerdown={handleSliderFocus}
						on:touchstart={handleSliderFocus}
						class="h-2 w-full cursor-pointer appearance-none rounded-lg bg-card accent-text"
					/>
					<div class="mt-1 flex justify-between text-xs text-text-muted">
						<span>{oceanTraits.agreeableness.lowLabel}</span>
						<span>{oceanTraits.agreeableness.highLabel}</span>
					</div>
				</div>

				<!-- Neuroticism -->
				<div class="mb-8">
					<label for="neuroticism" class="mb-2 block text-sm font-medium text-text">
						Neuroticism
						<span class="ml-2 text-text-muted">({formData.neuroticism.toFixed(1)})</span>
					</label>
					<p class="mb-3 text-xs text-text-muted">{oceanTraits.neuroticism.description}</p>

					<input
						id="neuroticism"
						type="range"
						min="0"
						max="10"
						step="0.1"
						bind:value={formData.neuroticism}
						on:pointerdown={handleSliderFocus}
						on:touchstart={handleSliderFocus}
						class="h-2 w-full cursor-pointer appearance-none rounded-lg bg-card accent-text"
					/>
					<div class="mt-1 flex justify-between text-xs text-text-muted">
						<span>{oceanTraits.neuroticism.lowLabel}</span>
						<span>{oceanTraits.neuroticism.highLabel}</span>
					</div>
				</div>
			</div>

			<div class="my-12 border-t border-white/10"></div>

			<!-- Additional traits Section -->
			<div class="mb-12">
				<h3 class="mb-2 text-xl font-semibold text-text">Additional Traits</h3>
				<p class="mb-6 text-sm text-text-muted">
					Fine-tune your persona's cognitive and emotional characteristics.
				</p>

				<!-- Intelligence -->
				<div class="mb-8">
					<label for="intelligence" class="mb-2 block text-sm font-medium text-text">
						Intelligence
						<span class="ml-2 text-text-muted">({formData.intelligence.toFixed(1)})</span>
					</label>
					<p class="mb-3 text-xs text-text-muted">{additionalTraits.intelligence.description}</p>

					<input
						id="intelligence"
						type="range"
						min="0"
						max="10"
						step="0.1"
						bind:value={formData.intelligence}
						on:pointerdown={handleSliderFocus}
						on:touchstart={handleSliderFocus}
						class="h-2 w-full cursor-pointer appearance-none rounded-lg bg-card accent-text"
					/>
					<div class="mt-1 flex justify-between text-xs text-text-muted">
						<span>{additionalTraits.intelligence.lowLabel}</span>
						<span>{additionalTraits.intelligence.highLabel}</span>
					</div>
				</div>

				<!-- Thinking Style -->
				<div class="mb-8">
					<label for="thinking_style" class="mb-2 block text-sm font-medium text-text">
						Thinking Style
						<span class="ml-2 text-text-muted">({formData.thinking_style.toFixed(1)})</span>
					</label>
					<p class="mb-3 text-xs text-text-muted">{additionalTraits.thinking_style.description}</p>

					<input
						id="thinking_style"
						type="range"
						min="0"
						max="10"
						step="0.1"
						bind:value={formData.thinking_style}
						on:pointerdown={handleSliderFocus}
						on:touchstart={handleSliderFocus}
						class="h-2 w-full cursor-pointer appearance-none rounded-lg bg-card accent-text"
					/>
					<div class="mt-1 flex justify-between text-xs text-text-muted">
						<span>{additionalTraits.thinking_style.lowLabel}</span>
						<span>{additionalTraits.thinking_style.highLabel}</span>
					</div>
				</div>

				<!-- Tone -->
				<div class="mb-8">
					<label for="tone" class="mb-2 block text-sm font-medium text-text"> Tone * </label>
					<p class="mb-3 text-xs text-text-muted">
						The overall style and manner of communication (e.g., friendly, professional, casual,
						warm).
					</p>
					<input
						id="tone"
						type="text"
						bind:value={formData.tone}
						class="w-full rounded-lg border border-white/10 bg-card px-4 py-2.5 text-text placeholder-text-muted focus:ring-2 focus:ring-white/20 focus:outline-none"
						placeholder="e.g., friendly, professional, casual"
					/>
				</div>

				<!-- Humor Level -->
				<div class="mb-8">
					<label for="humor_level" class="mb-2 block text-sm font-medium text-text">
						Humor Level
						<span class="ml-2 text-text-muted">({formData.humor_level.toFixed(1)})</span>
					</label>
					<p class="mb-3 text-xs text-text-muted">{additionalTraits.humor_level.description}</p>

					<input
						id="humor_level"
						type="range"
						min="0"
						max="10"
						step="0.1"
						bind:value={formData.humor_level}
						on:pointerdown={handleSliderFocus}
						on:touchstart={handleSliderFocus}
						class="h-2 w-full cursor-pointer appearance-none rounded-lg bg-card accent-text"
					/>
					<div class="mt-1 flex justify-between text-xs text-text-muted">
						<span>{additionalTraits.humor_level.lowLabel}</span>
						<span>{additionalTraits.humor_level.highLabel}</span>
					</div>
				</div>

				<!-- Mood Fluctuation -->
				<div class="mb-8">
					<label for="mood_fluctuation" class="mb-2 block text-sm font-medium text-text">
						Mood Fluctuation
						<span class="ml-2 text-text-muted">({formData.mood_fluctuation.toFixed(1)})</span>
					</label>
					<p class="mb-3 text-xs text-text-muted">
						{additionalTraits.mood_fluctuation.description}
					</p>

					<input
						id="mood_fluctuation"
						type="range"
						min="0"
						max="10"
						step="0.1"
						bind:value={formData.mood_fluctuation}
						on:pointerdown={handleSliderFocus}
						on:touchstart={handleSliderFocus}
						class="h-2 w-full cursor-pointer appearance-none rounded-lg bg-card accent-text"
					/>
					<div class="mt-1 flex justify-between text-xs text-text-muted">
						<span>{additionalTraits.mood_fluctuation.lowLabel}</span>
						<span>{additionalTraits.mood_fluctuation.highLabel}</span>
					</div>
				</div>
			</div>

			<div class="my-12 border-t border-white/10"></div>

			<!-- Interests Section -->
			<div class="mb-12">
				<h3 class="mb-2 text-xl font-semibold text-text">Interests</h3>
				<p class="mb-6 text-sm text-text-muted">
					Define what your persona enjoys and dislikes. Separate items with commas.
				</p>

				<div class="space-y-6">
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

			<div class="my-12 border-t border-white/10"></div>

			<!-- Language Section -->
			<div class="mb-12">
				<h3 class="mb-2 text-xl font-semibold text-text">Language Style</h3>
				<p class="mb-6 text-sm text-text-muted">
					Configure how your persona communicates and expresses themselves.
				</p>

				<div class="mb-8">
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
						on:pointerdown={handleSliderFocus}
						on:touchstart={handleSliderFocus}
						class="h-2 w-full cursor-pointer appearance-none rounded-lg bg-card accent-text"
					/>
					<div class="mt-1 flex justify-between text-xs text-text-muted">
						<span>Casual</span>
						<span>Formal</span>
					</div>
				</div>

				<div class="mb-8">
					<label for="fluency" class="mb-2 block text-sm font-medium text-text">Fluency *</label>
					<select
						id="fluency"
						bind:value={formData.fluency}
						class="w-full rounded-lg border border-white/10 bg-card px-4 py-2.5 text-text focus:ring-2 focus:ring-white/20 focus:outline-none"
					>
						<option value="" disabled selected>Select fluency level</option>
						<option value="beginner">Beginner</option>
						<option value="intermediate">Intermediate</option>
						<option value="proficient">Proficient</option>
					</select>
				</div>
			</div>

			<div class="my-12 border-t border-white/10"></div>

			<!-- Others Section -->
			<div class="mb-12">
				<h3 class="mb-2 text-xl font-semibold text-text">Other Settings</h3>
				<p class="mb-6 text-sm text-text-muted">
					Additional preferences for your persona's communication style.
				</p>

				<div class="space-y-6">
					<div>
						<label for="emoji_usage" class="mb-2 block text-sm font-medium text-text"
							>Emoji Usage *</label
						>
						<select
							id="emoji_usage"
							bind:value={formData.emoji_usage}
							class="w-full rounded-lg border border-white/10 bg-card px-4 py-2.5 text-text focus:ring-2 focus:ring-white/20 focus:outline-none"
						>
							<option value="" disabled selected>Select emoji usage level</option>
							<option value="frequent">Frequent</option>
							<option value="moderate">Moderate</option>
							<option value="rare">Rare</option>
							<option value="none">None</option>
						</select>
					</div>

					<div>
						<label for="response_length" class="mb-2 block text-sm font-medium text-text"
							>Response Length *</label
						>
						<select
							id="response_length"
							bind:value={formData.response_length}
							class="w-full rounded-lg border border-white/10 bg-card px-4 py-2.5 text-text focus:ring-2 focus:ring-white/20 focus:outline-none"
						>
							<option value="" disabled selected>Select response length</option>
							<option value="concise">Concise</option>
							<option value="moderate">Moderate</option>
							<option value="detailed">Detailed</option>
						</select>
					</div>
				</div>
			</div>

			<!-- Submit Button -->
			<div class="flex justify-end border-t border-white/10 pt-6">
				<button
					on:click={handleSubmit}
					disabled={!isFormValid || isLoading}
					class="inline-flex min-w-[140px] items-center justify-center gap-2 rounded-lg bg-text px-6 py-3 font-medium text-background transition hover:bg-text/90 disabled:cursor-not-allowed disabled:opacity-50"
				>
					{#if isLoading}
						<div
							class="h-4 w-4 animate-spin rounded-full border-2 border-current border-r-transparent"
						></div>
					{:else}
						Create Persona
					{/if}
				</button>
			</div>
		</div>
	</div>
</div>
