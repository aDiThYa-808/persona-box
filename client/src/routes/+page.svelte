<script lang="ts">
    
</script>

<svelte:head>
  <link rel="preconnect" href="https://fonts.googleapis.com">
  <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin="anonymous">
  <link href="https://fonts.googleapis.com/css2?family=Pixelify+Sans:wght@400..700&display=swap" rel="stylesheet">
</svelte:head>

<div class="min-h-screen flex flex-col text-white" style="font-family: 'Pixelify Sans', sans-serif;">
  <header class="flex justify-between items-center px-8 py-6">
    <div class="text-lg font-medium tracking-tight" style="font-family: 'Pixelify Sans', sans-serif;">Persona Box</div>
    <a href="/login" class="px-4 py-2 rounded-md hover:bg-white/10 transition-colors text-[15px]" style="font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', sans-serif;">
      Login
    </a>
  </header>

  <main class="flex-1 flex items-center justify-center px-8 pb-32">
    <div class="text-center max-w-2xl" style="font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', sans-serif;">
      <p class="text-sm text-white/60 mb-3">Introducing PersonaBox</p>
      <h1 class="text-5xl md:text-6xl font-semibold tracking-tight leading-tight mb-8">
        Create custom AI personas
      </h1>
      <a href="/dashboard" class="inline-block bg-white text-[#0f0f0f] px-6 py-3 rounded-lg font-medium hover:bg-white/90 transition-colors">
        Try PersonaBox
      </a>
    </div>
  </main>
</div>