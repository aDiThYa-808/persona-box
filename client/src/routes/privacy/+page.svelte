<script lang="ts">
	import githubIcon from '$lib/assets/icons/github.png';

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
	<title>Privacy Policy - Personabox</title>
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
				on:click={closeDrawer}
			>
				Personabox
			</a>
		</div>

		<!-- Navigation Links -->
		<div class="flex flex-1 flex-col justify-between px-8 py-12">
			<ul class="space-y-5">
				<li>
					<a
						href="/#science"
						class="block text-[15px] font-medium text-text-muted transition-colors hover:text-text"
						on:click={closeDrawer}
					>
						The Science
					</a>
				</li>
				<li>
					<a
						href="/#product"
						class="block text-[15px] font-medium text-text-muted transition-colors hover:text-text"
						on:click={closeDrawer}
					>
						Product
					</a>
				</li>
				<li>
					<a
						href="/#how-it-works"
						class="block text-[15px] font-medium text-text-muted transition-colors hover:text-text"
						on:click={closeDrawer}
					>
						How It Works
					</a>
				</li>
				<li>
					<a
						href="/#for-writers"
						class="block text-[15px] font-medium text-text-muted transition-colors hover:text-text"
						on:click={closeDrawer}
					>
						For Writers
					</a>
				</li>
				<li>
					<a
						href="/#other-uses"
						class="block text-[15px] font-medium text-text-muted transition-colors hover:text-text"
						on:click={closeDrawer}
					>
						Other Uses
					</a>
				</li>
				<li>
					<a
						href="/privacy"
						class="block text-[15px] font-medium text-text-muted transition-colors hover:text-text"
						on:click={closeDrawer}
					>
						Privacy Policy
					</a>
				</li>
				<li>
					<a
						href="/terms"
						class="block text-[15px] font-medium text-text-muted transition-colors hover:text-text"
						on:click={closeDrawer}
					>
						Terms of Use
					</a>
				</li>
			</ul>

			<div class="mt-8 pt-6">
				<p class="text-sm text-text-muted">
					Built by
					<a
						href="https://www.linkedin.com/in/adithya-s-088272287/"
						target="_blank"
						rel="noopener noreferrer"
						class="transition-colors hover:text-text"
					>
						S. Adithya
					</a>
				</p>
			</div>
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
		<!-- PRIVACY POLICY CONTENT -->
		<!-- ============================================ -->

		<div class="px-6 py-20 md:px-8 md:py-32">
			<div class="mx-auto w-full max-w-4xl">
				<!-- Header -->
				<div class="mb-16 md:mb-20">
					<h1 class="mb-4 text-4xl leading-tight font-bold tracking-[-0.02em] md:text-5xl">
						Privacy Policy
					</h1>
					<p class="text-base text-text-muted">Effective Date: 25 Feb 2026</p>
				</div>

				<!-- Introduction -->
				<div class="mb-12 space-y-4 text-base leading-relaxed text-text-muted">
					<p>
						This Privacy Policy describes how PersonaBox ("Service") collects, uses, and processes
						personal data. PersonaBox is operated by S. Adithya, an individual based in India ("I",
						"me", "my").
					</p>
					<p>By using the Service, you agree to this Privacy Policy.</p>
				</div>

				<!-- Sections -->
				<div class="space-y-12">
					<!-- Section 1 -->
					<section>
						<h2 class="mb-4 text-2xl font-bold md:text-3xl">1. Information Collected</h2>
						<div class="space-y-6">
							<div>
								<h3 class="mb-3 text-lg font-semibold text-text">A. Information You Provide</h3>
								<div class="space-y-4 text-base leading-relaxed text-text-muted">
									<p>When you use PersonaBox, the following may be collected:</p>
									<ul class="ml-6 list-disc space-y-2">
										<li>Google account information (such as name and email address)</li>
										<li>Persona configurations you create</li>
										<li>Chat messages and conversation history</li>
										<li>Communications sent to me (e.g., deletion requests)</li>
									</ul>
								</div>
							</div>

							<div>
								<h3 class="mb-3 text-lg font-semibold text-text">
									B. Automatically Collected Information
								</h3>
								<div class="space-y-4 text-base leading-relaxed text-text-muted">
									<p>
										Through hosting and infrastructure providers, technical data may be collected,
										including:
									</p>
									<ul class="ml-6 list-disc space-y-2">
										<li>IP address</li>
										<li>Device and browser information</li>
										<li>Log data</li>
										<li>Usage metrics</li>
									</ul>
								</div>
							</div>
						</div>
					</section>

					<!-- Section 2 -->
					<section>
						<h2 class="mb-4 text-2xl font-bold md:text-3xl">2. How Information Is Used</h2>
						<div class="space-y-4 text-base leading-relaxed text-text-muted">
							<p>Your information is used to:</p>
							<ul class="ml-6 list-disc space-y-2">
								<li>Provide and operate the Service</li>
								<li>Store and retrieve your personas and conversations</li>
								<li>Generate AI responses</li>
								<li>Maintain security and prevent misuse</li>
								<li>Respond to support or deletion requests</li>
								<li>Improve functionality and performance</li>
							</ul>
							<p>
								PersonaBox is an experimental project and does not use your data for advertising or
								resale.
							</p>
						</div>
					</section>

					<!-- Section 3 -->
					<section>
						<h2 class="mb-4 text-2xl font-bold md:text-3xl">3. AI Processing</h2>
						<div class="space-y-4 text-base leading-relaxed text-text-muted">
							<p>When you send messages in PersonaBox:</p>
							<ul class="ml-6 list-disc space-y-2">
								<li>Your prompts are transmitted to OpenAI's API for AI response generation.</li>
								<li>Conversations may be processed by OpenAI according to their policies.</li>
								<li>AI outputs are generated by third-party models and may not be accurate.</li>
							</ul>
							<p>By using the Service, you consent to this processing.</p>
						</div>
					</section>

					<!-- Section 4 -->
					<section>
						<h2 class="mb-4 text-2xl font-bold md:text-3xl">4. Third-Party Services</h2>
						<div class="space-y-4 text-base leading-relaxed text-text-muted">
							<p>PersonaBox relies on third-party providers, including but not limited to:</p>
							<ul class="ml-6 list-disc space-y-2">
								<li>AWS (cloud infrastructure and data storage)</li>
								<li>OpenAI (AI model provider)</li>
								<li>Google (authentication)</li>
								<li>Vercel (hosting and deployment)</li>
							</ul>
							<p>These providers may process data in accordance with their own privacy policies.</p>
							<p>I am not responsible for the independent practices of third-party providers.</p>
						</div>
					</section>

					<!-- Section 5 -->
					<section>
						<h2 class="mb-4 text-2xl font-bold md:text-3xl">5. Data Storage and Retention</h2>
						<div class="space-y-4 text-base leading-relaxed text-text-muted">
							<ul class="ml-6 list-disc space-y-2">
								<li>Chat data and persona configurations are stored in cloud infrastructure.</li>
								<li>
									Data is retained as long as the Service operates or until manually deleted upon
									request.
								</li>
								<li>There is no guaranteed retention period.</li>
							</ul>
							<p>If the Service is discontinued, data may be deleted without notice.</p>
						</div>
					</section>

					<!-- Section 6 -->
					<section>
						<h2 class="mb-4 text-2xl font-bold md:text-3xl">6. Data Deletion Requests</h2>
						<div class="space-y-4 text-base leading-relaxed text-text-muted">
							<p>There is currently no self-service deletion feature.</p>
							<p>You may request deletion of your account and associated data by contacting:</p>
							<a
								href="mailto:contactadiiiii@gmail.com"
								target="_blank"
								class="font-semibold text-text">contactadiiiii@gmail.com</a
							>
							<p>Deletion requests are handled manually and may take reasonable time to process.</p>
						</div>
					</section>

					<!-- Section 7 -->
					<section>
						<h2 class="mb-4 text-2xl font-bold md:text-3xl">7. Age Requirement</h2>
						<div class="space-y-4 text-base leading-relaxed text-text-muted">
							<p>PersonaBox is intended for users aged 13 and above.</p>
							<p>
								By using the Service, you represent that you meet this requirement. I do not
								independently verify age.
							</p>
						</div>
					</section>

					<!-- Section 8 -->
					<section>
						<h2 class="mb-4 text-2xl font-bold md:text-3xl">8. Data Security</h2>
						<div class="space-y-4 text-base leading-relaxed text-text-muted">
							<p>Reasonable technical measures are used to protect stored data.</p>
							<p>However:</p>
							<ul class="ml-6 list-disc space-y-2">
								<li>No system is completely secure.</li>
								<li>I cannot guarantee absolute security of data transmitted or stored.</li>
								<li>Use of the Service is at your own risk.</li>
							</ul>
						</div>
					</section>

					<!-- Section 9 -->
					<section>
						<h2 class="mb-4 text-2xl font-bold md:text-3xl">9. International Data Processing</h2>
						<div class="space-y-4 text-base leading-relaxed text-text-muted">
							<p>
								Because PersonaBox uses global cloud providers, your data may be processed or stored
								outside India.
							</p>
							<p>By using the Service, you consent to such transfers.</p>
						</div>
					</section>

					<!-- Section 10 -->
					<section>
						<h2 class="mb-4 text-2xl font-bold md:text-3xl">10. No Sale of Data</h2>
						<div class="space-y-4 text-base leading-relaxed text-text-muted">
							<p>PersonaBox does not sell personal data.</p>
						</div>
					</section>

					<!-- Section 11 -->
					<section>
						<h2 class="mb-4 text-2xl font-bold md:text-3xl">11. Changes to This Policy</h2>
						<div class="space-y-4 text-base leading-relaxed text-text-muted">
							<p>This Privacy Policy may be updated at any time.</p>
							<p>
								Continued use of the Service after changes constitutes acceptance of the updated
								policy.
							</p>
						</div>
					</section>

					<!-- Section 12 -->
					<section>
						<h2 class="mb-4 text-2xl font-bold md:text-3xl">12. Contact</h2>
						<div class="space-y-4 text-base leading-relaxed text-text-muted">
							<p>For questions or data-related requests, contact:</p>
							<a
								href="mailto:contactadiiiii@gmail.com"
								target="_blank"
								class="font-semibold text-text">contactadiiiii@gmail.com</a
							>
						</div>
					</section>
				</div>
			</div>
		</div>

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
					<a href="/privacy" class="font-medium text-text transition-colors hover:text-text"
						>Privacy</a
					>
					<a
						href="mailto:contactadiiiii@gmail.com"
						class="font-medium text-text-muted transition-colors hover:text-text">Contact</a
					>
				</div>
			</div>
		</footer>
	</main>
</div>
