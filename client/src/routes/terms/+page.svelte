<script lang="ts">
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
	<title>Terms of Use - Personabox</title>
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
		<!-- TERMS CONTENT -->
		<!-- ============================================ -->

		<div class="px-6 py-20 md:px-8 md:py-32">
			<div class="mx-auto w-full max-w-4xl">
				<!-- Header -->
				<div class="mb-16 md:mb-20">
					<h1 class="mb-4 text-4xl leading-tight font-bold tracking-[-0.02em] md:text-5xl">
						Terms of Use
					</h1>
					<p class="text-base text-text-muted">Effective Date: 25 Feb 2026</p>
				</div>

				<!-- Introduction -->
				<div class="mb-12 space-y-4 text-base leading-relaxed text-text-muted">
					<p>
						Welcome to PersonaBox. These Terms of Use ("Terms") govern your access to and use of
						PersonaBox (the "Service"), operated by S. Adithya, an individual based in India ("I",
						"me", "my").
					</p>
					<p>
						By accessing or using PersonaBox, you agree to be bound by these Terms. If you do not
						agree, do not use the Service.
					</p>
				</div>

				<!-- Sections -->
				<div class="space-y-12">
					<!-- Section 1 -->
					<section>
						<h2 class="mb-4 text-2xl font-bold md:text-3xl">1. Eligibility</h2>
						<div class="space-y-4 text-base leading-relaxed text-text-muted">
							<p>You must be at least 13 years old to use PersonaBox.</p>
							<p>By using the Service, you represent and warrant that:</p>
							<ul class="ml-6 list-disc space-y-2">
								<li>You are at least 13 years old.</li>
								<li>You have the legal capacity to agree to these Terms.</li>
								<li>All information you provide is accurate.</li>
							</ul>
							<p>I do not independently verify age.</p>
						</div>
					</section>

					<!-- Section 2 -->
					<section>
						<h2 class="mb-4 text-2xl font-bold md:text-3xl">2. Nature of the Service</h2>
						<div class="space-y-4 text-base leading-relaxed text-text-muted">
							<p>
								PersonaBox is an experimental AI-based web application that allows users to create
								AI personas and engage in AI-generated conversations.
							</p>
							<p>The Service:</p>
							<ul class="ml-6 list-disc space-y-2">
								<li>Uses third-party infrastructure and AI models.</li>
								<li>May generate unpredictable, inaccurate, or inappropriate outputs.</li>
								<li>Is provided for informational and entertainment purposes only.</li>
							</ul>
							<p>
								PersonaBox is a personal experimental project and is not a commercial SaaS product.
							</p>
						</div>
					</section>

					<!-- Section 3 -->
					<section>
						<h2 class="mb-4 text-2xl font-bold md:text-3xl">3. AI-Generated Content Disclaimer</h2>
						<div class="space-y-4 text-base leading-relaxed text-text-muted">
							<p>
								All responses are generated by artificial intelligence models provided by third
								parties (including OpenAI).
							</p>
							<p>You acknowledge and agree that:</p>
							<ul class="ml-6 list-disc space-y-2">
								<li>AI outputs may be inaccurate, misleading, offensive, or inappropriate.</li>
								<li>
									Responses do not constitute professional advice (including legal, medical,
									financial, or psychological advice).
								</li>
								<li>You use all outputs at your own risk.</li>
								<li>I am not responsible for decisions made based on AI-generated content.</li>
							</ul>
						</div>
					</section>

					<!-- Section 4 -->
					<section>
						<h2 class="mb-4 text-2xl font-bold md:text-3xl">4. Accounts and Access</h2>
						<div class="space-y-4 text-base leading-relaxed text-text-muted">
							<ul class="ml-6 list-disc space-y-2">
								<li>Access requires Google authentication.</li>
								<li>You are responsible for maintaining the security of your Google account.</li>
								<li>
									I reserve the right to suspend or terminate your access at any time, for any
									reason, without notice.
								</li>
								<li>
									There is no guarantee of continued availability of your account, data, or the
									Service.
								</li>
							</ul>
						</div>
					</section>

					<!-- Section 5 -->
					<section>
						<h2 class="mb-4 text-2xl font-bold md:text-3xl">5. User Content and License</h2>
						<div class="space-y-4 text-base leading-relaxed text-text-muted">
							<p>You may create personas and generate conversations ("User Content").</p>
							<p>You retain ownership of the content you create.</p>
							<p>
								However, by using PersonaBox, you grant me a worldwide, non-exclusive, royalty-free,
								irrevocable license to:
							</p>
							<ul class="ml-6 list-disc space-y-2">
								<li>Store</li>
								<li>Process</li>
								<li>Modify</li>
								<li>Display</li>
								<li>Remove</li>
							</ul>
							<p>your User Content as necessary to operate, maintain, or improve the Service.</p>
							<p>
								I reserve the right to remove any content that comes to my attention and is deemed
								inappropriate, unlawful, or harmful.
							</p>
							<p>PersonaBox does not guarantee monitoring or moderation of content.</p>
						</div>
					</section>

					<!-- Section 6 -->
					<section>
						<h2 class="mb-4 text-2xl font-bold md:text-3xl">6. Prohibited Conduct</h2>
						<div class="space-y-4 text-base leading-relaxed text-text-muted">
							<p>You agree not to use the Service to:</p>
							<ul class="ml-6 list-disc space-y-2">
								<li>Engage in illegal activity</li>
								<li>Generate or distribute hate speech</li>
								<li>Harass or threaten others</li>
								<li>Create sexually explicit or unlawful content</li>
								<li>Impersonate individuals</li>
								<li>Attempt to reverse engineer, exploit, or interfere with the Service</li>
							</ul>
							<p>Violation may result in immediate termination.</p>
						</div>
					</section>

					<!-- Section 7 -->
					<section>
						<h2 class="mb-4 text-2xl font-bold md:text-3xl">7. No Moderation Guarantee</h2>
						<div class="space-y-4 text-base leading-relaxed text-text-muted">
							<ul class="ml-6 list-disc space-y-2">
								<li>PersonaBox does not actively monitor conversations.</li>
								<li>Content may not be reviewed unless brought to my attention.</li>
								<li>
									You acknowledge that you may encounter offensive or inappropriate material and
									agree that you use the Service at your own risk.
								</li>
							</ul>
						</div>
					</section>

					<!-- Section 8 -->
					<section>
						<h2 class="mb-4 text-2xl font-bold md:text-3xl">8. Third-Party Services</h2>
						<div class="space-y-4 text-base leading-relaxed text-text-muted">
							<p>PersonaBox relies on third-party providers, including but not limited to:</p>
							<ul class="ml-6 list-disc space-y-2">
								<li>AWS (hosting and infrastructure)</li>
								<li>OpenAI (AI model provider)</li>
								<li>Google (authentication)</li>
								<li>Vercel (deployment and hosting)</li>
							</ul>
							<p>
								Your data may be processed by these providers in accordance with their respective
								terms and policies.
							</p>
							<p>I am not responsible for the practices or failures of third-party services.</p>
						</div>
					</section>

					<!-- Section 9 -->
					<section>
						<h2 class="mb-4 text-2xl font-bold md:text-3xl">9. Data Deletion</h2>
						<div class="space-y-4 text-base leading-relaxed text-text-muted">
							<p>There is currently no self-service account deletion feature.</p>
							<p>If you wish to request deletion of your data, you may email:</p>
							<a
								href="mailto:contactadiiiii@gmail.com"
								target="_blank"
								class="font-semibold text-text">contactadiiiii@gmail.com</a
							>
							<p>Deletion requests will be handled manually and at my discretion.</p>
						</div>
					</section>

					<!-- Section 10 -->
					<section>
						<h2 class="mb-4 text-2xl font-bold md:text-3xl">10. Donations</h2>
						<div class="space-y-4 text-base leading-relaxed text-text-muted">
							<p>PersonaBox is free to use.</p>
							<p>Users may voluntarily support the project through Buy Me a Coffee.</p>
							<p>
								All donations are voluntary, non-refundable, and do not grant additional rights,
								ownership, or service guarantees.
							</p>
						</div>
					</section>

					<!-- Section 11 -->
					<section>
						<h2 class="mb-4 text-2xl font-bold md:text-3xl">11. No Warranties</h2>
						<div class="space-y-4 text-base leading-relaxed text-text-muted">
							<p>The Service is provided "AS IS" and "AS AVAILABLE."</p>
							<p>I make no warranties, express or implied, including but not limited to:</p>
							<ul class="ml-6 list-disc space-y-2">
								<li>Accuracy</li>
								<li>Reliability</li>
								<li>Fitness for a particular purpose</li>
								<li>Non-infringement</li>
								<li>Continuous availability</li>
							</ul>
							<p>Your use of the Service is entirely at your own risk.</p>
						</div>
					</section>

					<!-- Section 12 -->
					<section>
						<h2 class="mb-4 text-2xl font-bold md:text-3xl">12. Limitation of Liability</h2>
						<div class="space-y-4 text-base leading-relaxed text-text-muted">
							<p>To the maximum extent permitted by law:</p>
							<p>
								I shall not be liable for any direct, indirect, incidental, consequential, special,
								or punitive damages arising from or related to:
							</p>
							<ul class="ml-6 list-disc space-y-2">
								<li>Use of the Service</li>
								<li>AI-generated content</li>
								<li>Service interruptions</li>
								<li>Data loss</li>
								<li>Third-party failures</li>
							</ul>
							<p>
								If liability cannot be excluded under applicable law, it shall be limited to the
								amount you have paid for the Service in the last 12 months (which is ₹0).
							</p>
						</div>
					</section>

					<!-- Section 13 -->
					<section>
						<h2 class="mb-4 text-2xl font-bold md:text-3xl">
							13. Termination and Service Discontinuation
						</h2>
						<div class="space-y-4 text-base leading-relaxed text-text-muted">
							<p>I reserve the right to:</p>
							<ul class="ml-6 list-disc space-y-2">
								<li>Modify</li>
								<li>Suspend</li>
								<li>Discontinue</li>
								<li>Shut down</li>
							</ul>
							<p>PersonaBox at any time, without notice or liability.</p>
							<p>There is no guarantee of data retention or service continuity.</p>
						</div>
					</section>

					<!-- Section 14 -->
					<section>
						<h2 class="mb-4 text-2xl font-bold md:text-3xl">14. Governing Law</h2>
						<div class="space-y-4 text-base leading-relaxed text-text-muted">
							<p>
								These Terms shall be governed by and construed in accordance with the laws of India.
							</p>
							<p>
								Any disputes shall be subject to the exclusive jurisdiction of the courts in India.
							</p>
						</div>
					</section>

					<!-- Section 15 -->
					<section>
						<h2 class="mb-4 text-2xl font-bold md:text-3xl">15. Changes to Terms</h2>
						<div class="space-y-4 text-base leading-relaxed text-text-muted">
							<p>
								I may update these Terms at any time. Continued use of the Service after changes
								constitutes acceptance of the revised Terms.
							</p>
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
					<a href="/terms" class="font-medium text-text transition-colors hover:text-text">Terms</a>
					<a href="/privacy" class="font-medium text-text-muted transition-colors hover:text-text"
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
