<script lang="ts">
	import personaFormImg from '$lib/assets/screenshots/persona-form.png';
	import chatInterfaceImg from '$lib/assets/screenshots/chat-interface.png';
	import personaFormMobileImg from '$lib/assets/screenshots/persona-form-mobile.png';
	import chatInterfaceMobileImg from '$lib/assets/screenshots/chat-interface-mobile.png';
	import githubIcon from '$lib/assets/icons/github-logo.png';

	let isOpen = false;

	function closeDrawer() {
		isOpen = false;
	}
</script>

<svelte:head>
	<link rel="preconnect" href="https://fonts.googleapis.com" />
	<link rel="preconnect" href="https://fonts.gstatic.com" crossorigin="anonymous" />
	<link
		href="https://fonts.googleapis.com/css2?family=Pixelify+Sans:wght@600&display=swap"
		rel="stylesheet"
	/>
</svelte:head>

<div class="flex min-h-screen overflow-x-hidden bg-background text-text">
	<!-- ============================================ -->
	<!-- LEFT NAVBAR -->
	<!-- ============================================ -->

	<!-- Mobile Menu Button -->
	<button
		on:click={() => (isOpen = true)}
		class="fixed top-4 left-4 z-40 rounded-md border border-white/10 bg-card p-2 text-text transition-colors hover:bg-card-hover lg:hidden"
		aria-label="Open menu"
	>
		<svg
			width="20"
			height="20"
			viewBox="0 0 24 24"
			fill="none"
			stroke="currentColor"
			stroke-width="2"
		>
			<path d="M3 12h18M3 6h18M3 18h18" />
		</svg>
	</button>

	<!-- Overlay (mobile only) -->
	{#if isOpen}
		<button
			on:click={closeDrawer}
			class="fixed inset-0 z-40 bg-background/80 backdrop-blur-sm lg:hidden"
			aria-label="Close menu"
		></button>
	{/if}

	<!-- Navbar -->
	<nav
		class="fixed top-0 left-0 z-50 flex h-full w-64 flex-col border-r border-white/10 bg-background transition-transform duration-300 lg:translate-x-0 {isOpen
			? 'translate-x-0'
			: '-translate-x-full'}"
	>
		<!-- Logo -->
		<div class="border-b border-white/10 px-8 py-8">
			<a
				href="/"
				class="text-xl font-semibold tracking-tight"
				style="font-family: 'Pixelify Sans', sans-serif;"
			>
				Personabox
			</a>
		</div>

		<!-- Navigation Links -->
		<div class="flex flex-1 flex-col px-8 py-12">
			<ul class="space-y-5">
				<li>
					<a
						href="#science"
						class="block text-[15px] font-medium text-text-muted transition-colors hover:text-text"
					>
						The Science
					</a>
				</li>
				<li>
					<a
						href="#product"
						class="block text-[15px] font-medium text-text-muted transition-colors hover:text-text"
					>
						Product
					</a>
				</li>
				<li>
					<a
						href="#how-it-works"
						class="block text-[15px] font-medium text-text-muted transition-colors hover:text-text"
					>
						How It Works
					</a>
				</li>
				<li>
					<a
						href="#for-writers"
						class="block text-[15px] font-medium text-text-muted transition-colors hover:text-text"
					>
						For Writers
					</a>
				</li>
				<li>
					<a
						href="#other-uses"
						class="block text-[15px] font-medium text-text-muted transition-colors hover:text-text"
					>
						Other Uses
					</a>
				</li>
			</ul>
		</div>
	</nav>

	<!-- ============================================ -->
	<!-- MAIN CONTENT -->
	<!-- ============================================ -->

	<main class="w-full flex-1 lg:ml-64">
		<!-- ============================================ -->
		<!-- TOP RIGHT BUTTONS -->
		<!-- ============================================ -->

		<div class="absolute top-6 right-8 z-30 flex items-center gap-4">
			<a
				href="https://github.com/aDiThYa-808/persona-box"
				target="_blank"
				rel="noopener noreferrer"
				class="transition-opacity hover:opacity-70"
				aria-label="GitHub"
			>
				<img src={githubIcon} alt="GitHub" class="h-4 w-auto" />
			</a>
			<a
				href="/login"
				class="text-[15px] font-medium text-text-muted transition-colors hover:text-text"
			>
				Login
			</a>
		</div>

		<!-- ============================================ -->
		<!-- HERO SECTION -->
		<!-- ============================================ -->

		<section class="flex min-h-screen flex-col items-center justify-center px-6 py-20 md:px-8">
			<div class="w-full max-w-5xl text-center">
				<p class="mb-6 text-xs font-semibold tracking-[0.2em] text-text-muted uppercase">
					Personality-Driven AI Chat
				</p>
				<h1
					class="mb-8 text-5xl leading-[1.08] font-bold tracking-[-0.02em] md:text-6xl lg:text-7xl xl:text-[5.5rem]"
				>
					Design AI personas with<br />real personality traits
				</h1>
				<div class="mx-auto max-w-3xl space-y-3">
					<p class="text-5 font-medium text-text md:text-2xl">
						Built on the Big Five personality model.
					</p>
					<p class="text-base leading-relaxed text-text-muted md:text-lg">
						Create AI personas with authentic psychological depth for character development,<br
							class="hidden md:inline"
						/> dialogue testing, and creative exploration.
					</p>
				</div>
				<div class="mt-12">
					<a
						href="/chat/new-persona"
						class="inline-block rounded-lg border-2 border-text bg-text px-8 py-3.5 text-base font-semibold text-background transition-all hover:bg-background hover:text-text"
					>
						Start Creating (Free)
					</a>
				</div>
			</div>
		</section>

		<!-- ============================================ -->
		<!-- BIG FIVE MODEL EXPLANATION -->
		<!-- ============================================ -->

		<section id="science" class="border-t border-white/10 px-6 py-24 md:px-8 md:py-40">
			<div class="mx-auto w-full max-w-6xl">
				<div class="mb-16 md:mb-20">
					<p class="mb-3 text-xs font-semibold tracking-[0.2em] text-text-muted uppercase">
						The Science
					</p>
					<h2
						class="mb-6 text-4xl leading-tight font-bold tracking-[-0.02em] md:text-5xl lg:text-6xl"
					>
						The Big Five<br />Personality Model
					</h2>
					<p class="max-w-3xl text-lg leading-relaxed text-text-muted md:text-xl">
						Unlike arbitrary personality types, the Big Five is the most scientifically validated
						framework in psychology. It measures personality across five core dimensions that
						predict how people think, feel, and behave.
					</p>
				</div>

				<!-- Mobile: Cards -->
				<div class="space-y-4 md:hidden">
					<!-- Openness -->
					<div
						class="overflow-hidden rounded-lg border border-white/10 bg-card shadow-2xl shadow-black/10"
					>
						<div class="border-b border-white/10 bg-background px-6 py-4">
							<div
								class="mb-2 inline-block rounded border border-white/10 bg-card px-2.5 py-1 text-xs font-semibold tracking-wider text-text-muted uppercase"
							>
								O
							</div>
							<div class="mt-2 text-lg font-bold">Openness</div>
						</div>
						<div class="px-6 py-6">
							<p class="mb-4 text-sm leading-relaxed text-text-muted">
								Imagination, curiosity, and preference for novelty. High openness means creative,
								abstract thinking. Low openness means practical, conventional approaches.
							</p>
							<p class="text-sm leading-relaxed text-text-muted">
								<span class="font-semibold text-text">In conversation:</span> Affects whether they explore
								hypotheticals, use metaphors, or stick to concrete facts.
							</p>
						</div>
					</div>

					<!-- Conscientiousness -->
					<div
						class="overflow-hidden rounded-lg border border-white/10 bg-card shadow-2xl shadow-black/10"
					>
						<div class="border-b border-white/10 bg-background px-6 py-4">
							<div
								class="mb-2 inline-block rounded border border-white/10 bg-card px-2.5 py-1 text-xs font-semibold tracking-wider text-text-muted uppercase"
							>
								C
							</div>
							<div class="mt-2 text-lg font-bold">Conscientiousness</div>
						</div>
						<div class="px-6 py-6">
							<p class="mb-4 text-sm leading-relaxed text-text-muted">
								Organization, discipline, and goal-orientation. High conscientiousness means
								structured and reliable. Low conscientiousness means spontaneous and flexible.
							</p>
							<p class="text-sm leading-relaxed text-text-muted">
								<span class="font-semibold text-text">In conversation:</span> Influences planning style,
								attention to detail, and follow-through on ideas.
							</p>
						</div>
					</div>

					<!-- Extraversion -->
					<div
						class="overflow-hidden rounded-lg border border-white/10 bg-card shadow-2xl shadow-black/10"
					>
						<div class="border-b border-white/10 bg-background px-6 py-4">
							<div
								class="mb-2 inline-block rounded border border-white/10 bg-card px-2.5 py-1 text-xs font-semibold tracking-wider text-text-muted uppercase"
							>
								E
							</div>
							<div class="mt-2 text-lg font-bold">Extraversion</div>
						</div>
						<div class="px-6 py-6">
							<p class="mb-4 text-sm leading-relaxed text-text-muted">
								Energy from social interaction and external stimulation. High extraversion means
								outgoing and energetic. Low extraversion means reserved and introspective.
							</p>
							<p class="text-sm leading-relaxed text-text-muted">
								<span class="font-semibold text-text">In conversation:</span> Shapes their enthusiasm
								level, verbosity, and how they engage with topics.
							</p>
						</div>
					</div>

					<!-- Agreeableness -->
					<div
						class="overflow-hidden rounded-lg border border-white/10 bg-card shadow-2xl shadow-black/10"
					>
						<div class="border-b border-white/10 bg-background px-6 py-4">
							<div
								class="mb-2 inline-block rounded border border-white/10 bg-card px-2.5 py-1 text-xs font-semibold tracking-wider text-text-muted uppercase"
							>
								A
							</div>
							<div class="mt-2 text-lg font-bold">Agreeableness</div>
						</div>
						<div class="px-6 py-6">
							<p class="mb-4 text-sm leading-relaxed text-text-muted">
								Compassion, cooperation, and trust. High agreeableness means warm and empathetic.
								Low agreeableness means competitive and skeptical.
							</p>
							<p class="text-sm leading-relaxed text-text-muted">
								<span class="font-semibold text-text">In conversation:</span> Determines how they handle
								disagreement, show empathy, and critique ideas.
							</p>
						</div>
					</div>

					<!-- Neuroticism -->
					<div
						class="overflow-hidden rounded-lg border border-white/10 bg-card shadow-2xl shadow-black/10"
					>
						<div class="border-b border-white/10 bg-background px-6 py-4">
							<div
								class="mb-2 inline-block rounded border border-white/10 bg-card px-2.5 py-1 text-xs font-semibold tracking-wider text-text-muted uppercase"
							>
								N
							</div>
							<div class="mt-2 text-lg font-bold">Neuroticism</div>
						</div>
						<div class="px-6 py-6">
							<p class="mb-4 text-sm leading-relaxed text-text-muted">
								Emotional sensitivity and tendency toward negative emotions. High neuroticism means
								anxious and reactive. Low neuroticism means calm and stable.
							</p>
							<p class="text-sm leading-relaxed text-text-muted">
								<span class="font-semibold text-text">In conversation:</span> Affects their emotional
								range, stress responses, and overall tone.
							</p>
						</div>
					</div>
				</div>

				<!-- Desktop: Table -->
				<div
					class="hidden overflow-hidden rounded-lg border border-white/10 bg-card shadow-2xl shadow-black/10 md:block"
				>
					<table class="w-full">
						<thead>
							<tr class="border-b border-white/10">
								<th
									class="px-8 py-5 text-left text-sm font-semibold tracking-wider text-text-muted uppercase"
								>
									Trait
								</th>
								<th
									class="px-8 py-5 text-left text-sm font-semibold tracking-wider text-text-muted uppercase"
								>
									Description
								</th>
								<th
									class="px-8 py-5 text-left text-sm font-semibold tracking-wider text-text-muted uppercase"
								>
									In Conversation
								</th>
							</tr>
						</thead>
						<tbody>
							<!-- Openness -->
							<tr class="border-b border-white/10">
								<td class="px-8 py-8 align-top">
									<div
										class="mb-2 inline-block rounded border border-white/10 bg-background px-2.5 py-1 text-xs font-semibold tracking-wider text-text-muted uppercase"
									>
										O
									</div>
									<div class="mt-2 text-xl font-bold">Openness</div>
								</td>
								<td class="px-8 py-8 align-top">
									<p class="text-base leading-relaxed text-text-muted">
										Imagination, curiosity, and preference for novelty. High openness means
										creative, abstract thinking. Low openness means practical, conventional
										approaches.
									</p>
								</td>
								<td class="px-8 py-8 align-top">
									<p class="text-base leading-relaxed text-text-muted">
										Affects whether they explore hypotheticals, use metaphors, or stick to concrete
										facts.
									</p>
								</td>
							</tr>

							<!-- Conscientiousness -->
							<tr class="border-b border-white/10">
								<td class="px-8 py-8 align-top">
									<div
										class="mb-2 inline-block rounded border border-white/10 bg-background px-2.5 py-1 text-xs font-semibold tracking-wider text-text-muted uppercase"
									>
										C
									</div>
									<div class="mt-2 text-xl font-bold">Conscientiousness</div>
								</td>
								<td class="px-8 py-8 align-top">
									<p class="text-base leading-relaxed text-text-muted">
										Organization, discipline, and goal-orientation. High conscientiousness means
										structured and reliable. Low conscientiousness means spontaneous and flexible.
									</p>
								</td>
								<td class="px-8 py-8 align-top">
									<p class="text-base leading-relaxed text-text-muted">
										Influences planning style, attention to detail, and follow-through on ideas.
									</p>
								</td>
							</tr>

							<!-- Extraversion -->
							<tr class="border-b border-white/10">
								<td class="px-8 py-8 align-top">
									<div
										class="mb-2 inline-block rounded border border-white/10 bg-background px-2.5 py-1 text-xs font-semibold tracking-wider text-text-muted uppercase"
									>
										E
									</div>
									<div class="mt-2 text-xl font-bold">Extraversion</div>
								</td>
								<td class="px-8 py-8 align-top">
									<p class="text-base leading-relaxed text-text-muted">
										Energy from social interaction and external stimulation. High extraversion means
										outgoing and energetic. Low extraversion means reserved and introspective.
									</p>
								</td>
								<td class="px-8 py-8 align-top">
									<p class="text-base leading-relaxed text-text-muted">
										Shapes their enthusiasm level, verbosity, and how they engage with topics.
									</p>
								</td>
							</tr>

							<!-- Agreeableness -->
							<tr class="border-b border-white/10">
								<td class="px-8 py-8 align-top">
									<div
										class="mb-2 inline-block rounded border border-white/10 bg-background px-2.5 py-1 text-xs font-semibold tracking-wider text-text-muted uppercase"
									>
										A
									</div>
									<div class="mt-2 text-xl font-bold">Agreeableness</div>
								</td>
								<td class="px-8 py-8 align-top">
									<p class="text-base leading-relaxed text-text-muted">
										Compassion, cooperation, and trust. High agreeableness means warm and
										empathetic. Low agreeableness means competitive and skeptical.
									</p>
								</td>
								<td class="px-8 py-8 align-top">
									<p class="text-base leading-relaxed text-text-muted">
										Determines how they handle disagreement, show empathy, and critique ideas.
									</p>
								</td>
							</tr>

							<!-- Neuroticism -->
							<tr>
								<td class="px-8 py-8 align-top">
									<div
										class="mb-2 inline-block rounded border border-white/10 bg-background px-2.5 py-1 text-xs font-semibold tracking-wider text-text-muted uppercase"
									>
										N
									</div>
									<div class="mt-2 text-xl font-bold">Neuroticism</div>
								</td>
								<td class="px-8 py-8 align-top">
									<p class="text-base leading-relaxed text-text-muted">
										Emotional sensitivity and tendency toward negative emotions. High neuroticism
										means anxious and reactive. Low neuroticism means calm and stable.
									</p>
								</td>
								<td class="px-8 py-8 align-top">
									<p class="text-base leading-relaxed text-text-muted">
										Affects their emotional range, stress responses, and overall tone.
									</p>
								</td>
							</tr>
						</tbody>
					</table>
				</div>

				<!-- Why This Matters -->
				<div
					class="mt-16 rounded-lg border border-white/10 bg-card p-8 shadow-2xl shadow-black/10 md:mt-20 md:p-12"
				>
					<h3 class="mb-4 text-2xl font-bold md:text-3xl">Why This Model Matters</h3>
					<p class="text-base leading-relaxed text-text-muted md:text-lg">
						Random personality traits produce random behavior. The Big Five creates <span
							class="font-semibold text-text">consistent, predictable personalities</span
						> because it's based on how human personality actually works. Each trait interacts with the
						others, creating complex but coherent characters, just like real people.
					</p>
				</div>
			</div>
		</section>
		<!-- ============================================ -->
		<!-- PERSONA CREATION SHOWCASE -->
		<!-- ============================================ -->

		<section id="product" class="border-t border-white/10 px-6 py-24 md:px-8 md:py-40">
			<div class="mx-auto w-full max-w-6xl">
				<div class="mb-16 md:mb-20">
					<p class="mb-3 text-xs font-semibold tracking-[0.2em] text-text-muted uppercase">
						The Product
					</p>
					<h2
						class="mb-6 text-4xl leading-tight font-bold tracking-[-0.02em] md:text-5xl lg:text-6xl"
					>
						Fine-tune every<br />dimension
					</h2>
					<p class="max-w-3xl text-lg leading-relaxed text-text-muted md:text-xl">
						Adjust all five traits plus additional specific facets. Create personas ranging from
						anxious perfectionists to carefree adventurers, and everything in between.
					</p>
				</div>

				<!-- Desktop Screenshot -->
				<div
					class="hidden overflow-hidden rounded-lg border border-white/10 bg-card shadow-2xl shadow-black/10 md:block"
				>
					<img
						src={personaFormImg}
						alt="Persona creation form showing Big Five personality sliders and traits"
						class="w-full"
					/>
				</div>

				<!-- Mobile Screenshot -->
				<div
					class="overflow-hidden rounded-lg border border-white/10 bg-card shadow-2xl shadow-black/10 md:hidden"
				>
					<img
						src={personaFormMobileImg}
						alt="Persona creation form showing Big Five personality sliders and traits"
						class="w-full"
					/>
				</div>
			</div>
		</section>

		<!-- ============================================ -->
		<!-- CHAT EXPERIENCE SHOWCASE -->
		<!-- ============================================ -->

		<section class="px-6 py-24 md:px-8 md:py-40">
			<div class="mx-auto w-full max-w-6xl">
				<div class="mb-16 md:mb-20">
					<h2
						class="mb-6 text-4xl leading-tight font-bold tracking-[-0.02em] md:text-5xl lg:text-6xl"
					>
						Personalities that<br />stay consistent
					</h2>
					<p class="max-w-3xl text-lg leading-relaxed text-text-muted md:text-xl">
						Each persona maintains their unique perspective across conversations. The personality
						traits you set directly influence how they think, respond, and interact.
					</p>
				</div>

				<!-- Desktop Screenshot -->
				<div
					class="hidden overflow-hidden rounded-lg border border-white/10 bg-card shadow-2xl shadow-black/10 md:block"
				>
					<img
						src={chatInterfaceImg}
						alt="Chat interface showing conversation with AI persona"
						class="w-full"
					/>
				</div>

				<!-- Mobile Screenshot -->
				<div
					class="overflow-hidden rounded-lg border border-white/10 bg-card shadow-2xl shadow-black/10 md:hidden"
				>
					<img
						src={chatInterfaceMobileImg}
						alt="Chat interface showing conversation with AI persona"
						class="w-full"
					/>
				</div>
			</div>
		</section>

		<!-- ============================================ -->
		<!-- HOW IT WORKS -->
		<!-- ============================================ -->

		<section id="how-it-works" class="border-t border-white/10 px-6 py-24 md:px-8 md:py-40">
			<div class="mx-auto w-full max-w-6xl">
				<div class="mb-20 md:mb-28">
					<p class="mb-3 text-xs font-semibold tracking-[0.2em] text-text-muted uppercase">
						How It Works
					</p>
					<h2 class="text-4xl leading-tight font-bold tracking-[-0.02em] md:text-5xl lg:text-6xl">
						Three steps to<br />realistic characters
					</h2>
				</div>

				<div class="space-y-20 md:space-y-28">
					<!-- Step 1 -->
					<div class="grid gap-8 md:grid-cols-[120px_1fr] md:gap-16">
						<div>
							<div
								class="text-5xl font-bold text-text-muted md:text-6xl"
								style="font-family: 'Pixelify Sans', sans-serif;"
							>
								01
							</div>
						</div>
						<div>
							<h3 class="mb-4 text-2xl font-bold md:text-3xl">Design</h3>
							<p class="text-base leading-relaxed text-text-muted md:text-lg">
								Adjust the Big Five traits and specific facets. High conscientiousness and low
								openness? You get a detail-oriented traditionalist. Low agreeableness and high
								openness? A provocative contrarian.
							</p>
						</div>
					</div>

					<!-- Step 2 -->
					<div class="grid gap-8 md:grid-cols-[120px_1fr] md:gap-16">
						<div>
							<div
								class="text-5xl font-bold text-text-muted md:text-6xl"
								style="font-family: 'Pixelify Sans', sans-serif;"
							>
								02
							</div>
						</div>
						<div>
							<h3 class="mb-4 text-2xl font-bold md:text-3xl">Chat</h3>
							<p class="text-base leading-relaxed text-text-muted md:text-lg">
								Have conversations that reveal the personality. Test dialogue, explore character
								motivations, or see how they'd react to specific scenarios. The traits you set shape
								every response.
							</p>
						</div>
					</div>

					<!-- Step 3 -->
					<div class="grid gap-8 md:grid-cols-[120px_1fr] md:gap-16">
						<div>
							<div
								class="text-5xl font-bold text-text-muted md:text-6xl"
								style="font-family: 'Pixelify Sans', sans-serif;"
							>
								03
							</div>
						</div>
						<div>
							<h3 class="mb-4 text-2xl font-bold md:text-3xl">Iterate</h3>
							<p class="text-base leading-relaxed text-text-muted md:text-lg">
								Explore different personalities by creating new personas. Start multiple
								conversations with the same persona to see how consistently they respond, or build
								variations to compare how different traits change the interaction.
							</p>
						</div>
					</div>
				</div>
			</div>
		</section>

		<!-- ============================================ -->
		<!-- FOR WRITERS -->
		<!-- ============================================ -->

		<section id="for-writers" class="border-t border-white/10 px-6 py-24 md:px-8 md:py-40">
			<div class="mx-auto w-full max-w-6xl">
				<div class="mb-16 md:mb-20">
					<p class="mb-3 text-xs font-semibold tracking-[0.2em] text-text-muted uppercase">
						For Writers
					</p>
					<h2
						class="mb-6 text-4xl leading-tight font-bold tracking-[-0.02em] md:text-5xl lg:text-6xl"
					>
						Build characters<br />with depth
					</h2>
					<p class="max-w-3xl text-lg leading-relaxed text-text-muted md:text-xl">
						Go beyond surface-level character traits. Create personas that think and respond like
						complete human beings.
					</p>
				</div>

				<div class="grid gap-6 md:grid-cols-2 md:gap-8">
					<!-- Use Case 1 -->
					<div
						class="rounded-lg border border-white/10 bg-card p-8 transition-colors hover:bg-card-hover md:p-10"
					>
						<h3 class="mb-4 text-xl font-bold md:text-2xl">Character Development & Dialogue</h3>
						<p class="text-base leading-relaxed text-text-muted">
							Test your character's perspective. See how their personality traits influence word
							choice, argument style, and emotional responses. Catch inconsistencies before they
							make it to the page.
						</p>
					</div>

					<!-- Use Case 2 -->
					<div
						class="rounded-lg border border-white/10 bg-card p-8 transition-colors hover:bg-card-hover md:p-10"
					>
						<h3 class="mb-4 text-xl font-bold md:text-2xl">Explore Different Reactions</h3>
						<p class="text-base leading-relaxed text-text-muted">
							Drop the same scenario into conversations with different personas. See how a neurotic
							introvert vs. a confident extravert would handle the same conflict. Use it to make
							your characters feel distinct and real.
						</p>
					</div>
				</div>

				<!-- CTA -->
				<div class="mt-16 text-center md:mt-20">
					<a
						href="/chat/new-persona"
						class="inline-block rounded-lg border-2 border-text bg-text px-8 py-3.5 text-base font-semibold text-background transition-all hover:bg-background hover:text-text"
					>
						Create Your First Persona (Free)
					</a>
				</div>
			</div>
		</section>

		<!-- ============================================ -->
		<!-- OTHER USES -->
		<!-- ============================================ -->

		<section id="other-uses" class="border-t border-white/10 px-6 py-24 md:px-8 md:py-40">
			<div class="mx-auto w-full max-w-6xl">
				<div class="mb-16 md:mb-20">
					<p class="mb-3 text-xs font-semibold tracking-[0.2em] text-text-muted uppercase">
						Other Uses
					</p>
					<h2
						class="mb-6 text-4xl leading-tight font-bold tracking-[-0.02em] md:text-5xl lg:text-6xl"
					>
						Or just experiment<br />for fun
					</h2>
					<p class="max-w-3xl text-lg leading-relaxed text-text-muted md:text-xl">
						Not everything needs a purpose. Create personalities, see how they think, and explore
						what makes people different.
					</p>
				</div>

				<div class="grid gap-6 md:grid-cols-2 md:gap-8">
					<!-- Use Case 1 -->
					<div
						class="rounded-lg border border-white/10 bg-card p-8 transition-colors hover:bg-card-hover md:p-10"
					>
						<h3 class="mb-4 text-xl font-bold md:text-2xl">Practice Difficult Conversations</h3>
						<p class="text-base leading-relaxed text-text-muted">
							Need to have a tough talk? Create a persona that matches the person you're dealing
							with and practice what you want to say. Test different approaches in a low-stakes
							environment.
						</p>
					</div>

					<!-- Use Case 2 -->
					<div
						class="rounded-lg border border-white/10 bg-card p-8 transition-colors hover:bg-card-hover md:p-10"
					>
						<h3 class="mb-4 text-xl font-bold md:text-2xl">Understand Different Perspectives</h3>
						<p class="text-base leading-relaxed text-text-muted">
							Talk to a highly agreeable person, then a highly disagreeable one. See how the same
							topic feels completely different. It's like trying on different ways of thinking.
						</p>
					</div>

					<!-- Use Case 3 -->
					<div
						class="rounded-lg border border-white/10 bg-card p-8 transition-colors hover:bg-card-hover md:p-10"
					>
						<h3 class="mb-4 text-xl font-bold md:text-2xl">Get Advice From Different Angles</h3>
						<p class="text-base leading-relaxed text-text-muted">
							Create a cautious, neurotic advisor and a bold, low-neuroticism one. Get the same
							question answered two completely different ways. Sometimes you need both perspectives.
						</p>
					</div>

					<!-- Use Case 4 -->
					<div
						class="rounded-lg border border-white/10 bg-card p-8 transition-colors hover:bg-card-hover md:p-10"
					>
						<h3 class="mb-4 text-xl font-bold md:text-2xl">Study Partner With Personality</h3>
						<p class="text-base leading-relaxed text-text-muted">
							Make a highly conscientious study buddy who keeps you on track, or an open-minded one
							who loves exploring tangents. Choose the personality that helps you learn best.
						</p>
					</div>
				</div>

				<!-- CTA -->
				<div class="mt-16 text-center md:mt-20">
					<a
						href="/chat/new-persona"
						class="inline-block rounded-lg border-2 border-text bg-text px-8 py-3.5 text-base font-semibold text-background transition-all hover:bg-background hover:text-text"
					>
						Start Experimenting (Free)
					</a>
				</div>
			</div>
		</section>

		<!-- ============================================ -->
		<!-- FOOTER -->
		<!-- ============================================ -->

		<footer class="border-t border-white/10 px-6 py-12 md:px-8">
			<div
				class="mx-auto flex w-full max-w-6xl flex-col items-center justify-between gap-4 sm:flex-row"
			>
				<p class="text-sm text-text-muted">© 2026 Personabox</p>
				<div class="flex gap-6 text-sm">
					<a href="/terms" class="font-medium text-text-muted transition-colors hover:text-text"
						>Terms</a
					>
					<a href="/privacy" class="font-medium text-text-muted transition-colors hover:text-text"
						>Privacy</a
					>
					<a href="/contact" class="font-medium text-text-muted transition-colors hover:text-text"
						>Contact</a
					>
				</div>
			</div>
		</footer>
	</main>
</div>
